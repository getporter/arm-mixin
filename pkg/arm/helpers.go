package arm

import (
	"testing"

	"get.porter.sh/porter/pkg/config"
	"get.porter.sh/porter/pkg/portercontext"
	"get.porter.sh/porter/pkg/runtime"
)

type TestMixin struct {
	*Mixin
	TestContext *portercontext.TestContext
}

// NewTestMixin initializes a mixin test client, with the output buffered, and an in-memory file system.
func NewTestMixin(t *testing.T) *TestMixin {
	testConfig := config.NewTestConfig(t)
	m := &TestMixin{
		Mixin: &Mixin{
			RuntimeConfig: runtime.NewConfigFor(testConfig.Config),
		},
		TestContext: testConfig.TestContext,
	}

	return m
}
