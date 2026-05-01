<!DOCTYPE html>
<html lang="ja">
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
                <div class="nav-links">
                    {{if .Config.GitHub}}<a href="{{.Config.GitHub}}" target="_blank">GitHub</a>{{end}}
                    {{if .Config.Bluesky}}<a href="{{.Config.Bluesky}}" target="_blank">Bluesky</a>{{end}}
                </div>
            </nav>
        </header>

        <div class="post-header">
            <a href="../index.html" class="post-back">← Writing</a>
            <h1>{{.Post.Title}}</h1>
            <p class="meta">{{.Post.Date}}{{if .Post.Tags}} · {{range $i, $v := .Post.Tags}}{{if $i}}, {{end}}{{$v}}{{end}}{{end}}</p>
        </div>

        <div class="post-body">
            {{.Post.Content}}
        </div>

        <footer>
            <p>{{.Config.Name}}</p>
        </footer>
    </div>
</body>
</html>
