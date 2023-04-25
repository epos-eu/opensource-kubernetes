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
    "github.com/spf13/cobra"
	"github.com/joho/godotenv"
    "os/exec"
    "os"
    "log"
    "time"
    "regexp"
    "strings"
	"github.com/a8m/envsubst"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

var deployCmd = & cobra.Command {
    Use: "deploy",
    Short: "Deploy an environment on Kubernetes",
    Long: `Deploy an enviroment with .env set up on Kubernetes`,
    Run: func(cmd * cobra.Command, args[] string) {

        env, _ := cmd.Flags().GetString("env")
        context, _ := cmd.Flags().GetString("context")
        namespace, _ := cmd.Flags().GetString("namespace")
        tag, _ := cmd.Flags().GetString("tag")

        if env == "" {
            env = generateTempFile(configurations)
        }

        fileContents, err := os.ReadFile(env)
        if err != nil {
            log.Fatalln(err)
        }

        fileLines := lineBreakRegExp.Split(string(fileContents), -1)

        for _, line := range fileLines {
            if strings.Contains(line, "="){
                res1 := strings.Split(line, "=")
                os.Setenv(res1[0], res1[1])
            }
        }

        os.Setenv("CONTEXT", context)
        os.Setenv("DEPLOY_TAG", tag)
        os.Setenv("NAMESPACE", namespace)
        os.Setenv("DEPLOY_PATH", "/"+namespace+"/")
        os.Setenv("BASE_CONTEXT", "/"+namespace)

        setupIPs()

        operator, err = envsubst.Bytes([]byte(operator))
        if err != nil {
            log.Fatal(err)
        }
        rabbitmqoperatorfile := generateTempFile(operator)

        rabbitmq, err = envsubst.Bytes([]byte(rabbitmq))
        if err != nil {
            log.Fatal(err)
        }
        rabbitmqfile := generateTempFile(rabbitmq)

        logging, err = envsubst.Bytes([]byte(logging))
        if err != nil {
            log.Fatal(err)
        }
        loggingfile := generateTempFile(logging)

        secrets, err = envsubst.Bytes([]byte(secrets))
        if err != nil {
            log.Fatal(err)
        }
        secretsfile := generateTempFile(secrets)

        backoffice, err = envsubst.Bytes([]byte(backoffice))
        if err != nil {
            log.Fatal(err)
        }
        backofficefile := generateTempFile(backoffice)

        dataMetadata, err = envsubst.Bytes([]byte(dataMetadata))
        if err != nil {
            log.Fatal(err)
        }
        datametadatafile := generateTempFile(dataMetadata)

        externalAccess, err = envsubst.Bytes([]byte(externalAccess))
        if err != nil {
            log.Fatal(err)
        }
        externalaccessfile := generateTempFile(externalAccess)

        ingestor, err = envsubst.Bytes([]byte(ingestor))
        if err != nil {
            log.Fatal(err)
        }
        ingestorfile := generateTempFile(ingestor)

        metadataDatabase, err = envsubst.Bytes([]byte(metadataDatabase))
        if err != nil {
            log.Fatal(err)
        }
        metadatadatabasefile := generateTempFile(metadataDatabase)

        redisDatabase, err = envsubst.Bytes([]byte(redisDatabase))
        if err != nil {
            log.Fatal(err)
        }
        redisdatabasefile := generateTempFile(redisDatabase)

        resources, err = envsubst.Bytes([]byte(resources))
        if err != nil {
            log.Fatal(err)
        }
        resourcesfile := generateTempFile(resources)

        gateway, err = envsubst.Bytes([]byte(gateway))
        if err != nil {
            log.Fatal(err)
        }
        gatewayfile := generateTempFile(gateway)

        converter, err = envsubst.Bytes([]byte(converter))
        if err != nil {
            log.Fatal(err)
        }
        converterfile := generateTempFile(converter)

        list_of_services:= [13]string{rabbitmqoperatorfile, rabbitmqfile, loggingfile, secretsfile, 
            backofficefile, datametadatafile, externalaccessfile, ingestorfile, metadatadatabasefile, 
            redisdatabasefile, resourcesfile, gatewayfile, converterfile}

        if err := godotenv.Load(env);
        err != nil {
            log.Fatal("Error loading env variables from "+env+"\n")
            log.Fatal(err)
        }

        cmd.Printf(">> Deploy environment\n   >> Context: %s \n   >> Namespace: %s \n   >> Env file: %s \n   >> LocalAddress ip %s\n", context, namespace, env, os.Getenv("LOCAL_IP"))

        execute_command(cmd,exec.Command("kubectl",
        "config",
        "use-context",
        context))

        execute_command(cmd,exec.Command("kubectl",
        "apply",
        "-f",
        "https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.1/deploy/static/provider/cloud/deploy.yaml"))

        execute_command(cmd,exec.Command("kubectl",
        "create",
        "ns",
        namespace))

        time.Sleep(10 * time.Second)

        execute_command(cmd,exec.Command("kubectl",
            "apply",
            "-f",
            list_of_services[0]))

        time.Sleep(10 * time.Second)

        for i:= 1; i < 13; i++{

            execute_command(cmd,exec.Command(
            "kubectl",
            "apply",
            "-f",
            list_of_services[i],
            "-n",
            namespace))
            

            execute_command(cmd,exec.Command("kubectl",
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

func execute_command(cmd * cobra.Command, command * exec.Cmd) {
    cmd.Println(command.String())
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr
    if err := command.Run();
    err != nil {
        print(err)
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