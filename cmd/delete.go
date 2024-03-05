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
//file: ./cmd/delete.go
package cmd

import (
	_ "embed"

	"github.com/epos-eu/opensource-kubernetes/cmd/methods"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an environment on Kubernetes",
	Long:  `Delete an enviroment on Kubernetes using Namespace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		context, _ := cmd.Flags().GetString("context")
		namespace, _ := cmd.Flags().GetString("namespace")

		if err := methods.DeleteEnvironment(context, namespace); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	deleteCmd.Flags().String("context", "", "Kubernetes context")
	exportCmd.MarkFlagRequired("context")
	deleteCmd.Flags().String("namespace", "", "Kubernetes namespace")
	exportCmd.MarkFlagRequired("namespace")
}
