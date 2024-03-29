package executer

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Executer struct {
}

func (e *Executer) GoCodeExec(request []byte, reply *[]byte) error {
	fmt.Println("using go executer", string(request))

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tempDir, err := os.MkdirTemp("", "gosandbox")
	if err != nil {
		// http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return err
	}
	defer os.RemoveAll(tempDir) // Clean up the temporary directory

	filePath := filepath.Join(tempDir, "main.go")
	if err := os.WriteFile(filePath, request, 0644); err != nil {
		// http.Error(w, "Error writing code to file", http.StatusInternalServerError)
		return err
	}

	// Run the Go code
	output, err := runGoCode(ctx, tempDir)
	if err != nil {

		return err
	}
	fmt.Fprintln(os.Stdout, []any{"err ", err}...)

	*reply = output
	return nil

}

func runGoCode(ctx context.Context, directory string) ([]byte, error) {
	cmd := exec.Command("go", "run", filepath.Join(directory, "main.go"))

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("code execution time out")
		}
		return nil, err
	}

	return output, nil
}
