# Traefik IP Allowlist Service

A lightweight Go service that provides IP-based access control for Traefik using the forwardAuth middleware.

## Features

- ✅ Simple IP and CIDR-based allowlisting
- ✅ Works with Traefik's forwardAuth middleware
- ✅ Handles X-Forwarded-For and X-Real-IP headers
- ✅ Minimal Docker image (~2MB using scratch base)
- ✅ Health check endpoint
- ✅ Configurable via environment variables

## Quick Start

### Building the Docker Image

```bash
docker build -t traefik-ipallowlist:latest .
```

### Running the Service

```bash
docker run -d \
  --name ipallowlist \
  -p 8080:8080 \
  -e ALLOWED_IPS="192.168.1.0/24,10.0.0.1,2001:db8::/32" \
  traefik-ipallowlist:latest
```

### Configuration

The service is configured using environment variables:

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `ALLOWED_IPS` | Yes | - | Comma-separated list of allowed IPs and CIDRs |
| `PORT` | No | `8080` | Port the service listens on |

#### Examples:

**Single IP:**
```bash
ALLOWED_IPS="192.168.1.100"
```

**Multiple IPs:**
```bash
ALLOWED_IPS="192.168.1.100,10.0.0.50"
```

**CIDR ranges:**
```bash
ALLOWED_IPS="192.168.1.0/24,10.0.0.0/8"
```

**Mixed:**
```bash
ALLOWED_IPS="192.168.1.100,10.0.0.0/8,172.16.0.0/12"
```

**IPv6 support:**
```bash
ALLOWED_IPS="2001:db8::/32,192.168.1.0/24"
```

## Traefik Integration

### Docker Compose Example

```yaml
version: '3.8'

services:
  traefik:
    image: traefik:v2.10
    command:
      - "--api.insecure=true"
      - "--providers.docker=true"
      - "--entrypoints.web.address=:80"
    ports:
      - "80:80"
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

  ipallowlist:
    build: .
    environment:
      - ALLOWED_IPS=192.168.1.0/24,10.0.0.1
      - PORT=8080
    labels:
      - "traefik.enable=true"
      - "traefik.http.services.ipallowlist.loadbalancer.server.port=8080"

      # ForwardAuth middleware
      - "traefik.http.middlewares.ip-allowlist.forwardauth.address=http://ipallowlist:8080"
      - "traefik.http.middlewares.ip-allowlist.forwardauth.trustForwardHeader=true"

  myapp:
    image: nginx:alpine
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.myapp.rule=Host(`myapp.local`)"
      - "traefik.http.routers.myapp.entrypoints=web"

      # Apply the IP allowlist middleware
      - "traefik.http.routers.myapp.middlewares=ip-allowlist@docker"
```

### Kubernetes Example

```yaml
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: ipallowlist-config
data:
  ALLOWED_IPS: "192.168.1.0/24,10.0.0.0/8"

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ipallowlist
spec:
  replicas: 2
  selector:
    matchLabels:
      app: ipallowlist
  template:
    metadata:
      labels:
        app: ipallowlist
    spec:
      containers:
      - name: ipallowlist
        image: traefik-ipallowlist:latest
        ports:
        - containerPort: 8080
        envFrom:
        - configMapRef:
            name: ipallowlist-config
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 3
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 3
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: ipallowlist
spec:
  selector:
    app: ipallowlist
  ports:
  - port: 8080
    targetPort: 8080

---
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: ip-allowlist
spec:
  forwardAuth:
    address: http://ipallowlist:8080
    trustForwardHeader: true

---
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: myapp
spec:
  entryPoints:
    - web
  routes:
  - match: Host(`myapp.example.com`)
    kind: Rule
    middlewares:
    - name: ip-allowlist
    services:
    - name: myapp
      port: 80
```

## Testing

### Test the health endpoint:
```bash
curl http://localhost:8080/health
```

### Test from an allowed IP:
```bash
curl -H "X-Forwarded-For: 192.168.1.100" http://localhost:8080/
# Should return 200 OK
```

### Test from a denied IP:
```bash
curl -H "X-Forwarded-For: 1.2.3.4" http://localhost:8080/
# Should return 403 Forbidden
```

## Endpoints

- `/` - Main authentication endpoint (used by Traefik forwardAuth)
  - Returns `200 OK` if IP is allowed
  - Returns `403 Forbidden` if IP is denied

- `/health` - Health check endpoint
  - Always returns `200 OK` if service is running

## IP Detection

The service checks headers in the following order:

1. `X-Forwarded-For` (uses first IP in the list)
2. `X-Real-IP`
3. `RemoteAddr` (direct connection IP)

This ensures compatibility with various proxy configurations.

## Building from Source

```bash
go build -o ipallowlist main.go
```

## Running Locally

```bash
ALLOWED_IPS="127.0.0.1,192.168.1.0/24" ./ipallowlist
```

## Security Considerations

- Always use `trustForwardHeader: true` in Traefik when behind a proxy
- Ensure Traefik is properly configured to set X-Forwarded-For headers
- The service trusts the X-Forwarded-For header, so ensure it's only accessible from Traefik
- Do not expose this service directly to the internet
- Regularly review and update your allowlist

## License

MIT
