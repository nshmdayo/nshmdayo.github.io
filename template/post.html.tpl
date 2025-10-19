<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Post.Title}} - {{.Config.Name}}</title>
    <meta name="description" content="{{.Post.Description}}">
    <link rel="stylesheet" href="../css/style.css">
</head>
<body>
    <main class="container">
        <article>
            <h1>{{.Post.Title}}</h1>
            <div class="date">{{.Post.Date}}</div>
            {{if .Post.Tags}}
            <div class="tags" style="margin: 1rem 0;">
                {{range .Post.Tags}}
                <span class="tag">{{.}}</span>
                {{end}}
            </div>
            {{end}}
            <div>
                {{.Post.Content}}
            </div>
            <div style="margin-top: 3rem;">
                <a href="../index.html#blog" class="btn">← Back to Blog</a>
            </div>
        </article>
    </main>
</body>
</html>