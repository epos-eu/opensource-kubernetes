/*
   EPOS Open Source - Local installation with Kubernetes
   Copyright (C) 2023  EPOS ERIC

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
//file: ./cmd/populate.go
package cmd

import (
	_ "embed"

	"github.com/epos-eu/opensource-kubernetes/cmd/methods"
	"github.com/spf13/cobra"
)

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populate the existing environment with metadata information",
	Long:  `Populate the existing environment with metadata information in a specific folder`,
	RunE: func(cmd *cobra.Command, args []string) error {

		context, _ := cmd.Flags().GetString("context")
		path, _ := cmd.Flags().GetString("folder")
		env, _ := cmd.Flags().GetString("env")
		namespace, _ := cmd.Flags().GetString("namespace")
		tag, _ := cmd.Flags().GetString("tag")

		if err := methods.PopulateEnvironment(context, env, path, namespace, tag); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	populateCmd.Flags().String("context", "", "Kubernetes context")
	populateCmd.MarkFlagRequired("context")
	populateCmd.Flags().String("folder", "", "Folder where ttl files are located")
	populateCmd.MarkFlagRequired("folder")
	populateCmd.Flags().String("env", "", "Environment variable file")
	populateCmd.Flags().String("namespace", "", "Kubernetes namespace")
	populateCmd.MarkFlagRequired("namespace")
	populateCmd.Flags().String("tag", "", "Version Tag")
	populateCmd.MarkFlagRequired("tag")
}
