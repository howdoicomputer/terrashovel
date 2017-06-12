// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/GeertJohan/go.rice"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Generate a module directory.",
	Long:    `Generate a module directory.`,
	Aliases: []string{"create"},
	PreRunE: ArgLenCheck,
	Run:     NewModule,
}

func ArgLenCheck(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("A path argument is required.")
	}

	return nil
}

func NewModule(cmd *cobra.Command, args []string) {
	newModulePath := strings.Join(args, "")
	err := newModuleDirs(newModulePath)

	if err != nil {
		jww.ERROR.Println(err)
		return
	}

	err = newModuleTemplates(newModulePath)

	if err != nil {
		jww.ERROR.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf(`
Congratulations! Your new module has been instantiated!

If you want to try it out:

cd %s
bundle install
bundle exec kitchen test
`, newModulePath))
}

func newModuleDirs(newModulePath string) error {
	err := os.Mkdir(newModulePath, os.ModePerm)

	if err != nil {
		return err
	}

	dirs := []string{
		filepath.Join(newModulePath, "test"),
		filepath.Join(newModulePath, "test", "fixtures", "default"),
		filepath.Join(newModulePath, "test", "integration", "default"),
		filepath.Join(newModulePath, "test", "integration", "default", "controls"),
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func newModuleTemplates(newModulePath string) error {
	specDir := "test/integration/default/controls"
	fixturesDir := "test/fixtures/default"

	templates := map[string]string{
		"main.tf":                       filepath.Join(newModulePath, "main.tf"),
		"outputs.tf":                    filepath.Join(newModulePath, "outputs.tf"),
		"variables.tf":                  filepath.Join(newModulePath, "variables.tf"),
		"gitignore":                     filepath.Join(newModulePath, ".gitignore"),
		"README.md":                     filepath.Join(newModulePath, "README.md"),
		"Gemfile":                       filepath.Join(newModulePath, "Gemfile"),
		"kitchen.yml":                   filepath.Join(newModulePath, ".kitchen.yml"),
		fixturesDir + "/module.tf":      filepath.Join(newModulePath, fixturesDir, "module.tf"),
		fixturesDir + "/outputs.tf":     filepath.Join(newModulePath, fixturesDir, "outputs.tf"),
		specDir + "/ec2_server_spec.rb": filepath.Join(newModulePath, specDir, "ec2_server_spec.rb"),
		specDir + "/state_file_spec.rb": filepath.Join(newModulePath, specDir, "state_file_spec.rb"),
	}

	for source, destination := range templates {
		_, fileName := filepath.Split(source)

		templateBox, err := rice.FindBox("../templates")
		if err != nil {
			return err
		}

		templateString, err := templateBox.String(source)
		if err != nil {
			return err
		}

		template := template.Must(template.New(fileName).Parse(templateString))
		file, err := os.Create(destination)
		if err != nil {
			return err
		}

		err = template.Execute(file, template)
		if err != nil {
			return err
		}

		file.Close()
	}

	return nil
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
