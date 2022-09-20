package templates

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"get.porter.sh/mixin/arm/pkg/arm/auth"
	"get.porter.sh/porter/pkg/portercontext"
	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"
)

func newTestDeployer(ctx *portercontext.TestContext) Deployer {
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
	ctx := portercontext.NewTestContext(t)
	b, err := ioutil.ReadFile("testdata/test-arm.json")
	require.NoError(t, err)
	ctx.AddTestFile("testdata/test-arm.json", "/cnab/app/arm/aci.json")
	d := newTestDeployer(ctx)
	tpl, err := d.FindTemplate("arm/aci.json")
	assert.NoError(t, err)
	assert.Equal(t, string(b), string(tpl))

}
