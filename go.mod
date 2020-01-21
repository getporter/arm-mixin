module get.porter.sh/mixin/arm

go 1.13

require (
	get.porter.sh/porter v0.22.1-beta.1
	github.com/Azure/azure-sdk-for-go v38.2.0+incompatible
	github.com/Azure/go-autorest v12.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.0
	github.com/Azure/go-autorest/autorest/adal v0.5.0
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/gobuffalo/packr/v2 v2.7.1
	github.com/huandu/xstrings v1.3.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
