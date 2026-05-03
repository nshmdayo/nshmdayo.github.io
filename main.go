package main

import (
	"bytes"
	"errors"
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

const postDir = "content/post"

var md = goldmark.New()

type Post struct {
	Title       string   `yaml:"title"`
	Date        string   `yaml:"date"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Slug        string
	Content     template.HTML
	ParsedDate  time.Time
}

type Project struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tech        []string `yaml:"tech"`
	GitHub      string   `yaml:"github"`
	URL         string   `yaml:"url"`
}

type Config struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	About       string    `yaml:"about"`
	Projects    []Project `yaml:"projects"`
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
		buildSite()
		go watchAndRebuild()
		http.Handle("/", http.FileServer(http.Dir("artifact")))
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	} else {
		buildSite()
	}
}

func watchAndRebuild() {
	staticPaths := []string{
		"config.yaml",
		"template/index.html.tpl",
		"template/post.html.tpl",
		"template/style.css.tpl",
	}

	type fileState struct {
		modTime time.Time
		size    int64
	}

	snapshot := func() map[string]fileState {
		state := make(map[string]fileState)
		for _, p := range staticPaths {
			if info, err := os.Stat(p); err == nil {
				state[p] = fileState{info.ModTime(), info.Size()}
			}
		}
		entries, _ := os.ReadDir(postDir)
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
				if info, err := e.Info(); err == nil {
					state[filepath.Join(postDir, e.Name())] = fileState{info.ModTime(), info.Size()}
				}
			}
		}
		return state
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	prev := snapshot()
	for range ticker.C {
		curr := snapshot()
		rebuild := len(curr) != len(prev)
		if !rebuild {
			for p, cs := range curr {
				if prev[p] != cs {
					rebuild = true
					break
				}
			}
		}
		if rebuild {
			log.Println("Change detected, rebuilding...")
			buildSite()
			prev = curr
		}
	}
}

func buildSite() {
	for _, dir := range []string{"artifact/css", "artifact/post"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("mkdir %s: %v", dir, err)
		}
	}

	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	posts, err := loadPosts(postDir)
	if err != nil {
		log.Fatalf("load posts: %v", err)
	}

	if err := renderTemplate("template/style.css.tpl", "artifact/css/style.css", nil); err != nil {
		log.Fatalf("render style: %v", err)
	}
	if err := renderTemplate("template/index.html.tpl", "artifact/index.html", PageData{Config: config, Posts: posts}); err != nil {
		log.Fatalf("render index: %v", err)
	}

	generateBlogPages(config, posts)
}

func executeTemplate(tmpl *template.Template, outPath string, data any) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, data)
}

func renderTemplate(tplPath, outPath string, data any) error {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}
	return executeTemplate(tmpl, outPath, data)
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
	entries, err := os.ReadDir(dir)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
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
			log.Printf("skipping %s: %v", entry.Name(), err)
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

	var dateErr error
	post.ParsedDate, dateErr = time.Parse("2006-01-02", post.Date)
	if dateErr != nil {
		log.Printf("warning: %s: invalid date %q", path, post.Date)
	}

	var buf bytes.Buffer
	if err := md.Convert(parts[2], &buf); err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(buf.String())

	return post, nil
}

func generateBlogPages(config Config, posts []Post) {
	tmpl, err := template.ParseFiles("template/post.html.tpl")
	if err != nil {
		log.Fatalf("parse post template: %v", err)
	}

	for _, post := range posts {
		outPath := filepath.Join("artifact/post", post.Slug+".html")
		if err := executeTemplate(tmpl, outPath, PageData{Config: config, Post: &post}); err != nil {
			log.Printf("render post %s: %v", post.Slug, err)
		}
	}
}
