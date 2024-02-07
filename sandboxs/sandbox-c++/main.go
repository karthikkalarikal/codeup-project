// main.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func executeCppCode(code string) (string, error) {
	tmpfile, err := os.CreateTemp("", "cpp_code_*.cpp")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(code); err != nil {
		return "", err
	}

	cmd := exec.Command("g++", tmpfile.Name(), "-o", "/app/cpp_executable")
	compileOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Compilation error", err)
		fmt.Println("compilation output", string(compileOutput))
		return string(compileOutput), err
	}

	cmd = exec.Command("/app/cpp_executable")
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	code := string(body)

	output, err := executeCppCode(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing C++ code: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, output)
}

func main() {
	http.HandleFunc("/execute-cpp", codeExecutionHandler)
	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
