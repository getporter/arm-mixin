package templates

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"

	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/deislabs/porter-arm/pkg/arm/auth"
	"github.com/deislabs/porter/pkg/context"
	"github.com/stretchr/testify/assert"
)

func newTestDeplyer(ctx *context.TestContext) Deployer {
	env, _ := azure.EnvironmentFromName("")
	authorizer, _ := auth.GetBearerTokenAuthorizer(env, "", "", "")
	resourceDeploymentsClient := resourcesSDK.NewDeploymentsClientWithBaseURI(
		"",
		"",
	)
	resourceDeploymentsClient.Authorizer = authorizer

	resourceGroupsClient := resourcesSDK.NewGroupsClientWithBaseURI(
		"",
		"",
	)
	resourceGroupsClient.Authorizer = authorizer
	return NewDeployer(ctx.Context, resourceGroupsClient, resourceDeploymentsClient)

}
func TestLoadTemplate(t *testing.T) {
	ctx := context.NewTestContext(t)
	b, err := ioutil.ReadFile("testdata/test-arm.json")
	require.NoError(t, err)
	ctx.AddTestFile("testdata/test-arm.json", "/cnab/app/arm/aci.json")
	d := newTestDeplyer(ctx)
	tpl, err := d.FindTemplate("arm/aci.json")
	assert.NoError(t, err)
	assert.Equal(t, string(b), string(tpl))

}
