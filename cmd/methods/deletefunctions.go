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
//file: ./cmd/methods/deletefunctions.go
package methods

import (
	_ "embed"
	"os/exec"
)

func DeleteEnvironment(context string, namespace string) error {
	SetupIPs()
	PrintSetupShort(context, namespace)

	if err := ExecuteCommand(exec.Command("kubectl",
		"config",
		"use-context",
		context)); err != nil {
		return err
	}

	if err := ExecuteCommand(exec.Command("kubectl",
		"delete",
		"ns",
		namespace)); err != nil {
		return err
	}

	return nil
}
