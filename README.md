
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
  delete      Delete an environment on kubernetes
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
      --context string     Kubernetes context
  -h, --help               help for deploy
      --namespace string   Kubernetes namespace
      --tag string         Version Tag
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
      --env string         Environment variable file
      --folder string      Folder where ttl files are located
  -h, --help               help for populate
      --namespace string   Kubernetes namespace
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
| API_HOST | ${API_HOST} | API Host IP, if not set is generated automatically using machine IP |
| EXECUTE_HOST | ${API_HOST} | Internal variable to setup redirections for the external access service, if not set is generated automatically using machine IP |
| DEPLOY_PATH | / | Context path of the environment|
| BASE_CONTEXT | empty value | Context path name of the environment (similar to DEPLOY_PATH but without the initial /) |
| API_PATH | /api/v1 | API GATEWAY access path|
| GUI_PORT | 8000 | Port used by EPOS Data Portal or other GUIs |
| BACKOFFICE_GUI_PORT | 9000 | Port used by EPOS Backoffice UI or other Backoffice GUIs |
| API_PORT | 8080 | Port used by EPOS API Gateway |
| IS_MONITORING_AUTH | false | Variable used to protect monitoring endpoint via JWT |
| IS_AAI_ENABLED | false | Variable used to protect backoffice endpoints via AAI service |

### RabbitMQ configuration

| Name | Standard Value | Description |
|--|--|--|
| BROKER_USERNAME | changeme | RabbitMQ username |
| BROKER_PASSWORD | changeme | RabbitMQ password |
| BROKER_VHOST | changeme | RabbitMQ vhost |

### RabbitMQ configuration

| Name | Standard Value | Description |
|--|--|--|
| POSTGRES_USER | postgres | Database user |
| POSTGRESQL_PASSWORD | changeme | Database password |
| POSTGRES_DB | cerif | Database name |
| POSTGRESQL_CONNECTION_STRING | jdbc:postgresql://postgrescerif:5432/${POSTGRES_DB}?user=${POSTGRES_USER}&password=${POSTGRESQL_PASSWORD} | Database connection string based on previous configurations |
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
| MONITORING_URL | changeme | Monitoring service url |
| MONITORING_USER | changeme | Monitoring service username |
| MONITORING_PWD | changeme | Monitoring service password |

### Monitoring Service configuration

| Name | Standard Value | Description |
|--|--|--|
| Kubernetes_REGISTRY | epos | Kubernetes registry url |
| REGISTRY_USERNAME | changeme | Kubernetes registry username |
| REGISTRY_PASSWORD | changeme | Kubernetes registry password |

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
| FACETS_DEFAULT | false | |
| FACETS_TYPE_DEFAULT | categories | |
| REDIS_SERVER | redis-server | |
| INGESTOR_HASH | 3F58A1895982CC81A2E5CEDA7DD9AC7009DF9998 | |


### Kubernetes Images for Open Source 

| Variable name | Image name | Default Tag |
|--|--|--|
|--|--|--|
| MESSAGE_BUS_IMAGE | rabbitmq | 3.11.7-management |
| REDIS_IMAGE | redis | 7.0.11 |
| GATEWAY_IMAGE | epos-api-gateway | 1.1.0 |
| RESOURCES_SERVICE_IMAGE | resources-service | 1.3.2 |
| INGESTOR_IMAGE | ingestor-service | 1.3.1 |
| EXTERNAL_ACCESS_IMAGE | external-access-service | 1.3.2 |
| BACKOFFICE_SERVICE_IMAGE | backoffice-service | 2.1.0 |
| CONVERTER_IMAGE | converter-service | 1.1.5 |
| DATA_METADATA_SERVICE_IMAGE | data-metadata-service | 2.3.17 |
| METADATA_DB_IMAGE | metadata-database-deploy | 2.2.0 |

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
