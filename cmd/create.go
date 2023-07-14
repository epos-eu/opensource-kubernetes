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
	"github.com/a8m/envsubst"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an environment on Kubernetes",
	Long:  `Deploy an enviroment with .env set up on Kubernetes`,
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		context, _ := cmd.Flags().GetString("context")
		namespace, _ := cmd.Flags().GetString("namespace")
		tag, _ := cmd.Flags().GetString("tag")
		isDefaultEnv := false
		if env == "" {
			env = generateTempFile(configurations)
			isDefaultEnv = true
		}
		fileContents, err := os.ReadFile(env)
		if err != nil {
			printError("Loading env variables from " + env + " cause: " + err.Error())
			os.Exit(0)
		}
		fileLines := lineBreakRegExp.Split(string(fileContents), -1)
		for _, line := range fileLines {
			if strings.Contains(line, "=") {
				res1 := strings.Split(line, "=")
				os.Setenv(res1[0], res1[1])
			}
		}

		os.Setenv("CONTEXT", context)
		os.Setenv("DEPLOY_TAG", tag)
		os.Setenv("NAMESPACE", namespace)
		os.Setenv("DEPLOY_PATH", "/"+namespace+"/")
		os.Setenv("BASE_CONTEXT", "/"+namespace)
		os.Setenv("POSTGRESQL_CONNECTION_STRING", "jdbc:postgresql://"+os.Getenv("POSTGRESQL_HOST")+"/"+os.Getenv("POSTGRES_DB")+"?user="+os.Getenv("POSTGRES_USER")+"&password="+os.Getenv("POSTGRESQL_PASSWORD")+"")
		if isDefaultEnv {
			checkImagesUpdate()
		}
		setupIPs()

		operator, err = envsubst.Bytes([]byte(operator))
		if err != nil {
			printError("Updating env variables of " + string(operator) + " cause: " + err.Error())
			os.Exit(0)
		}
		rabbitmqoperatorfile := generateTempFile(operator)

		rabbitmq, err = envsubst.Bytes([]byte(rabbitmq))
		if err != nil {
			printError("Updating env variables of " + string(rabbitmq) + " cause: " + err.Error())
			os.Exit(0)
		}
		rabbitmqfile := generateTempFile(rabbitmq)

		logging, err = envsubst.Bytes([]byte(logging))
		if err != nil {
			printError("Updating env variables of " + string(logging) + " cause: " + err.Error())
			os.Exit(0)
		}
		loggingfile := generateTempFile(logging)

		secrets, err = envsubst.Bytes([]byte(secrets))
		if err != nil {
			printError("Updating env variables of " + string(secrets) + " cause: " + err.Error())
			os.Exit(0)
		}
		secretsfile := generateTempFile(secrets)

		backoffice, err = envsubst.Bytes([]byte(backoffice))
		if err != nil {
			printError("Updating env variables of " + string(backoffice) + " cause: " + err.Error())
			os.Exit(0)
		}
		backofficefile := generateTempFile(backoffice)

		dataMetadata, err = envsubst.Bytes([]byte(dataMetadata))
		if err != nil {
			printError("Updating env variables of " + string(dataMetadata) + " cause: " + err.Error())
			os.Exit(0)
		}
		datametadatafile := generateTempFile(dataMetadata)

		externalAccess, err = envsubst.Bytes([]byte(externalAccess))
		if err != nil {
			printError("Updating env variables of " + string(externalAccess) + " cause: " + err.Error())
			os.Exit(0)
		}
		externalaccessfile := generateTempFile(externalAccess)

		ingestor, err = envsubst.Bytes([]byte(ingestor))
		if err != nil {
			printError("Updating env variables of " + string(ingestor) + " cause: " + err.Error())
			os.Exit(0)
		}
		ingestorfile := generateTempFile(ingestor)

		metadataDatabase, err = envsubst.Bytes([]byte(metadataDatabase))
		if err != nil {
			printError("Updating env variables of " + string(metadataDatabase) + " cause: " + err.Error())
			os.Exit(0)
		}
		metadatadatabasefile := generateTempFile(metadataDatabase)

		redisDatabase, err = envsubst.Bytes([]byte(redisDatabase))
		if err != nil {
			printError("Updating env variables of " + string(redisDatabase) + " cause: " + err.Error())
			os.Exit(0)
		}
		redisdatabasefile := generateTempFile(redisDatabase)

		resources, err = envsubst.Bytes([]byte(resources))
		if err != nil {
			printError("Updating env variables of " + string(resources) + " cause: " + err.Error())
			os.Exit(0)
		}
		resourcesfile := generateTempFile(resources)

		gateway, err = envsubst.Bytes([]byte(gateway))
		if err != nil {
			printError("Updating env variables of " + string(gateway) + " cause: " + err.Error())
			os.Exit(0)
		}
		gatewayfile := generateTempFile(gateway)

		converter, err = envsubst.Bytes([]byte(converter))
		if err != nil {
			printError("Updating env variables of " + string(converter) + " cause: " + err.Error())
			os.Exit(0)
		}
		converterfile := generateTempFile(converter)

		list_of_services := [13]string{rabbitmqoperatorfile, rabbitmqfile, loggingfile, secretsfile, metadatadatabasefile,
			backofficefile, datametadatafile, externalaccessfile, ingestorfile,
			redisdatabasefile, resourcesfile, gatewayfile, converterfile}

		if err := godotenv.Load(env); err != nil {
			printError("Error loading env variables from " + env + " cause: " + err.Error())
			os.Exit(0)
		}

		printSetup(env, context, namespace, tag)
		printTask("Switching context to " + context)

		execute_command(cmd, exec.Command("kubectl",
			"config",
			"use-context",
			context))

		printTask("Deploy of ingress nginx")

		execute_command(cmd, exec.Command("kubectl",
			"apply",
			"-f",
			"https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.1/deploy/static/provider/cloud/deploy.yaml"))

		printTask("Creating namespace: " + namespace)
		execute_command(cmd, exec.Command("kubectl",
			"create",
			"ns",
			namespace))

		time.Sleep(10 * time.Second)

		printTask("Deploy of " + list_of_services[0])

		execute_command(cmd, exec.Command("kubectl",
			"apply",
			"-f",
			list_of_services[0]))

		time.Sleep(10 * time.Second)

		for i := 1; i < 13; i++ {
			printTask("Deploy of " + list_of_services[i])
			execute_command(cmd, exec.Command(
				"kubectl",
				"apply",
				"-f",
				list_of_services[i],
				"-n",
				namespace))
			printTask("Waiting for conditions met")
			execute_command(cmd, exec.Command("kubectl",
				"wait",
				"--for=condition=Ready",
				"pods",
				"--all",
				"-n",
				namespace))

			time.Sleep(10 * time.Second)
		}

		print_urls()

	},
}

func execute_command(cmd *cobra.Command, command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		printError("Error on executing command, cause: " + err.Error())
	}
}

func init() {
	deployCmd.Flags().String("context", "", "Kubernetes context")
	deployCmd.MarkFlagRequired("context")
	deployCmd.Flags().String("namespace", "", "Kubernetes namespace")
	deployCmd.MarkFlagRequired("namespace")
	deployCmd.Flags().String("tag", "", "Version Tag")
	deployCmd.MarkFlagRequired("tag")
}
