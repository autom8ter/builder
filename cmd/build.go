// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
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
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type BuildCmd struct {
	Project string
	AppName string
	Version string
}

func NewBuildCmd() (*BuildCmd, error) {
	var version, proj, app string
	var err error
	proj, err = util.EnvPrompt("PROJECT_ID")
	if err != nil {
		return nil, err
	}
	app, err = util.EnvPrompt("APP_NAME")
	if err != nil {
		return nil, err
	}
	version, err = util.EnvPrompt("VERSION")
	if err != nil {
		return nil, err
	}
	return &BuildCmd{
		Project: proj,
		AppName: app,
		Version: version,
	}, nil
}

// buildCmd rep resents the build command
var buildCmd = &cobra.Command{
	Use: "build",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := NewBuildCmd()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(string(util.Shell(fmt.Sprintf(`docker build -t gcr.io/%s/%s:%s .`, c.Project, c.AppName, c.Version))))
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

}
