package main

//TODO: add ability to archive projects
// - pass arguments to projects
// archive
// delete project

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"

	"github.com/urfave/cli"
)

var config = struct {
	ProjectFileDir string
	EditCommand    string
	Extension      string
	Template       string
}{}

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	// DEFAULT CONFIG
	config.EditCommand = "code"
	config.ProjectFileDir = path.Join(usr.HomeDir, "/.pro")
	config.Extension = "sh"
	config.Template = "#!/bin/bash\n"

	app := cli.NewApp()
	app.Name = "pro"
	app.Usage = "Manage your projects and their environments"
	app.Version = "0.0.1"
	// default action, which is to run the project file if it exists
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			return nil
		}

		projectname := c.Args().First()

		if projectExists(projectname) {
			cmd := exec.Command(ntof(projectname))
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Run()
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:        "new",
			Aliases:     []string{"n"},
			Description: "create a new project",
			Usage:       "pro new my-project",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					println(c.Command.UsageText)
					return nil
				}
				return createProject(c.Args().First())
			},
		},
		{
			Name:        "edit",
			Aliases:     []string{"e"},
			Description: "edit a project file",
			Usage:       "edit my-project",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					println(c.Command.UsageText)
					return nil
				}
				return editProject(c.Args().First())
			},
		},
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Description: "list all available projects",
			Usage:       "list",
			Action: func(c *cli.Context) error {
				return listProjects()
			},
		},
		{
			Name:        "repair",
			Aliases:     []string{"r"},
			Description: "attempt to repair project file permissions",
			Usage:       "repair [optional-file]",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					files, err := listProjectFiles()
					if err != nil {
						return err
					}

					allProjects := []string{}
					for _, f := range files {
						allProjects = append(allProjects, f.Name())
					}

					return repairProjects(allProjects)
				}
				return repairProjects(c.Args())
			},
		},
		{
			Name:        "delete",
			Aliases:     []string{"d"},
			Description: "delete a project file",
			Usage:       "delete my-project",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return nil
				}
				return deleteProject(c.Args().First())
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
