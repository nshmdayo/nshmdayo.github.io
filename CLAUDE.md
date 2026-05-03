# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Install dependencies
go mod download

# Build the site (outputs to artifact/)
go run main.go

# Build and serve locally with auto-rebuild on file changes
go run main.go -serve

# Serve on a custom port
go run main.go -serve -port 3000
```

There are no tests in this project.

## Architecture

This is a Go-based static site generator that builds a portfolio and blog site deployed to GitHub Pages.

**Build pipeline** (`main.go`): `buildSite()` reads `config.yaml` and Markdown files from `content/post/`, then renders Go `html/template` templates from `template/` into `artifact/`. The output directory structure is `artifact/css/style.css`, `artifact/index.html`, and `artifact/blog/<slug>.html`.

**Data flow**:
- `config.yaml` → `Config` struct (name, description, about, skills, projects)
- `content/post/*.md` → `Post` structs (YAML front matter delimited by `---`, body parsed with goldmark)
- Both are combined into `PageData` and passed to templates

**Templates** (`template/*.tpl`): Standard Go `html/template` files. `index.html.tpl` renders the home page with tabbed About/Blog sections. `post.html.tpl` renders individual blog posts. `style.css.tpl` renders the stylesheet (no dynamic data, rendered with `nil`).

**Watch mode**: `watchAndRebuild()` polls `config.yaml`, templates, and `content/post/` every 500ms and calls `buildSite()` on any change.

**Deployment**: GitHub Actions (`.github/workflows/deploy.yml`) runs `go run main.go` on push to `main`, uploads `artifact/` as a Pages artifact, and deploys it.

## Adding Content

**Blog posts** — add `.md` files to `content/post/` with this front matter:
```markdown
---
title: "Post Title"
date: "2025-10-18"
description: "Post summary"
tags: ["Tag1", "Tag2"]
---
```
Posts are sorted by date descending. The filename (without `.md`) becomes the URL slug.

**Site config** — edit `config.yaml` to update name, description, about text, skills, and projects.
