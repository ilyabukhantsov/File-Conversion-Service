package conventor

import (
	"fmt"
	"os"
	"os/exec"
)

func Convert(id int, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found")
		}
		return fmt.Errorf("other error: %w", err)
	}
	cmd := exec.Command("echo", "Hello World")

	out, err := cmd.Output()
	fmt.Println(string(out), string(info.Name()))
	return nil
}
