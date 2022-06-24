package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func main() {
	cErr := color.New(color.FgWhite).Add(color.BgRed)

	if len(os.Args) < 2 {
		cErr.Println("Usage: cani <search query>")
		os.Exit(1)
	}

	queryArray := os.Args[1:]

	query := ""
	for _, q := range queryArray {
		query += q + " "
	}

	fmt.Printf("Searching for %s in Can I use\n", query)

	openBrowser("https://caniuse.com/?search=" + query)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("run", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		os.Exit(1)
	}
}
