package arm

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMixin_UnmarshalInstallStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/install-input.yaml")
	require.NoError(t, err)

	var step InstallStep
	err = yaml.Unmarshal(b, &step)
	require.NoError(t, err)

	assert.Equal(t, "Create Azure MySQL", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, AzureOutput{"MYSQL_HOST", "MYSQL_HOST"}, step.Outputs[0])

	assert.Equal(t, "mysql", step.Type)
	assert.Equal(t, "mysql-azure-porter-demo", step.Name)
	assert.Equal(t, "porter-test", step.ResourceGroup)
	assert.Equal(t, map[string]interface{}{"location": "eastus", "serverName": "myserver"}, step.Parameters)
}

func TestMixin_UnmarshalInstallAction(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/install-input-with-action.yaml")
	require.NoError(t, err)

	var action InstallAction
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)

	require.Equal(t, 1, len(action.Steps))
	step := action.Steps[0]

	assert.Equal(t, "Create Azure MySQL", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, AzureOutput{"MYSQL_HOST", "MYSQL_HOST"}, step.Outputs[0])

	assert.Equal(t, "mysql", step.Type)
	assert.Equal(t, "mysql-azure-porter-demo", step.Name)
	assert.Equal(t, "porter-test", step.ResourceGroup)
	assert.Equal(t, map[string]interface{}{"location": "eastus", "serverName": "myserver"}, step.Parameters)
}
