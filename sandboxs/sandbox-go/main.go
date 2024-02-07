package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the submitted Go code
	code, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading code", http.StatusInternalServerError)
		return
	}

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "gosandbox")
	if err != nil {
		http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir) // Clean up the temporary directory

	// Write the Go code to a temporary file
	filePath := filepath.Join(tempDir, "main.go")
	if err := os.WriteFile(filePath, code, 0644); err != nil {
		http.Error(w, "Error writing code to file", http.StatusInternalServerError)
		return
	}

	// Run the Go code
	output, err := runGoCode(tempDir)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error running code: %s", err), http.StatusInternalServerError)
		return
	}

	// Send the output back to the user
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func runGoCode(directory string) ([]byte, error) {
	cmd := exec.Command("go", "run", filepath.Join(directory, "main.go"))

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
