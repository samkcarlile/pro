package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func listProjectFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(config.ProjectFileDir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// n(ame)tof(ile)
func ntof(projectname string) string {
	return path.Join(config.ProjectFileDir, projectname)
}

func projectExists(projectname string) bool {
	if _, err := os.Stat(ntof(projectname)); err != nil {
		return false
	}
	return true
}

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	reply, _ := reader.ReadString('\n')
	return reply
}
