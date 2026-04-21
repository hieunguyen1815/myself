package handler

import (
	"io"
	"net/http"
	"strings"
)

const mediumBase = "https://medium.com"
const proxyPrefix = "/api/medium"

func Proxy(w http.ResponseWriter, r *http.Request) {
	// Vercel injects catch-all path segments as the "path" query param.
	// e.g. /api/medium/foo/bar  ->  path=foo/bar
	mediumPath := r.URL.Query().Get("path")

	targetURL := mediumBase + "/" + mediumPath

	// Forward any original query params (minus the injected "path" one).
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

	// Mimic a browser so Medium returns full HTML rather than a bot page.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	// Disable compression so we can rewrite the body as plain text.
	req.Header.Set("Accept-Encoding", "identity")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "upstream fetch failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read upstream response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	html := string(body)

	// Rewrite all absolute Medium links so they flow through this proxy.
	html = strings.ReplaceAll(html, "https://medium.com", proxyPrefix)
	html = strings.ReplaceAll(html, "http://medium.com", proxyPrefix)

	// Copy safe response headers; skip hop-by-hop and content-length
	// (body length changed after rewriting).
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
