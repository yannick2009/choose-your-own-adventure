package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	//go:embed assets/stories.json
	jsonData []byte
	//go:embed views/index.html
	htmlData string
)

// Chapter represents a chapter in the story.
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option represents a choice in the story.
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// Story is a map of chapter names to Chapter structs.
type Story = map[string]Chapter

func main() {
	port := flag.Uint("port", 3000, "port to start the server on")
	path := flag.String("path", "", "path to JSON file containing the stories")
	flag.Parse()

	var data Story
	if strings.TrimSpace(*path) != "" {
		rawJSON, err := os.ReadFile(*path)
		if err != nil {
			panic(err)
		}
		data = loadStories(rawJSON)
	} else {
		data = loadStories(nil)
	}

	tmpl := template.Must(template.New("html").Parse(htmlData))

	handler := NewHandler(tmpl, data)

	http.HandleFunc("/", handler.ServeChapter)

	fmt.Println(fmt.Sprintf("Server is running: http://localhost:%d", *port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

// loadStories loads the story data from a JSON byte slice.
// If the data is nil, it uses the embedded JSON data. It returns a map of chapter names to Chapter structs.
func loadStories(data []byte) Story {
	histories := make(Story)

	if data == nil {
		data = jsonData
	}

	err := json.Unmarshal(data, &histories)
	if err != nil {
		panic(err)
	}

	return histories
}

// Handler is a struct that holds the template and story data.
type Handler struct {
	Templ *template.Template
	Data  Story
}

// NewHandler creates a new Handler with the given template and story data.
func NewHandler(templ *template.Template, data Story) *Handler {
	return &Handler{
		Templ: templ,
		Data:  data,
	}
}

// Redirect handles the root URL and serves the story chapters.
func (h *Handler) ServeChapter(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "" {
		path = "intro"
	}
	if chapter, ok := h.Data[path]; ok {
		if err := h.Templ.Execute(w, chapter); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		return
	}
}
