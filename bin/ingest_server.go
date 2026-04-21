package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		filename := fmt.Sprintf("/home/sam/opencode-obsidian/wiki/inbox/memo-%d.md", time.Now().Unix())
		content := fmt.Sprintf("---\ntype: memo\ncreated: %s\nsource: telegram\n---\n\n%s", time.Now().Format(time.RFC3339), string(body))

		err = ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Memo saved: %s\n", filename)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Saved successfully")
	})

	fmt.Println("Bunker Ingest Server running on :9090")
	http.ListenAndServe(":9090", nil)
}
