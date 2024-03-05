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
//file: ./cmd/create.go
package cmd

import (
	"github.com/epos-eu/opensource-kubernetes/cmd/methods"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an environment on Kubernetes",
	Long:  `Deploy an enviroment with .env set up on Kubernetes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		env, _ := cmd.Flags().GetString("env")
		context, _ := cmd.Flags().GetString("context")
		namespace, _ := cmd.Flags().GetString("namespace")
		tag, _ := cmd.Flags().GetString("tag")
		autoupdate, _ := cmd.Flags().GetString("autoupdate")
		update, _ := cmd.Flags().GetString("update")

		if err := methods.CreateEnvironment(env, context, namespace, tag, autoupdate, update); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	deployCmd.Flags().String("context", "", "Kubernetes context")
	deployCmd.MarkFlagRequired("context")
	deployCmd.Flags().String("namespace", "", "Kubernetes namespace")
	deployCmd.MarkFlagRequired("namespace")
	deployCmd.Flags().String("tag", "", "Version Tag")
	deployCmd.MarkFlagRequired("tag")
	deployCmd.Flags().String("autoupdate", "", "Auto update the images versions (true|false)")
	deployCmd.Flags().String("update", "", "Update of an existing deployment (true|false), default false")
}
