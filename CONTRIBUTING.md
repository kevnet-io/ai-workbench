# Contributing

Contributions welcome! This repo grows organically - here's how to add content.

## Quick Capture (inbox/)

For fast, unstructured capture:

1. Create a file in `inbox/` with the format `YYYY-MM-DD__slug.md`
2. Dump your content - no formatting required
3. That's it

Example: `inbox/2025-12-15__claude-code-tips.md`

## Promoting to Library

When content has proven useful (used more than twice):

1. Move from `inbox/` to appropriate location in `library/`
2. Add YAML frontmatter (see README for format)
3. Give it a versioned `id` (e.g., `coding.review.v1`)

## Adding Tool-Specific Content

Content specific to a particular AI tool goes in `ecosystems/`:

- `ecosystems/claude/` - Claude, Claude Code
- `ecosystems/chatgpt/` - ChatGPT, GPTs, actions
- `ecosystems/codex/` - OpenAI Codex, CLI
- `ecosystems/local/` - Ollama, LMStudio, etc.

## Creating Experiments

For research or architecture exploration:

1. Create a numbered folder: `experiments/NNNN__descriptive-name/`
2. Include a `README.md` with hypothesis and setup
3. Use `notes.md` for working log
4. Put variants in `variants/`, results in `results/`

## Pull Requests

- Keep PRs focused - one concept per PR
- Follow existing patterns and naming conventions
- Ensure no secrets or private data are included
- Update `meta/changelog.md` for significant additions

## Issues

Use issues for:

- Suggesting new patterns or templates
- Reporting problems with existing content
- Proposing structural changes
