# Claude Instructions

Instructions for Claude Code and Claude-based agents working in this repository.

## Repository Purpose

This is an AI Workbench - a public notebook for prompts, patterns, templates, and experiments across AI ecosystems. The philosophy is **Capture → Curate → Canonicalize**.

## Key Conventions

### File Locations

- `inbox/` - Quick capture, unstructured. Use `YYYY-MM-DD__slug.md` naming.
- `library/` - Curated, reusable content with frontmatter metadata.
- `ecosystems/claude/` - Claude-specific content goes here, not in general library.
- `experiments/NNNN__name/` - Numbered experiment folders for research.
- `meta/` - Repo metadata (taxonomy, glossary, changelog).

### Prompt Format

Prompts in `library/` should use YAML frontmatter:

```yaml
---
id: category.name.v1
tags: [tag1, tag2]
targets: [claude, chatgpt, codex]
inputs:
  - name: variable_name
output: markdown
---

## Prompt

Your prompt text with {{variable_name}} placeholders.
```

### When Adding Content

1. **New ideas**: Create in `inbox/` with date prefix
2. **Claude-specific**: Put in `ecosystems/claude/`
3. **General utility**: Put in `library/` with proper frontmatter
4. **Research**: Create numbered folder in `experiments/`

### What NOT to Do

- Never commit secrets, API keys, or credentials
- Don't put Claude-specific quirks in general `library/` - use `ecosystems/claude/`
- Don't overwrite prompt history - create new versions instead (v1 → v2)

## Helpful Context

- This repo is MIT licensed and fully public
- Structure is documented but folders are created as needed
- See README.md for full structure and workflow details
- See CONTRIBUTING.md for contribution guidelines
