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
//file: ./cmd/export.go
package cmd

import (
	_ "embed"

	"github.com/epos-eu/opensource-kubernetes/cmd/methods"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export configuration files in output folder, options: [env]",
	Long:  `Export configuration files for customization in output folder, options: [env]`,
	RunE: func(cmd *cobra.Command, args []string) error {

		file, _ := cmd.Flags().GetString("file")
		output, _ := cmd.Flags().GetString("output")

		if err := methods.ExportVariablesEnvironment(file, output); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	exportCmd.Flags().String("file", "", "File to export, available options: [env]")
	exportCmd.MarkFlagRequired("file")
	exportCmd.Flags().String("output", "", "Output folder")
	exportCmd.MarkFlagRequired("output")
}
