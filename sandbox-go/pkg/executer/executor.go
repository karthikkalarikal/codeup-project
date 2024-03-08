package executer

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Executer struct {
}

func (e *Executer) GoCodeExec(request []byte, reply *[]byte) error {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// Read the submitted Go code
	// code, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Error reading code", http.StatusInternalServerError)
	// 	return
	// }

	// Create a temporary directory
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tempDir, err := os.MkdirTemp("", "gosandbox")
	if err != nil {
		// http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return err
	}
	defer os.RemoveAll(tempDir) // Clean up the temporary directory

	// Write the Go code to a temporary file
	filePath := filepath.Join(tempDir, "main.go")
	if err := os.WriteFile(filePath, request, 0644); err != nil {
		// http.Error(w, "Error writing code to file", http.StatusInternalServerError)
		return err
	}

	// Run the Go code
	output, err := runGoCode(ctx, tempDir)
	if err != nil {
		// http.Error(w, fmt.Sprintf("Error running code: %s", err), http.StatusInternalServerError)
		return err
	}

	*reply = output
	return nil
	// Send the output back to the user
	// w.Header().Set("Content-Type", "text/plain")
	// w.Write(output)
}

func runGoCode(ctx context.Context, directory string) ([]byte, error) {
	cmd := exec.Command("go", "run", filepath.Join(directory, "main.go"))

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
