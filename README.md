# AI Workbench

A public notebook, pattern library, and experiment space for AI-related research, prompts, templates, and tooling across multiple ecosystems.

## Philosophy

**Capture → Curate → Canonicalize**

- **Capture**: Frictionless inbox for random ideas, prompts, and notes
- **Curate**: Periodically move valuable content into a structured library
- **Canonicalize**: Templates with metadata, examples, and versioning

Start small, grow organically. Minimal structure, maximum utility.

## Structure

Folders are created as needed. This is the expected layout:

```
ai-workbench/
├── inbox/                    # Quick capture - dump ideas here
├── library/
│   ├── prompts/             # Curated, reusable prompts
│   ├── templates/           # Prompt templates with frontmatter
│   └── patterns/            # Prompt engineering patterns
├── ecosystems/
│   ├── claude/              # Claude-specific content
│   ├── chatgpt/             # ChatGPT, GPTs, actions
│   ├── codex/               # Codex workflows
│   └── local/               # Ollama, LMStudio, local models
├── experiments/             # Numbered experiment folders
└── meta/                    # Taxonomy, glossary, changelog
```

### inbox/

Drop anything here with the naming convention `YYYY-MM-DD__slug.md`. No structure required - this is for fast capture.

### library/

The "source of truth" for reusable content. Prompts here should follow the frontmatter format (see below).

### ecosystems/

Tool-specific content that doesn't belong in the general library. Claude quirks stay in `ecosystems/claude/`, ChatGPT-specific GPT configs go in `ecosystems/chatgpt/`, etc.

### experiments/

For architecture explorations and research. Each experiment gets a numbered folder:

```
experiments/0001__prompt-routing/
├── README.md        # Hypothesis, setup, how to reproduce
├── notes.md         # Messy working log
├── variants/        # Prompt variants, configs
└── results/         # Outputs, tables, screenshots
```

### meta/

Repo-level metadata: taxonomy definitions, glossary, changelog.

## Prompt Format

Use **one prompt per file** with **YAML frontmatter** for searchability:

```yaml
---
id: coding.review.v1
tags: [coding, review, rubric]
targets: [claude, chatgpt, codex]
inputs:
  - name: context
  - name: diff
output: markdown
---

## Prompt

You are reviewing a PR. Use this rubric:
1) correctness 2) readability 3) security 4) performance 5) tests

Context:
{{context}}

Diff:
{{diff}}
```

This format enables:
- Searching by `id:` or `tags:`
- Building tooling to render `{{variables}}`
- Tracking prompt versions

## Workflow

### Adding Content

1. **Quick capture**: Drop into `inbox/` with date-prefixed filename
2. **Weekly triage** (10 mins): Move items from inbox to appropriate location
   - General utility → `library/`
   - Tool-specific → `ecosystems/<tool>/`
   - Research/architecture → `experiments/`
   - One-off junk → delete

### Versioning

If a prompt changes materially, bump the version in `id` (e.g., `v1` → `v2`) and keep the old file. Don't overwrite history unless it's a typo fix.

### Evergreen Rule

If you've used something more than twice, it doesn't belong in inbox.

## License

MIT - see [LICENSE](LICENSE)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on adding content, promoting from inbox, and creating experiments.

## Security

See [SECURITY.md](SECURITY.md) - no secrets, no API keys, no private data.
