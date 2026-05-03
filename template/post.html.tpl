<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Post.Title}} — {{.Config.Name}}</title>
    <meta name="description" content="{{.Post.Description}}">
    <link rel="stylesheet" href="../css/style.css">
</head>
<body>
    <div class="page">
        <header>
            <nav>
                <a href="/" class="nav-home">{{.Config.Name}}</a>
            </nav>
        </header>

        <div class="post-header">
            <a href="../index.html" class="post-back">← Writing</a>
            <h1>{{.Post.Title}}</h1>
            <div class="meta">{{.Post.Date}}</div>
            {{if .Post.Tags}}
            <div class="tags">
                {{range .Post.Tags}}
                <span class="tag">{{.}}</span>
                {{end}}
            </div>
            {{end}}
        </div>

        <article class="post-body">
            {{.Post.Content}}
        </article>

        <div class="post-footer">
            <a href="../index.html#tab-post" class="post-back">← Writing</a>
        </div>
    </div>

    <footer>
        <span>© 2026 {{.Config.Name}}</span>
    </footer>
</body>
</html>
