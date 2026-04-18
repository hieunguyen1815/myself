package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
)

// Markdown serves a markdown file rendered as HTML at the root path.
// It reads the file on every request so changes on disk are reflected immediately.
type Markdown struct {
	filepath string
}

func New(filepath string) *Markdown {
	return &Markdown{filepath: filepath}
}

func (m *Markdown) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", m.serveMarkdown)
	mux.HandleFunc("GET /health", serveHealth)
}

func (m *Markdown) serveMarkdown(w http.ResponseWriter, r *http.Request) {
	src, err := os.ReadFile(m.filepath)
	if err != nil {
		http.Error(w, "could not read file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var body bytes.Buffer
	if err := goldmark.Convert(src, &body); err != nil {
		http.Error(w, "could not render markdown: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, htmlPage, body.String())
}

func serveHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

const htmlPage = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <style>
    body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; max-width: 800px; margin: 40px auto; padding: 0 20px; line-height: 1.6; color: #333; }
    h1, h2, h3 { border-bottom: 1px solid #eee; padding-bottom: 8px; }
    code { background: #f4f4f4; padding: 2px 6px; border-radius: 4px; font-size: 0.9em; }
    pre { background: #f4f4f4; padding: 16px; border-radius: 6px; overflow-x: auto; }
    pre code { background: none; padding: 0; }
    a { color: #0070f3; }
    blockquote { border-left: 4px solid #ddd; margin: 0; padding-left: 16px; color: #666; }
  </style>
</head>
<body>
%s
</body>
</html>`
