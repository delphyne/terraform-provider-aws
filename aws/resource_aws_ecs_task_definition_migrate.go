package aws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/hashicorp/terraform/terraform"
)

func migrateEcsTaskDefinitionStateV0toV1(is *terraform.InstanceState, conn *ecs.ECS) (*terraform.InstanceState, error) {
	arn := is.Attributes["arn"]

	// We need to pull definitions from the API b/c they're unrecoverable from the checksum
	td, err := conn.DescribeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(arn),
	})
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(td.TaskDefinition.ContainerDefinitions)
	if err != nil {
		return nil, err
	}

	is.Attributes["container_definitions"] = string(b)

	return is, nil
}
