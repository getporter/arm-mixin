package templates

import (
	"testing"

	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/stretchr/testify/assert"
)

func TestGetOutputs_ValidOutputs(t *testing.T) {
	deployment := &resourcesSDK.DeploymentExtended{
		Properties: &resourcesSDK.DeploymentPropertiesExtended{
			Outputs: map[string]interface{}{
				"output1": map[string]interface{}{
					"value": "amaterasu",
				},
				"output2": map[string]interface{}{
					"value": 108,
				},
			},
		},
	}

	outputs, err := getOutputs(deployment)
	assert.NoError(t, err)
	assert.Equal(t, "amaterasu", outputs["output1"])
	assert.Equal(t, 108, outputs["output2"])
}

func TestGetOutputs_InvalidOutputs(t *testing.T) {
	deployment := &resourcesSDK.DeploymentExtended{
		Properties: &resourcesSDK.DeploymentPropertiesExtended{
			Outputs: "invalid",
		},
	}

	outputs, err := getOutputs(deployment)
	assert.Error(t, err)
	assert.Nil(t, outputs)
}

func TestGetOutputs_NoOutputs(t *testing.T) {
	deployment := &resourcesSDK.DeploymentExtended{
		Properties: &resourcesSDK.DeploymentPropertiesExtended{
			Outputs: nil,
		},
	}

	outputs, err := getOutputs(deployment)
	assert.NoError(t, err)
	assert.Empty(t, outputs)
}
