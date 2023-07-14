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
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an environment on Kubernetes",
	Long:  `Delete an enviroment on Kubernetes using Namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		context, _ := cmd.Flags().GetString("context")
		namespace, _ := cmd.Flags().GetString("namespace")
		setupIPs()
		printSetupShort(context, namespace)
		command := exec.Command("kubectl",
			"config",
			"use-context",
			context)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			printError("Error on executing command, cause:" + err.Error())
		}
		command = exec.Command("kubectl",
			"delete",
			"ns",
			namespace)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			printError("Error on executing command, cause:" + err.Error())
		}
	},
}

func init() {
	deleteCmd.Flags().String("context", "", "Kubernetes context")
	exportCmd.MarkFlagRequired("context")
	deleteCmd.Flags().String("namespace", "", "Kubernetes namespace")
	exportCmd.MarkFlagRequired("namespace")
}
