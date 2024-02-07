// main.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func executePythonCode(code string) (string, error) {
	// Create a temporary file to store the Python code
	tmpfile, err := os.CreateTemp("", "python_code_*.py")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name()) // Clean up the temporary file

	// Write the Python code to the temporary file
	if _, err := tmpfile.WriteString(code); err != nil {
		return "", err
	}

	// Command to run Python code
	cmd := exec.Command("python3", tmpfile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

func codeExecutionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the Python code from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Convert the request body (Python code) to a string
	code := string(body)

	// Execute the Python code
	output, err := executePythonCode(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing Python code: %v", err), http.StatusInternalServerError)
		return
	}

	// Send the output as the response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, output)
}

func main() {
	http.HandleFunc("/execute-python", codeExecutionHandler)
	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
