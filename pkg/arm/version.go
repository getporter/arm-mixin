package arm

import (
	"github.com/deislabs/porter-arm/pkg"
	"github.com/deislabs/porter/pkg/mixin"
	"github.com/deislabs/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "arm",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "DeisLabs",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
