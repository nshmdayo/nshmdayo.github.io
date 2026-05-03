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
	serve := flag.Bool("serve", false, "start local server for preview")
	port := flag.String("port", "8000", "server port")
	flag.Parse()

	if *serve {
		startServer(*port)
		return
	}

	buildSite()
}

func buildSite() {
	for _, dir := range []string{"artifact/css", "artifact/blog"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("mkdir %s: %v", dir, err)
		}
	}

	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	posts, _ := loadPosts("content/post")

	if err := renderTemplate("template/style.css.tpl", "artifact/css/style.css", nil); err != nil {
		log.Fatalf("render style: %v", err)
	}
	if err := renderTemplate("template/index.html.tpl", "artifact/index.html", PageData{Config: config, Posts: posts}); err != nil {
		log.Fatalf("render index: %v", err)
	}
	generateBlogPages(config, posts)
}

func startServer(port string) {
	if _, err := os.Stat("artifact"); os.IsNotExist(err) {
		buildSite()
	}
	http.Handle("/", http.FileServer(http.Dir("artifact")))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(tplPath, outPath string, data any) error {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, data)
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
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var posts []Post
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
	if err := yaml.Unmarshal(parts[1], &post); err != nil {
		return Post{}, err
	}

	post.ParsedDate, _ = time.Parse("2006-01-02", post.Date)

	var buf bytes.Buffer
	if err := goldmark.New().Convert(parts[2], &buf); err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(buf.String())

	return post, nil
}

func generateBlogPages(config Config, posts []Post) {
	tmpl, err := template.ParseFiles("template/post.html.tpl")
	if err != nil {
		log.Fatalf("Error parsing post.html.tpl: %v", err)
	}

	for _, post := range posts {
		f, err := os.Create(filepath.Join("artifact/blog", post.Slug+".html"))
		if err != nil {
			log.Printf("Error creating post %s: %v", post.Slug, err)
			continue
		}
		if err := tmpl.Execute(f, PageData{Config: config, Post: &post}); err != nil {
			log.Printf("Error executing template for post %s: %v", post.Slug, err)
		}
		f.Close()
	}
}
