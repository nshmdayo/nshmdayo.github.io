<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Config.Name}}</title>
    <meta name="description" content="{{.Config.Description}}">
    <link rel="stylesheet" href="css/style.css">
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

        <main>
            <section class="intro">
                <h1>{{.Config.Name}}</h1>
                <p>{{.Config.Description}}</p>
                {{if .Config.About}}<p class="about">{{.Config.About}}</p>{{end}}
            </section>

{{if .Config.Projects}}
            <section>
                <h2>Projects</h2>
                <ul class="project-list">
                    {{range .Config.Projects}}
                    <li class="project-item">
                        <div class="project-title-row">
                            <span class="project-name">{{.Title}}</span>
                            <span class="project-links">
                                {{if .GitHub}}<a href="{{.GitHub}}" target="_blank">↗ GitHub</a>{{end}}
                                {{if .URL}}<a href="{{.URL}}" target="_blank">↗ Live</a>{{end}}
                            </span>
                        </div>
                        <p class="project-desc">{{.Description}}</p>
                        <p class="project-tech">{{range $i, $v := .Tech}}{{if $i}} · {{end}}{{$v}}{{end}}</p>
                    </li>
                    {{end}}
                </ul>
            </section>
            {{end}}

            {{if .Posts}}
            <section>
                <h2>Writing</h2>
                <ul class="post-list">
                    {{range .Posts}}
                    <li class="post-item">
                        <a href="blog/{{.Slug}}.html">{{.Title}}</a>
                        <span class="post-date">{{.Date}}</span>
                    </li>
                    {{end}}
                </ul>
            </section>
            {{end}}
        </main>

        <footer>
            <p>{{.Config.Name}}</p>
        </footer>
    </div>
</body>
</html>
