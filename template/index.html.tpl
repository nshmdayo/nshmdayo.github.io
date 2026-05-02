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
            </nav>
            <div class="tabs">
                <button class="tab-btn active" data-tab="about">About</button>
                <button class="tab-btn" data-tab="blog">Blog</button>
            </div>
        </header>

        <main>
            <div class="tab-panel active" id="tab-about">
                <section class="intro">
                    <h1>{{.Config.Name}}</h1>
                    <p>{{.Config.Description}}</p>
                    {{if .Config.About}}<p class="about">{{.Config.About}}</p>{{end}}
                </section>
{{if .Config.Projects}}
                <section>
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
            </div>

            <div class="tab-panel" id="tab-blog">
            {{if .Posts}}
                <section>
                    <ul class="post-list">
                        {{range .Posts}}
                        <li class="post-item">
                            <a href="blog/{{.Slug}}.html">{{.Title}}</a>
                            <span class="post-date">{{.Date}}</span>
                        </li>
                        {{end}}
                    </ul>
                </section>
            {{else}}
                <p class="empty-state">No posts yet.</p>
            {{end}}
            </div>
        </main>

        <footer>
            <p>{{.Config.Name}}</p>
        </footer>
    </div>

    <script>
        const btns = document.querySelectorAll('.tab-btn');
        const panels = document.querySelectorAll('.tab-panel');

        btns.forEach(btn => {
            btn.addEventListener('click', () => {
                btns.forEach(b => b.classList.remove('active'));
                panels.forEach(p => p.classList.remove('active'));
                btn.classList.add('active');
                document.getElementById('tab-' + btn.dataset.tab).classList.add('active');
            });
        });
    </script>
</body>
</html>
