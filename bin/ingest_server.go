package main

import (
	"bytes"
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

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		rawContent := string(body)

		// --- CATEGORIZACIÓN ---
		tag := "memo"
		isArchitecture := false
		lowerContent := strings.ToLower(rawContent)
		if strings.Contains(lowerContent, "error") || strings.Contains(lowerContent, "bug") {
			tag = "bug"
		} else if strings.Contains(lowerContent, "decidimos") || strings.Contains(lowerContent, "estándar") || strings.Contains(lowerContent, "architecture") {
			tag = "decision"
			isArchitecture = true
		} else if strings.Contains(lowerContent, "tarea") || strings.Contains(lowerContent, "revisar") {
			tag = "task"
		}

		// --- SENTINEL: Auditoría Reactiva (v3.5) ---
		complianceReport := ""
		if isArchitecture {
			complianceReport = "\n\n## 🛡️ Bunker Sentinel: Compliance Report\n"
			
			// Ejemplo: Si la decisión menciona "ioutil", buscamos violaciones
			if strings.Contains(lowerContent, "ioutil") {
				cmd := exec.Command("grep", "-r", "ioutil", "bin/", "src/")
				var out bytes.Buffer
				cmd.Stdout = &out
				cmd.Run() // No importa si falla (si no encuentra nada)
				
				violations := strings.TrimSpace(out.String())
				if violations == "" {
					complianceReport += "✅ **Estado: Compliant.** No se encontraron usos de `ioutil` en el código actual.\n"
				} else {
					complianceReport += "⚠️ **Estado: Non-Compliant.** Se detectaron usos de `ioutil` que deben ser refactorizados:\n"
					lines := strings.Split(violations, "\n")
					for i, line := range lines {
						if i >= 5 { break } // Limitar a 5 para no saturar
						complianceReport += fmt.Sprintf("- `%s`\n", line)
					}
				}
			} else {
				complianceReport += "ℹ️ **Estado: Sin auditoría automática para este patrón.** Guardado como norma general.\n"
			}
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

		// --- ENRIQUECIMIENTO MCP ---
		contextInfo := ""
		// (Mantener lógica de smart-search anterior...)
		
		// --- ESCRITURA FINAL ---
		filename := filepath.Join(cwd, "wiki", "inbox", fmt.Sprintf("%s.md", titleSlug))
		os.MkdirAll(filepath.Dir(filename), 0755)

		frontmatter := fmt.Sprintf("---\ntype: %s\ncreated: %s\nsource: smart-ingest-v3.5\ntags: [%s]\n---\n\n", tag, time.Now().Format(time.RFC3339), tag)
		fullContent := frontmatter + rawContent + complianceReport + contextInfo

		os.WriteFile(filename, []byte(fullContent), 0644)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Saved as %s with Sentinel Audit\n", titleSlug)
	})

	fmt.Println("Bunker Smart-Ingest Server v3.5 (SENTINEL) running on :9090")
	http.ListenAndServe(":9090", nil)
}
