
# EPOS Open Source - Kubernetes installer

## Introduction

EPOS Open Source - Kubernetes installer is part of the EPOS Open Source project for local installation using Kubernetes.
It contains a set of Kubernetes images to deploy the EPOS ecosystem. 

Use `opensource-kubernetes` binary to spin up local environment on Linux, Mac OS X or Windows.

## Prerequisites

Kubernetes Tools installed on your local machine and access to a Kubernetes Cluster.
For further information follow the official guidelines: https://kubernetes.io/docs/home/

## Installation

Download the binary file according to your OS.

Give permissions on `opensource-kubernetes` file and move on binary folder from a Terminal (in Linux/MacOS):

```
chmod +x opensource-kubernetes
sudo mv opensource-kubernetes /usr/local/bin/opensource-kubernetes
```

## Usage

```
./opensource-kubernetes <command>
```

The `<command>` field value is one of the following listed below:

```
EPOS Open Source CLI installer to deploy the EPOS System using Kubernetes

Usage:
  opensource-kubernetes [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete an environment on Kubernetes
  deploy      Deploy an environment on Kubernetes
  export      Export configuration files in output folder, options: [env]
  help        Help about any command
  populate    Populate the existing environment with metadata information

Flags:
  -h, --help      help for opensource-kubernetes
  -v, --version   version for opensource-kubernetes

Use "opensource-kubernetes [command] --help" for more information about a command.
```

## Deploy a new environment

```
Deploy an enviroment with .env set up on Kubernetes

Usage:
  opensource-kubernetes deploy [flags]

Flags:
      --autoupdate string   Auto update the images versions (true|false)
      --context string      Kubernetes context
  -h, --help                help for deploy
      --namespace string    Kubernetes namespace
      --tag string          Version Tag
      --update string       Update of an existing deployment (true|false), default false
```

## Delete the existing environment

```
Delete an enviroment on Kubernetes using Namespace

Usage:
  opensource-kubernetes delete [flags]

Flags:
      --context string     Kubernetes context
  -h, --help               help for delete
      --namespace string   Kubernetes namespace
```

## Populate the existing environment with metadata

### Automatic option: 

Download or create TTL files according to EPOS-DCAT-AP and use the following command:

```
Populate the existing environment with metadata information in a specific folder

Usage:
  opensource-kubernetes populate [flags]

Flags:
      --context string     Kubernetes context
      --env string         Environment variable file
      --folder string      Folder where ttl files are located
  -h, --help               help for populate
      --namespace string   Kubernetes namespace
      --tag string         Version Tag
```

### Manual option

Use the API Gateway endpoint to manually ingest metadata TTL files into the catalogue.

## Export configuration file and Kubernetes-compose

```
Export configuration files for customization in output folder, options: [env]

Usage:
  opensource-kubernetes export [flags]

Flags:
      --file string     File to export, available options: [env]
  -h, --help            help for export
      --output string   Output folder
```


## Access URLs

EPOS Data Portal: 
```
http://<your-ip>/<DEPLOY_PATH>
```

EPOS Backoffice: 
```
http://<your-ip>/<DEPLOY_PATH>
```

EPOS API Gateway: 
```
http://<your-ip>/<DEPLOY_PATH>/<API_PATH>
```

## Environment Variables
### Base environment configuration

| Name | Standard Value | Description |
|--|--|--|
| PROTOCOL | http | |
| API_PATH | /api/v1 | API GATEWAY access path |
| INGRESS_CLASS | nginx | Ingress class used in the Kubernetes cluster |


### RabbitMQ configuration

| Name | Standard Value | Description |
|--|--|--|
| BROKER_HOST | rabbitmq | RabbitMQ host |
| BROKER_USERNAME | changeme | RabbitMQ username |
| BROKER_PASSWORD | changeme | RabbitMQ password |
| BROKER_VHOST | changeme | RabbitMQ vhost |

### PostgreSQL configuration

