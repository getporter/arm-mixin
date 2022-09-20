//go:generate packr2
package arm

import (
	"bufio"
	"fmt"
	"io/ioutil"

	resourcesSDK "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"

	"get.porter.sh/mixin/arm/pkg/arm/auth"
	arm "get.porter.sh/mixin/arm/pkg/arm/templates"
	"get.porter.sh/porter/pkg/runtime"

	"github.com/pkg/errors"
)

type Mixin struct {
	runtime.RuntimeConfig
	cfg Config
	//also add the azure clients here
}

// New arm mixin client, initialized with useful defaults.
func New() *Mixin {
	return &Mixin{
		RuntimeConfig: runtime.NewConfig(),
	}

}

func (m *Mixin) LoadConfigFromEnvironment() error {
	cfg, err := GetConfigFromEnvironment()
	if err != nil {
		return err
	}
	m.cfg = cfg
	return nil
}

func (m *Mixin) getPayloadData() ([]byte, error) {
	reader := bufio.NewReader(m.In)
	data, err := ioutil.ReadAll(reader)
	return data, errors.Wrap(err, "could not read the payload from STDIN")
}

func (m *Mixin) getARMDeployer() (arm.Deployer, error) {

	azureConfig := m.cfg
	azureSubscriptionID := azureConfig.SubscriptionID

	authorizer, err := auth.GetBearerTokenAuthorizer(
		azureConfig.Environment,
		azureConfig.TenantID,
		azureConfig.ClientID,
		azureConfig.ClientSecret,
	)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("couldn't build ARM deployer %s", err))
	}

	resourceDeploymentsClient := resourcesSDK.NewDeploymentsClientWithBaseURI(
		azureConfig.Environment.ResourceManagerEndpoint,
		azureSubscriptionID,
	)
	resourceDeploymentsClient.Authorizer = authorizer

	resourceGroupsClient := resourcesSDK.NewGroupsClientWithBaseURI(
		azureConfig.Environment.ResourceManagerEndpoint,
		azureSubscriptionID,
	)
	resourceGroupsClient.Authorizer = authorizer

	armDeployer := arm.NewDeployer(
		m.Context,
		resourceGroupsClient,
		resourceDeploymentsClient,
	)

	return armDeployer, nil
}
