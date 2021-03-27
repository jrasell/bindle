package command

import (
	"strings"

	"github.com/mikenomitch/bindle/utils"
)

type Source struct{}

func (f *Source) Help() string {
	helpText := `
Some helper text goes here
`
	return strings.TrimSpace(helpText)
}

func (f *Source) Synopsis() string {
	return "Add a new source for a Bindle package"
}

func (f *Source) Name() string { return "source" }

func (f *Source) Run(args []string) int {
	catalogsFilePath := ".bindle/catalogs"
	catalogName := args[0]
	catalogDir := catalogsFilePath + "/" + catalogName

	gitRepoUrl := args[1]

	err := utils.CloneRepoToDir(gitRepoUrl, catalogDir)
	utils.Handle(err, "error cloning default catalog")

	utils.Log("Successfully added source \"" + catalogName + "\"")

	return 0
}