| Name | Standard Value | Description |
|--|--|--|
| POSTGRESQL_HOST | metadatacatalogue:5432 | Database service |
| POSTGRES_USER | postgres | Database user |
| POSTGRESQL_PASSWORD | changeme | Database password |
| POSTGRES_DB | cerif | Database name |
| PERSISTENCE_NAME | EPOSDataModel | Persistence Name of scientific metadata |
| PERSISTENCE_NAME_PROCESSING | EPOSProcessing | Persistence Name of processing metadata |

### Data Metadata Service configuration

| Name | Standard Value | Description |
|--|--|--|
| NUM_OF_PUBLISHERS | 10 | Number of publishers on rabbitmq |
| NUM_OF_CONSUMERS | 10 | Number of consumers on rabbitmq |
| CONNECTION_POOL_INIT_SIZE | 1 | Initial number of connections to database |
| CONNECTION_POOL_MIN_SIZE | 1 | Minimum number of connections to database |
| CONNECTION_POOL_MAX_SIZE | 20 | Maximum number of connections to database |

### Monitoring Service configuration

| Name | Standard Value | Description |
|--|--|--|
| MONITORING | false | True if activate interaction between system and monitoring service |
| MONITORING_URL | empty | Monitoring service url |
| MONITORING_USER | empty | Monitoring service username |
| MONITORING_PWD | empty | Monitoring service password |

### Monitoring Service configuration

| Name | Standard Value | Description |
|--|--|--|
| DOCKER_REGISTRY | epos | Docker registry url |
| REGISTRY_USERNAME | changeme | Docker registry username |
| REGISTRY_PASSWORD | changeme | Docker registry password |

### Other Environment variables

| Name | Standard Value | Description |
|--|--|--|
| LOAD_RESOURCES_API | true | |
| LOAD_INGESTOR_API | true | |
| LOAD_EXTERNAL_ACCESS_API | true | |
| LOAD_BACKOFFICE_API | true | |
| LOAD_PROCESSING_API | false | |
| IS_MONITORING_AUTH | false | |
| IS_AAI_ENABLED | false | |
| SECURITY_KEY | empty | |
| AAI_SERVICE_ENDPOINT | empty | |
| FACETS_DEFAULT | true | |
| FACETS_TYPE_DEFAULT | categories | |
| INGESTOR_HASH | FA9BEB99E4029AD5A6615399E7BBAE21356086B3 | "changeme" Security key|

## Maintenance

We regularly update images used in this stack.

## Contributing

If you want to contribute to a project and make it better, your help is very welcome. Contributing is also a great way to learn more about social coding on Github, new technologies and and their ecosystems and how to make constructive, helpful bug reports, feature requests and the noblest of all contributions: a good, clean pull request.

### How to make a clean pull request

Look for a project's contribution instructions. If there are any, follow them.

- Create a personal fork of the project on Github/GitLab.
- Clone the fork on your local machine. Your remote repo on Github/GitLab is called `origin`.
- Add the original repository as a remote called `upstream`.
- If you created your fork a while ago be sure to pull upstream changes into your local repository.
- Create a new branch to work on! Branch from `develop` if it exists, else from `master` or  `main`.
- Implement/fix your feature, comment your code.
- Follow the code style of the project, including indentation.
- If the project has tests run them!
- Write or adapt tests as needed.
- Add or change the documentation as needed.
- Squash your commits into a single commit with git's [interactive rebase](https://help.github.com/articles/interactive-rebase). Create a new branch if necessary.
- Push your branch to your fork on Github/GitLab, the remote `origin`.
- From your fork open a pull request in the correct branch. Target the project's `develop` branch if there is one, else go for `master` or  `main`!
- …
- If the maintainer requests further changes just push them to your branch. The PR will be updated automatically.
- Once the pull request is approved and merged you can pull the changes from `upstream` to your local repo and delete
your extra branch(es).

And last but not least: Always write your commit messages in the present tense. Your commit message should describe what the commit, when applied, does to the code – not what you did to the code.
