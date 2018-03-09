package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func createProject(projectname string) error {
	filename := ntof(projectname)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating new project file: %q", err)
	}
	defer file.Close()

	file.Chmod(0774) // we need execute permission
	file.WriteString(config.Template)
	file.Sync()

	return exec.Command(config.EditCommand, filename).Run()
}

func editProject(projectname string) error {
	filename := ntof(projectname)
	if projectExists(projectname) {
		return exec.Command(config.EditCommand, filename).Run()
	}
	return errors.New("project doesn't exist")
}

func listProjects() error {
	files, err := listProjectFiles()

	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	return nil
}

func repairProjects(projects []string) error {
	for _, project := range projects {
		if err := os.Chmod(ntof(project), 0774); err != nil {
			return err
		}
	}
	return nil
}

func deleteProject(projectname string) error {
	reply := input(fmt.Sprintf("delete project '%s'? (y/n)", projectname))
	if strings.Trim(reply, " \n\t") != "y" {
		println("abort")
		return nil
	}
	if err := os.Remove(ntof(projectname)); err != nil {
		return err
	}
	fmt.Printf("project '%s' deleted.\n", projectname)
	return nil
}
