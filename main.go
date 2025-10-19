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

	parts := bytes.SplitN(data, []byte("---"), 3)
	if len(parts) < 3 {
		return Post{}, fmt.Errorf("invalid post format")
	}

	var post Post
	err = yaml.Unmarshal(parts[1], &post)
	if err != nil {
		return Post{}, err
	}

	post.ParsedDate, _ = time.Parse("2006-01-02", post.Date)

	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert(parts[2], &buf); err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(buf.String())

	return post, nil
}

func generateCSS() {
	tmpl, err := template.ParseFiles("template/style.css.tpl")
	if err != nil {
		log.Fatalf("Error parsing css.tpl: %v", err)
	}

	f, err := os.Create("docs/css/style.css")
	if err != nil {
		log.Fatalf("Error creating style.css: %v", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, nil); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func generateIndexPage(config Config, posts []Post) {
	tmpl, err := template.ParseFiles("template/index.html.tpl")
	if err != nil {
		log.Fatalf("Error parsing index.html.tpl: %v", err)
	}

	f, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatalf("Error creating index.html: %v", err)
	}
	defer f.Close()

	data := PageData{
		Config: config,
		Posts:  posts,
	}

	err = tmpl.Execute(f, data)
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
