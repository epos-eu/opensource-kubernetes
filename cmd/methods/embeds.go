// file: ./cmd/methods/embeds.go
package methods

import (
	_ "embed"
)

var (

	//go:embed "kubernetes-resources/data-portal-service.yaml"
	dataPortal []byte

	//go:embed "kubernetes-resources/backoffice-service.yaml"
	backoffice []byte

	//go:embed "kubernetes-resources/converter-service.yaml"
	converter []byte

	//go:embed "kubernetes-resources/external-access-service.yaml"
	externalAccess []byte

	//go:embed "kubernetes-resources/gateway-service.yaml"
	gateway []byte

	//go:embed "kubernetes-resources/ingestor-service.yaml"
	ingestor []byte

	//go:embed "kubernetes-resources/logging.yaml"
	logging []byte

	//go:embed "kubernetes-resources/metadata-database.yaml"
	metadataDatabase []byte

	//go:embed "kubernetes-resources/rabbitmq-operator.yaml"
	operator []byte

	//go:embed "kubernetes-resources/rabbitmq.yaml"
	rabbitmq []byte

	//go:embed "kubernetes-resources/resources-service.yaml"
	resources []byte

	//go:embed "kubernetes-resources/secrets.yaml"
	secrets []byte

	//go:embed "configurations/env.env"
	configurations []byte
)

func GetConfigurationsEmbed() []byte {
	return configurations
}

func GetDataPortalResourceEmbed() []byte {
	return dataPortal
}

func GetBackofficeResourceEmbed() []byte {
	return backoffice
}
func GetConverterResourceEmbed() []byte {
	return converter
}
func GetExternalAccessResourceEmbed() []byte {
	return externalAccess
}
func GetGatewayResourceEmbed() []byte {
	return gateway
}
func GetIngestorResourceEmbed() []byte {
	return ingestor
}
func GetLoggingResourceEmbed() []byte {
	return logging
}
func GetMetadataDatabaseResourceEmbed() []byte {
	return metadataDatabase
}
func GetOperatorResourceEmbed() []byte {
	return operator
}
func GetRabbitMQResourceEmbed() []byte {
	return rabbitmq
}
func GetResourcesResourceEmbed() []byte {
	return resources
}
func GetSecretsResourceEmbed() []byte {
	return secrets
}
