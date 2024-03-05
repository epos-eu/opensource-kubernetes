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
//file: ./cmd/methods/createfunctions.go
package methods

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/a8m/envsubst"
	"github.com/joho/godotenv"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func CreateEnvironment(env string, context string, namespace string, tag string, autoupdate string, update string) error {

	if update != "true" && update != "false" {
		update = "false"
	}

	envtagname := ""

	if namespace != "" {
		envtagname += namespace
	}
	if tag != "" {
		envtagname += tag
	}
	if envtagname != "" {
		envtagname += "-"
	}

	envtagname = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(envtagname, "-")
	os.Setenv("PREFIX", envtagname)
	dname := GenerateDirectoryName()

	if !strings.HasPrefix(tag, "\"") {
		tag = "\"" + tag + "\""
	}

	if err := RemoveContents(dname); err != nil {
		PrintError("Error on removing the content from directory " + err.Error())
		return err
	}
	if err := CreateDirectory(dname); err != nil {
		PrintError("Error on creating the directory " + err.Error())
		return err
	}

	isDefaultEnv := false
	if env == "" {
		ret_env, err := GenerateTempFile(dname, "configurations", GetConfigurationsEmbed())
		if err != nil {
			return err
		}
		env = ret_env
		isDefaultEnv = true
	}
	fileContents, err := os.ReadFile(env)
	if err != nil {
		PrintError("Loading env variables from " + env + " cause: " + err.Error())
		return err
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
	if autoupdate == "true" || isDefaultEnv {
		if err := CheckImagesUpdate(); err != nil {
			PrintError("Error on updating the docker container images " + err.Error())
			return err
		}
	}
	SetupIPs()

	copyoperator := GetOperatorResourceEmbed()
	operator, err = envsubst.Bytes([]byte(copyoperator))
	if err != nil {
		PrintError("Updating env variables of " + string(operator) + " cause: " + err.Error())

	}
	rabbitmqoperatorfile, err := GenerateTempFile(dname, "operator", operator)
	if err != nil {
		return err
	}

	copyrabbitmq := GetRabbitMQResourceEmbed()
	rabbitmq, err = envsubst.Bytes([]byte(copyrabbitmq))
	if err != nil {
		PrintError("Updating env variables of " + string(rabbitmq) + " cause: " + err.Error())
		return err
	}

	rabbitmqfile, err := GenerateTempFile(dname, "rabbitmq", rabbitmq)
	if err != nil {
		return err
	}

	copylogging := GetLoggingResourceEmbed()
	logging, err = envsubst.Bytes([]byte(copylogging))
	if err != nil {
		PrintError("Updating env variables of " + string(logging) + " cause: " + err.Error())
		return err
	}
	loggingfile, err := GenerateTempFile(dname, "logging", logging)
	if err != nil {
		return err
	}

	copysecrets := GetSecretsResourceEmbed()
	secrets, err = envsubst.Bytes([]byte(copysecrets))
	if err != nil {
		PrintError("Updating env variables of " + string(secrets) + " cause: " + err.Error())

	}
	secretsfile, err := GenerateTempFile(dname, "secrets", secrets)
	if err != nil {
		return err
	}

	copybackoffice := GetBackofficeResourceEmbed()
	backoffice, err = envsubst.Bytes([]byte(copybackoffice))
	if err != nil {
		PrintError("Updating env variables of " + string(backoffice) + " cause: " + err.Error())

	}
	backofficefile, err := GenerateTempFile(dname, "backoffice", backoffice)
	if err != nil {
		return err
	}

	copyexternalAccess := GetExternalAccessResourceEmbed()
	externalAccess, err = envsubst.Bytes([]byte(copyexternalAccess))
	if err != nil {
		PrintError("Updating env variables of " + string(externalAccess) + " cause: " + err.Error())

	}
	externalaccessfile, err := GenerateTempFile(dname, "externalAccess", externalAccess)
	if err != nil {
		return err
	}

	copyingestor := GetIngestorResourceEmbed()
	ingestor, err = envsubst.Bytes([]byte(copyingestor))
	if err != nil {
		PrintError("Updating env variables of " + string(ingestor) + " cause: " + err.Error())

	}
	ingestorfile, err := GenerateTempFile(dname, "ingestor", ingestor)
	if err != nil {
		return err
	}

	copymetadataDatabase := GetMetadataDatabaseResourceEmbed()
	metadataDatabase, err = envsubst.Bytes([]byte(copymetadataDatabase))
	if err != nil {
		PrintError("Updating env variables of " + string(metadataDatabase) + " cause: " + err.Error())

	}
	metadatadatabasefile, err := GenerateTempFile(dname, "metadataDatabase", metadataDatabase)
	if err != nil {
		return err
	}

	copyresources := GetResourcesResourceEmbed()
	resources, err = envsubst.Bytes([]byte(copyresources))
	if err != nil {
		PrintError("Updating env variables of " + string(resources) + " cause: " + err.Error())

	}
	resourcesfile, err := GenerateTempFile(dname, "resources", resources)
	if err != nil {
		return err
	}

	copygateway := GetGatewayResourceEmbed()
	gateway, err = envsubst.Bytes([]byte(copygateway))
	if err != nil {
		PrintError("Updating env variables of " + string(gateway) + " cause: " + err.Error())

	}
	gatewayfile, err := GenerateTempFile(dname, "gateway", gateway)
	if err != nil {
		return err
	}

	copydataPortal := GetDataPortalResourceEmbed()
	dataPortal, err = envsubst.Bytes([]byte(copydataPortal))
	if err != nil {
		PrintError("Updating env variables of " + string(dataPortal) + " cause: " + err.Error())

	}
	dataportalfile, err := GenerateTempFile(dname, "dataPortal", dataPortal)
	if err != nil {
		return err
	}

	copyconverter := GetConverterResourceEmbed()
	converter, err = envsubst.Bytes([]byte(copyconverter))
	if err != nil {
		PrintError("Updating env variables of " + string(converter) + " cause: " + err.Error())

	}
	converterfile, err := GenerateTempFile(dname, "converter", converter)
	if err != nil {
		return err
	}

	list_of_services := [14]string{rabbitmqoperatorfile, rabbitmqfile, loggingfile, secretsfile, metadatadatabasefile,
		backofficefile, externalaccessfile, ingestorfile, resourcesfile, gatewayfile, dataportalfile, converterfile}

	if err := godotenv.Overload(env); err != nil {
		PrintError("Error loading env variables from " + env + " cause: " + err.Error())

	}

	PrintSetup(env, context, namespace, tag)
	PrintTask("Switching context to " + context)

	if err := ExecuteCommand(exec.Command("kubectl",
		"config",
		"use-context",
		context)); err != nil {
		return err
	}

	if update == "true" {
		PrintTask("Check if namespace exists: " + namespace)
		if err := ExecuteCommand(exec.Command("kubectl",
			"get",
			"ns",
			namespace)); err != nil {
			PrintTask("Namespace doesn't exists, creating namespace: " + namespace)
			if err := ExecuteCommand(exec.Command("kubectl",
				"create",
				"ns",
				namespace)); err != nil {
				return err
			}
		}
	} else {
		PrintTask("Creating namespace: " + namespace)
		if err := ExecuteCommand(exec.Command("kubectl",
			"create",
			"ns",
			namespace)); err != nil {
			return err
		}
	}

	time.Sleep(10 * time.Second)

	PrintTask("Deploy of " + list_of_services[0])

	if err := ExecuteCommand(exec.Command("kubectl",
		"apply",
		"-f",
		list_of_services[0])); err != nil {
		return err
	}

	time.Sleep(10 * time.Second)

	for i := 1; i < 12; i++ {
		PrintTask("Deploy of " + list_of_services[i])
		if err := ExecuteCommand(exec.Command(
			"kubectl",
			"apply",
			"-f",
			list_of_services[i],
			"-n",
			namespace)); err != nil {
			PrintError("Error deploying the service " + list_of_services[i])
			return err
		}
	}
	PrintWait("Waiting for conditions met")
	time.Sleep(20 * time.Second)
	if err := ExecuteCommand(exec.Command("kubectl",
		"wait",
		"--for=condition=Ready",
		"pods",
		"--all",
		"-n",
		namespace)); err != nil {
		PrintError("Error on waiting for the conditions met")
		return err
	}
	PrintWait("Waiting for all services up and running")

	time.Sleep(40 * time.Second)
	PrintTask("Restarting gateway")

	if err := ExecuteCommand(exec.Command("kubectl",
		"rollout",
		"restart",
		"deployment",
		"-n",
		namespace,
		"gateway-deployment")); err != nil {
		PrintError("Error restarting the gateway service")
		return err
	}
	if err := ExecuteCommand(exec.Command("kubectl",
		"wait",
		"--for=condition=Ready",
		"pods",
		"--all",
		"-n",
		namespace)); err != nil {
		PrintError("Error on waiting for the conditions met")
		return err
	}
	time.Sleep(20 * time.Second)

	c1 := exec.Command("kubectl", "get", "ingress", "-n", namespace, "gateway-ingress", "-o", "jsonpath={.status.loadBalancer.ingress[*].hostname}")

	out, err := ExportHostname(c1)
	if err != nil {
		PrintError("Error on waiting for the conditions met")
		return err
	}
	os.Setenv("LOCAL_IP", out)
	PrintUrls()
	return nil
}
