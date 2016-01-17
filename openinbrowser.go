package openinbrowser

import (
	"log"
	"os/exec"
	"runtime"
)

func Open(path string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path}
	case "windows":
		args = []string{"cmd", "/c", "start", path}
	default:
		args = []string{"xdg-open", path}
	}
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Printf("openinbrowser: %v\n", err)
	}
}
