package main

import (
	"bytes"
	"crypto/subtle"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const maxBodyBytes = 10 << 20 // 10 MiB

func envDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func unauthorized(w http.ResponseWriter) {
	http.Error(w, "unauthorized", http.StatusUnauthorized)
}

func requireToken(w http.ResponseWriter, r *http.Request) bool {
	token := os.Getenv("BUNKER_INGEST_TOKEN")
	if token == "" {
		return true
	}
	provided := r.Header.Get("X-Bunker-Token")
	if provided == "" {
		provided = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	}
	if subtle.ConstantTimeCompare([]byte(provided), []byte(token)) != 1 {
		unauthorized(w)
		return false
	}
	return true
}

func safeSlug(raw string) string {
	words := strings.Fields(raw)
	if len(words) == 0 {
		return fmt.Sprintf("memo-%d", time.Now().Unix())
	}
	limit := 5
	if len(words) < limit {
		limit = len(words)
	}
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := strings.Trim(reg.ReplaceAllString(strings.ToLower(strings.Join(words[:limit], "-")), "-"), "-")
	if slug == "" {
		return fmt.Sprintf("memo-%d", time.Now().Unix())
	}
	if len(slug) > 80 {
		slug = slug[:80]
	}
	return slug
}

func categorize(raw string) (tag string, isArchitecture bool) {
	lower := strings.ToLower(raw)
	switch {
	case strings.Contains(lower, "error"), strings.Contains(lower, "bug"), strings.Contains(lower, "fallo"), strings.Contains(lower, "falla"):
		return "bug", false
	case strings.Contains(lower, "decidimos"), strings.Contains(lower, "estándar"), strings.Contains(lower, "estandar"), strings.Contains(lower, "arquitectura"), strings.Contains(lower, "regla"):
		return "decision", true
	case strings.Contains(lower, "todo"), strings.Contains(lower, "tarea"), strings.Contains(lower, "revisar"):
		return "task", false
	default:
		return "memo", false
	}
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error getting cwd: %v\n", err)
		os.Exit(1)
	}

	bind := envDefault("BUNKER_INGEST_BIND", "127.0.0.1:9090")
	inboxDir := envDefault("BUNKER_INBOX_DIR", filepath.Join(cwd, "wiki", "inbox"))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok","service":"bunker-ingest"}` + "\n"))
	})

	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !requireToken(w, r) {
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "body too large or read failed", http.StatusRequestEntityTooLarge)
			return
		}
		rawContent := strings.TrimSpace(string(body))
		if rawContent == "" {
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		tag, isArchitecture := categorize(rawContent)
		lowerContent := strings.ToLower(rawContent)
		titleSlug := safeSlug(rawContent)

		complianceReport := ""
		if isArchitecture && strings.Contains(lowerContent, "ioutil") {
			cmd := exec.Command("grep", "-r", "ioutil", "bin/")
			var out bytes.Buffer
			cmd.Stdout = &out
			_ = cmd.Run()
			if strings.TrimSpace(out.String()) != "" {
				complianceReport = "\n\n## Bunker Sentinel: Compliance Report\nNon-compliant pattern detected: `ioutil`.\n"
			}
		}

		if err := os.MkdirAll(inboxDir, 0o755); err != nil {
			http.Error(w, "failed to create inbox", http.StatusInternalServerError)
			return
		}

		filename := filepath.Join(inboxDir, fmt.Sprintf("%s.md", titleSlug))
		frontmatter := fmt.Sprintf("---\ntype: %s\ncreated: %s\nsource: smart-ingest-local\ntags: [%s]\n---\n\n", tag, time.Now().Format(time.RFC3339), tag)
		fullContent := frontmatter + rawContent + complianceReport + "\n"

		if err := os.WriteFile(filename, []byte(fullContent), 0o644); err != nil {
			http.Error(w, "failed to write note", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "saved as %s\n", filepath.Base(filename))
	})

	fmt.Printf("Bunker Smart-Ingest Server running on http://%s\n", bind)
	fmt.Printf("Inbox: %s\n", inboxDir)
	if os.Getenv("BUNKER_INGEST_TOKEN") == "" {
		fmt.Println("Warning: BUNKER_INGEST_TOKEN is not set; local clients do not need a token.")
	}
	if err := http.ListenAndServe(bind, nil); err != nil {
		fmt.Printf("server error: %v\n", err)
		os.Exit(1)
	}
}
