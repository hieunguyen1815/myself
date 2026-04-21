package handler

import (
	"compress/gzip"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const mediumBase = "https://medium.com"
const proxyPrefix = "/api/proxy"

// Matches root-relative paths like href="/foo" but not protocol-relative href="//cdn.x.com"
var relPathRe = regexp.MustCompile(`(href|src|action)="(/[^/][^"]*)`)

func Proxy(w http.ResponseWriter, r *http.Request) {
	mediumPath := r.URL.Query().Get("path")
	if mediumPath == "" {
		http.Error(w, "missing 'path' query parameter", http.StatusBadRequest)
		return
	}

	targetURL := mediumBase + "/" + mediumPath

	// Forward original query params (minus the injected "path" one).
	q := r.URL.Query()
	q.Del("path")
	if qs := q.Encode(); qs != "" {
		targetURL += "?" + qs
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, targetURL, nil)
	if err != nil {
		http.Error(w, "bad target URL: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Mimic a real browser as closely as possible.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Referer", "https://www.google.com/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// Forward cookies from the browser (e.g. cf_clearance) to Medium.
	if cookie := r.Header.Get("Cookie"); cookie != "" {
		req.Header.Set("Cookie", cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "upstream fetch failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Decompress gzip if Medium sent it.
	bodyReader := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			http.Error(w, "gzip decode failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer gr.Close()
		bodyReader = gr
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		http.Error(w, "failed to read upstream response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	html := string(body)

	// 1. Rewrite absolute Medium URLs.
	html = strings.ReplaceAll(html, "https://medium.com", proxyPrefix)
	html = strings.ReplaceAll(html, "http://medium.com", proxyPrefix)

	// 2. Rewrite root-relative paths (e.g. /@user/article, /tag/go).
	//    Skips protocol-relative URLs like //cdn.example.com.
	html = relPathRe.ReplaceAllString(html, `$1="/api/proxy$2`)

	// Copy safe response headers; drop hop-by-hop and encoding headers
	// (body is now decompressed and length has changed).
	skipHeaders := map[string]bool{
		"Content-Length":    true,
		"Transfer-Encoding": true,
		"Content-Encoding":  true,
		"Connection":        true,
	}
	for k, vv := range resp.Header {
		if skipHeaders[k] {
			continue
		}
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, _ = io.WriteString(w, html)
}
