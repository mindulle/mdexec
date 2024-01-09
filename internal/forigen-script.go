package internal

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func getRunTimeName(relPath string) string {
	fileExt := path.Ext(relPath)

	switch fileExt {
	case ".js":
		return "node"
	case ".py":
		return "python3"
	default:
		return "bash"
	}
}

func getDest(relPath string) string {
	return BaseDir + "/" + relPath
}

func RunForigenScriptFrom(relPath string, trailingArgs ...string) {
	cmdSlice := append([]string{getDest(relPath)}, trailingArgs...)

	runtimeStr := getRunTimeName(relPath)
	cmd := exec.Command(runtimeStr, cmdSlice...)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatal(err, "\nRun this command again with --help option for help.")
	}
}
