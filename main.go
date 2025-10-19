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

type TemplateFile struct {
	name      string
	root      string
	extension string
}

type OutputFile struct {
	name      string
	root      string
	dir       string
	extension string
}

var styleTpl = TemplateFile{name: "style", root: "template/", extension: ".css"}
var indexTpl = TemplateFile{name: "index", root: "template/", extension: ".html"}
var postTpl = TemplateFile{name: "post", root: "template/", extension: ".html"}
var styleArt = OutputFile{name: "style", root: "artifact/", dir: "css/", extension: ".css"}
var indexArt = OutputFile{name: "index", root: "artifact/", dir: "", extension: ".html"}
var blogDoc = OutputFile{name: "", root: "artifact/", dir: "blog/", extension: ".html"}

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
	os.MkdirAll(indexArt.root, 0755)
	os.MkdirAll(indexArt.root+indexArt.dir, 0755)
	os.MkdirAll(styleArt.root+styleArt.dir, 0755)

	// 設定ファイルの読み込み
	config, err := loadConfig("config.yaml")
	if err != nil {
	}

	// ブログ記事の読み込み
	posts, _ := loadPosts("content/blog")

	generateCSS(styleTpl, styleArt)

	generateIndexPage(indexTpl, indexArt, config, posts)
	generateBlogPages(config, posts)
}

func startServer(port string) {
	if _, err := os.Stat("artifact"); os.IsNotExist(err) {
		buildSite()
	}

	fs := http.FileServer(http.Dir("artifact"))
	http.Handle("/", fs)

	addr := ":" + port

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

func generateCSS(input TemplateFile, output OutputFile) {
	tmpl, err := template.ParseFiles(input.root + input.name + input.extension + ".tpl")
	if err != nil {
		log.Fatalf("Error parsing %s: %v", input.root+input.name+input.extension+".tpl", err)
	}

	f, err := os.Create(output.root + output.dir + output.name + output.extension)
	if err != nil {
		log.Fatalf("Error creating %s: %v", output.root+output.dir+output.name+output.extension+".tpl", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, nil); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func generateIndexPage(input TemplateFile, output OutputFile, config Config, posts []Post) {
	tmpl, err := template.ParseFiles(input.root + input.name + input.extension + ".tpl")
	if err != nil {
		log.Fatalf("Error parsing index.html.tpl: %v", err)
	}

	f, err := os.Create(output.root + output.dir + output.name + output.extension)
	if err != nil {
		log.Fatalf("Error creating %s: %v", output.root+output.dir+output.name+output.extension+".tpl", err)
	}
	defer f.Close()

	data := PageData{
		Config: config,
		Posts:  posts,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func generateBlogPages(config Config, posts []Post) {
	tmpl, err := template.ParseFiles("template/post.html.tpl")
	if err != nil {
		log.Fatalf("Error parsing post.html.tpl: %v", err)
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

		err = tmpl.Execute(f, data)
		if err != nil {
			log.Printf("Error executing template for post %s: %v", post.Slug, err)
		}
		f.Close()
	}
}
