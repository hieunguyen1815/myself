package handler

import (
	"bytes"
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed assets/hieu_profile.md
var profileMarkdown []byte

func Handler(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	if _, err := body.Write(profileMarkdown); err != nil {
		http.Error(w, "could not read profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, body.String())
}
