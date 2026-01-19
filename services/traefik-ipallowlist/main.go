package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var allowedNets []*net.IPNet

func init() {
	// Load allowed IPs/CIDRs from environment variable
	allowlistStr := os.Getenv("ALLOWED_IPS")
	if allowlistStr == "" {
		log.Fatal("ALLOWED_IPS environment variable is required")
	}

	// Parse comma-separated list of IPs and CIDRs
	entries := strings.Split(allowlistStr, ",")
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}

		// Check if it's a CIDR or a single IP
		var ipNet *net.IPNet
		var err error

		if strings.Contains(entry, "/") {
			// It's a CIDR
			_, ipNet, err = net.ParseCIDR(entry)
		} else {
			// It's a single IP, convert to CIDR
			ip := net.ParseIP(entry)
			if ip == nil {
				log.Fatalf("Invalid IP address: %s", entry)
			}
			// Determine if IPv4 or IPv6 and create appropriate mask
			if ip.To4() != nil {
				_, ipNet, err = net.ParseCIDR(entry + "/32")
			} else {
				_, ipNet, err = net.ParseCIDR(entry + "/128")
			}
		}

		if err != nil {
			log.Fatalf("Invalid IP or CIDR: %s - %v", entry, err)
		}

		allowedNets = append(allowedNets, ipNet)
		log.Printf("Added to allowlist: %s", ipNet.String())
	}

	if len(allowedNets) == 0 {
		log.Fatal("No valid IPs or CIDRs found in ALLOWED_IPS")
	}
}

// extractClientIP extracts the real client IP from the request
func extractClientIP(r *http.Request) (string, error) {
	// Check X-Forwarded-For header first (Traefik typically sets this)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0]), nil
		}
	}

	// Check X-Real-IP header
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return strings.TrimSpace(xri), nil
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", fmt.Errorf("failed to parse RemoteAddr: %v", err)
	}

	return ip, nil
}

// isIPAllowed checks if the given IP is in the allowlist
func isIPAllowed(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	for _, allowedNet := range allowedNets {
		if allowedNet.Contains(ip) {
			return true
		}
	}

	return false
}

// authHandler handles the authentication requests from Traefik
func authHandler(w http.ResponseWriter, r *http.Request) {
	clientIP, err := extractClientIP(r)
	if err != nil {
		log.Printf("Error extracting client IP: %v", err)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if isIPAllowed(clientIP) {
		log.Printf("Allowed: %s", clientIP)
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Printf("Denied: %s", clientIP)
	http.Error(w, "Forbidden", http.StatusForbidden)
}

// healthHandler provides a health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", authHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("Starting IP allowlist service on port %s", port)
	log.Printf("Allowlist contains %d entries", len(allowedNets))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
