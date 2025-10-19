<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Not a Shell</title>
    <meta name="description" content="{{.Config.Description}}">
    <link rel="stylesheet" href="css/style.css">
</head>
<body>
    <div class="hero">
        <div class="container" style="text-align: left;">
            <h2>Hello</h2>
            <h2>I'm {{.Config.Name}},</h2>
            <h2>not the command-line kind.</h2>
            <p>{{.Config.Description}}</p>
        </div>
    </div>

    <main class="container">
        <div class="social-links">
            {{if .Config.GitHub}}<a href="{{.Config.GitHub}}" target="_blank">GitHub</a>{{end}}
            {{if .Config.Bluesky}}<a href="{{.Config.Bluesky}}" target="_blank">Bluesky</a>{{end}}
        </div>

        {{if .Posts}}
        <section id="blog">
            <h2>Latest Blog Posts</h2>
            <div class="grid">
                {{range .Posts}}
                <div class="card">
                    <h3><a href="blog/{{.Slug}}.html">{{.Title}}</a></h3>
                    <div class="date">{{.Date}}</div>
                    <p>{{.Description}}</p>
                    {{if .Tags}}
                    <div class="tags">
                        {{range .Tags}}
                        <span class="tag">{{.}}</span>
                        {{end}}
                    </div>
                    {{end}}
                    <a href="blog/{{.Slug}}.html" class="btn">Read More</a>
                </div>
                {{end}}
            </div>
        </section>
        {{end}}
    </main>
</body>
</html>