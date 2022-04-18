package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var dependabotFile, terragruntRepoPath string

func main() {

	flag.StringVar(&dependabotFile, "dependabotFile", "dependabot.yml", "Dependabot file path")
	flag.StringVar(&terragruntRepoPath, "terragruntRepoPath", "", "Terragrunt repo path")
	flag.Parse()

	if len(terragruntRepoPath) == 0 {
		panic("Please specify -terragruntRepoPath")
	}

	dirs, err := findTerragrunt(terragruntRepoPath)
	if err != nil {
		fmt.Println(err)
	}
	createDependabotConfig(dirs, dependabotFile)
}

func findTerragrunt(dir string) ([]string, error) {
	var projects []string

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileExtension := filepath.Ext(path)
		if fileExtension != ".hcl" {
			return nil
		}
		if strings.Contains(path, ".terragrunt-cache") {
			return nil
		}

		if strings.Contains(path, "terragrunt.hcl") {
			projects = append(projects, filepath.Dir(path))
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer func() {
			_ = file.Close()
		}()

		return nil
	})
	return projects, err
}

func createDependabotConfig(dirs []string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("templates/config.tmpl")
	err = t.Execute(f, dirs)
	if err != nil {
		return err
	}
	return err
}
