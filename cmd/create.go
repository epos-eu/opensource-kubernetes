/*
    EPOS Open Source - Local installation with Docker
    Copyright (C) 2022  EPOS ERIC

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
    "bytes"
    "fmt"
    "io/ioutil"
	"github.com/a8m/envsubst"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

var deployCmd = & cobra.Command {
    Use: "deploy",
    Short: "Deploy an environment on docker",
    Long: `Deploy an enviroment with .env set up on docker`,
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

        os.Setenv("DEPLOY_TAG", tag)
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

        processingAccess, err = envsubst.Bytes([]byte(processingAccess))
        if err != nil {
            log.Fatal(err)
        }
        processingaccessfile := generateTempFile(processingAccess)

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

        eposGUI, err = envsubst.Bytes([]byte(eposGUI))
        if err != nil {
            log.Fatal(err)
        }
        eposguifile := generateTempFile(eposGUI)

        converter, err = envsubst.Bytes([]byte(converter))
        if err != nil {
            log.Fatal(err)
        }
        converterfile := generateTempFile(converter)

        list_of_services:= [14]string{rabbitmqoperatorfile, rabbitmqfile, loggingfile, secretsfile, 
            backofficefile, datametadatafile, externalaccessfile, ingestorfile, metadatadatabasefile, 
            processingaccessfile, resourcesfile, gatewayfile, eposguifile, converterfile}

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

        time.Sleep(1 * time.Second)

        execute_command(cmd,exec.Command("kubectl",
            "apply",
            "-f",
            list_of_services[0]))

        time.Sleep(1 * time.Second)

        for i:= 1; i < 14; i++{

            execute_command(cmd,exec.Command(
            "kubectl",
            "apply",
            "-f",
            list_of_services[i]))
            

            execute_command(cmd,exec.Command("kubectl",
            "wait",
            "--for=condition=Ready",
            "pods",
            "--all",
            "-n",
            namespace))
        }

        print_urls()

    },
}

func find_and_replace(file string) {
    input, err := ioutil.ReadFile(file)
    if err != nil {
            fmt.Println(err)
            os.Exit(1)
    }

    output := bytes.Replace(input, []byte("replaceme"), []byte("ok"), -1)

    if err = ioutil.WriteFile(file, output, 0666); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func execute_command(cmd * cobra.Command, command * exec.Cmd) {
    cmd.Printf(command.String())
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