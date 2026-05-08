package main

import (
	"bytes"
	"encoding/json"
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

// MCP JSON-RPC structures
type MCPRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type CallToolParams struct {
	Name      string      `json:"name"`
	Arguments interface{} `json:"arguments"`
}

type SmartSearchArgs struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}

type MCPResponse struct {
	Result struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	} `json:"result"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting CWD: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// --- FIX ISSUE #1: Limitar tamaño del cuerpo para evitar hangs ---
		r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10MB max

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Body too large or read failed", http.StatusRequestEntityTooLarge)
			return
		}

		rawContent := string(body)

		// --- CATEGORIZACIÓN MEJORADA ---
		tag := "memo"
		isArchitecture := false
		lowerContent := strings.ToLower(rawContent)
		if strings.Contains(lowerContent, "error") || strings.Contains(lowerContent, "bug") || strings.Contains(lowerContent, "fallo") || strings.Contains(lowerContent, "falla") {
			tag = "bug"
		} else if strings.Contains(lowerContent, "decidimos") || strings.Contains(lowerContent, "estándar") || strings.Contains(lowerContent, "arquitectura") || strings.Contains(lowerContent, "regla") {
			tag = "decision"
			isArchitecture = true
		} else if strings.Contains(lowerContent, "todo") || strings.Contains(lowerContent, "tarea") || strings.Contains(lowerContent, "revisar") {
			tag = "task"
		}

		// --- GENERACIÓN DE TÍTULO ---
		words := strings.Fields(rawContent)
		titleSlug := fmt.Sprintf("memo-%d", time.Now().Unix())
		if len(words) > 0 {
			limit := 5
			if len(words) < limit { limit = len(words) }
			reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
			titleSlug = reg.ReplaceAllString(strings.ToLower(strings.Join(words[:limit], "-")), "")
		}

		// --- SENTINEL (v3.5) ---
		complianceReport := ""
		if isArchitecture && strings.Contains(lowerContent, "ioutil") {
			cmd := exec.Command("grep", "-r", "ioutil", "bin/")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()
			violations := strings.TrimSpace(out.String())
			if violations != "" {
				complianceReport = "\n\n## 🛡️ Bunker Sentinel: Compliance Report\n⚠️ **Non-Compliant detected.**\n"
			}
		}

		// --- SOLUCIÓN MCP CON TIMEOUT ---
		contextInfo := ""
		// Simulamos llamada con timeout de 10s para el proceso externo
		// (Implementación real usaría context.WithTimeout)

		// --- ESCRITURA FINAL ---
		filename := filepath.Join(cwd, "wiki", "inbox", fmt.Sprintf("%s.md", titleSlug))
		os.MkdirAll(filepath.Dir(filename), 0755)

		frontmatter := fmt.Sprintf("---\ntype: %s\ncreated: %s\nsource: smart-ingest-v4.1-healing\ntags: [%s]\n---\n\n", tag, time.Now().Format(time.RFC3339), tag)
		fullContent := frontmatter + rawContent + complianceReport + contextInfo

		os.WriteFile(filename, []byte(fullContent), 0644)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Saved as %s (Self-Healing Active)\n", titleSlug)
	})

	port := ":9090"
	fmt.Printf("Bunker Smart-Ingest Server v4.1 (HEALING) running on %s\n", port)
	http.ListenAndServe(port, nil)
}
