# Agent Instructions

General guidance for AI agents working in this repository.

## What This Repo Is

An AI Workbench - a public collection of:

- **Prompts**: Reusable prompt templates with metadata
- **Patterns**: Prompt engineering techniques and strategies
- **Experiments**: Research explorations with reproducible setups
- **Ecosystem notes**: Tool-specific content (Claude, ChatGPT, Codex, local models)

## Repository Structure

```
ai-workbench/
├── inbox/              # Unstructured capture (YYYY-MM-DD__slug.md)
├── library/
│   ├── prompts/        # Curated prompts with frontmatter
│   ├── templates/      # Prompt templates
│   └── patterns/       # Engineering patterns
├── ecosystems/
│   ├── claude/         # Claude-specific
│   ├── chatgpt/        # ChatGPT, GPTs
│   ├── codex/          # Codex workflows
│   └── local/          # Ollama, LMStudio, etc.
├── experiments/        # Numbered research folders
└── meta/               # Taxonomy, glossary, changelog
```

Folders are created as needed - not all may exist yet.

## Content Format

Prompts use YAML frontmatter for metadata:

```yaml
---
id: category.name.v1
tags: [relevant, tags]
targets: [claude, chatgpt]
inputs:
  - name: input_variable
output: markdown
---

## Prompt

Prompt content with {{input_variable}} placeholders.
```

## Guidelines for Agents

### Do

- Follow existing naming conventions
- Use frontmatter for prompts in `library/`
- Put tool-specific content in appropriate `ecosystems/` subfolder
- Create versioned IDs for prompts (v1, v2, etc.)
- Keep experiment folders self-contained with README

### Don't

- Commit secrets, API keys, or credentials
- Mix tool-specific content with general library content
- Overwrite existing prompts - create new versions
- Create empty placeholder folders

## Key Files

- `README.md` - Full documentation
- `CONTRIBUTING.md` - How to add content
- `SECURITY.md` - What not to commit
- `CLAUDE.md` - Claude-specific agent instructions
