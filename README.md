# Portfolio & Blog

A portfolio and blog site built with a static site generator written in Go.

## 🚀 Tech Stack

- **Go 1.25** - Static site generation
- **goldmark** - Markdown parser
- **yaml.v3** - YAML configuration files
- **GitHub Actions** - CI/CD
- **GitHub Pages** - Hosting

## 📦 Features

- **Responsive Design** - Mobile-first optimization
- **Blog** - Markdown-based technical blog
- **Fast Build** - High-speed static site generation with Go
- **Auto Deploy** - Automatic build and deployment with GitHub Actions on push

## 🛠️ Development & Deployment

### Install Dependencies

```bash
go mod download
```

### Generate Site Locally

```bash
go run main.go
```

The generated site will be output to the `artifact/` directory.

### Local Preview

```bash
# Preview with Go's simple server
go run main.go -serve

# Or specify a port number
go run main.go -serve -port 3000

# Open http://localhost:8000 in your browser
```

### Deploy to GitHub Pages

When you push to the `main` branch, GitHub Actions automatically builds and deploys the site.

```bash
git add .
git commit -m "Update content"
git push origin main
```

## 📝 Content Management

### Configuration File

Configure basic site information in `config.yaml`.

```yaml
name: "Your Name"
title: "Portfolio & Blog"
description: "Software Engineer | Go, TypeScript, Cloud"
github: "https://github.com/yourusername"
```

### Adding Blog Posts

Add `.md` files to the `content/post/` directory.

```markdown
---
title: "Post Title"
date: "2025-10-18"
description: "Post summary"
tags: ["Tag1", "Tag2"]
---

# Post Content

Write your post in Markdown.
```

## 🏗️ Project Structure

```
.
├── main.go                 # Main program
├── go.mod                  # Go module definition
├── config.yaml             # Site configuration
├── content/
│   └── post/              # Blog posts (Markdown)
├── artifact/              # Generated HTML (for GitHub Pages)
│   ├── index.html
│   ├── css/
│   │   └── style.css
│   └── blog/
│       └── *.html
├── .github/
│   └── workflows/
│       └── deploy.yml     # GitHub Actions configuration
└── README.md
```

## 🌍 GitHub Pages Setup

1. Go to your GitHub repository Settings → Pages
2. Source: Select "GitHub Actions"
3. Automatic deployment occurs when pushing to the `main` branch

## 📄 License

This project is licensed under the MIT License.