package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

type Post struct {
	Title       string   `yaml:"title"`
	Date        string   `yaml:"date"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Slug        string
	Content     template.HTML
	ParsedDate  time.Time
}

type Config struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	GitHub      string `yaml:"github"`
	Bluesky     string `yaml:"Bluesky"`
}

type PageData struct {
	Config Config
	Posts  []Post
	Post   *Post
}

func main() {
	// コマンドラインフラグの定義
	serve := flag.Bool("serve", false, "ローカルサーバーを起動してプレビュー")
	port := flag.String("port", "8000", "サーバーのポート番号")
	flag.Parse()

	// サーバーモードの場合
	if *serve {
		startServer(*port)
		return
	}

	buildSite()
}

func buildSite() {
	// ディレクトリの作成
	os.MkdirAll("docs", 0755)
	os.MkdirAll("docs/blog", 0755)
	os.MkdirAll("docs/css", 0755)

	// 設定ファイルの読み込み
	config, err := loadConfig("config.yaml")
	if err != nil {
	}

	// ブログ記事の読み込み
	posts, _ := loadPosts("content/blog")

	// CSSファイルの生成
	generateCSS()

	// ページの生成
	generateIndexPage(config, posts)
	generateBlogPages(config, posts)
}

func startServer(port string) {
	// docsディレクトリが存在しない場合は先に生成
	if _, err := os.Stat("docs"); os.IsNotExist(err) {
		log.Println("📝 docsディレクトリが見つかりません。サイトを生成します...")
		buildSite()
		fmt.Println("✅ サイトが生成されました！")
		fmt.Println()
	}

	// 静的ファイルサーバーの設定
	fs := http.FileServer(http.Dir("docs"))
	http.Handle("/", fs)

	addr := ":" + port
	fmt.Println("🚀 ローカルサーバーを起動しました")
	fmt.Printf("📍 http://localhost%s\n", addr)
	fmt.Println("⏹  停止するには Ctrl+C を押してください")
	fmt.Println()

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func loadPosts(dir string) ([]Post, error) {
	var posts []Post

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return posts, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		post, err := parsePost(filepath.Join(dir, entry.Name()))
		if err != nil {
			log.Printf("Error parsing post %s: %v", entry.Name(), err)
			continue
		}

		post.Slug = strings.TrimSuffix(entry.Name(), ".md")
		posts = append(posts, post)
	}

	// 日付でソート（新しい順）
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ParsedDate.After(posts[j].ParsedDate)
	})

	return posts, nil
}

func parsePost(path string) (Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Post{}, err
	}

	// Front matterの解析
	parts := bytes.SplitN(data, []byte("---"), 3)
	if len(parts) < 3 {
		return Post{}, fmt.Errorf("invalid post format")
	}

	var post Post
	err = yaml.Unmarshal(parts[1], &post)
	if err != nil {
		return Post{}, err
	}

	// 日付のパース
	post.ParsedDate, _ = time.Parse("2006-01-02", post.Date)

	// Markdownのパース
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert(parts[2], &buf); err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(buf.String())

	return post, nil
}

func generateCSS() {
	css := `
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    line-height: 1.6;
    color: #333;
    background: #f8f9fa;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
}

nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

nav h1 {
    font-size: 1.5rem;
    color: #2563eb;
}

nav ul {
    display: flex;
    list-style: none;
    gap: 2rem;
}

nav a {
    color: #333;
    text-decoration: none;
    transition: color 0.3s;
}

nav a:hover {
    color: #2563eb;
}

.hero {
    text-align: center;
    padding: 4rem 0;
    margin-bottom: 3rem;
}

.hero h2 {
    font-size: 2.5rem;
    margin-bottom: 1rem;
}

.hero p {
    font-size: 1.2rem;
    opacity: 0.9;
}

section {
    margin-bottom: 4rem;
}

h2 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
    color: #1e293b;
}

.grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 2rem;
}

.card {
    background: white;
    border-radius: 8px;
    padding: 1.5rem;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    transition: transform 0.3s, box-shadow 0.3s;
}

.card:hover {
    transform: translateY(-4px);
    box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.card h3 {
    color: #2563eb;
    margin-bottom: 0.5rem;
}

.card .date {
    color: #64748b;
    font-size: 0.9rem;
    margin-bottom: 0.5rem;
}

.card p {
    color: #475569;
    margin-bottom: 1rem;
}

.tags {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
    margin-top: 1rem;
}

.tag {
    background: #e0e7ff;
    color: #3730a3;
    padding: 0.25rem 0.75rem;
    border-radius: 9999px;
    font-size: 0.875rem;
}

.btn {
    display: inline-block;
    padding: 0.5rem 1rem;
    background: #2563eb;
    color: white;
    text-decoration: none;
    border-radius: 4px;
    transition: background 0.3s;
}

.btn:hover {
    background: #1d4ed8;
}

.social-links {
    display: flex;
    gap: 1rem;
    justify-content: center;
    margin-top: 2rem;
}

.social-links a {
    color: #64748b;
    transition: color 0.3s;
}

.social-links a:hover {
    color: #2563eb;
}

article {
    background: white;
    padding: 3rem;
    border-radius: 8px;
    max-width: 800px;
    margin: 0 auto;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

article h1 {
    color: #1e293b;
    margin-bottom: 0.5rem;
}

article img {
    max-width: 100%;
    height: auto;
    border-radius: 4px;
    margin: 1rem 0;
}

article pre {
    background: #1e293b;
    color: #e2e8f0;
    padding: 1rem;
    border-radius: 4px;
    overflow-x: auto;
    margin: 1rem 0;
}

article code {
    background: #f1f5f9;
    padding: 0.2rem 0.4rem;
    border-radius: 4px;
    font-family: 'Courier New', monospace;
}

article pre code {
    background: none;
    padding: 0;
}

@media (max-width: 768px) {
    .hero h2 {
        font-size: 1.8rem;
    }
    
    nav ul {
        gap: 1rem;
    }
    
    .grid {
        grid-template-columns: 1fr;
    }
    
    article {
        padding: 1.5rem;
    }
}
`
	os.WriteFile("docs/css/style.css", []byte(css), 0644)
}

func generateIndexPage(config Config, posts []Post) {
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Config.Name}} - {{.Config.Title}}</title>
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
</html>`

	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := PageData{
		Config: config,
		Posts:  posts,
	}

	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}

func generateBlogPages(config Config, posts []Post) {
	tmpl := `<!DOCTYPE html>
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
</html>`

	t, err := template.New("post").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		f, err := os.Create(filepath.Join("docs/blog", post.Slug+".html"))
		if err != nil {
			log.Printf("Error creating post %s: %v", post.Slug, err)
			continue
		}

		data := PageData{
			Config: config,
			Post:   &post,
		}

		err = t.Execute(f, data)
		if err != nil {
			log.Printf("Error executing template for post %s: %v", post.Slug, err)
		}
		f.Close()
	}
}
