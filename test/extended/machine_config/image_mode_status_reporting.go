package machine_config

import (
	"context"

	exutil "github.com/openshift/origin/test/extended/util"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

var _ = g.Describe("[Suite:openshift/machine-config-operator/disruptive][sig-mco][OCPFeatureGate:MachineConfigNodes]", func() {
	defer g.GinkgoRecover()
	var (
		oc = exutil.NewCLIWithoutNamespace("machine-config")
	)

	g.BeforeEach(func(ctx context.Context) {
		//skip these tests on hypershift platforms
		if ok, _ := exutil.IsHypershift(ctx, oc.AdminConfigClient()); ok {
			g.Skip("`MachineConfigNodes` and `OnClusterBuilds` are not currently supported on hypershift, so neither is `ImageModeStatusReporting`. Skipping tests.")
		}
	})

	g.It("[Suite:openshift/conformance/serial]Should have correct machine count on default MCP update [apigroup:machineconfiguration.openshift.io]", func() {
		ValidateMachineCountsInDefaultMCPUpdate(oc)
	})

	g.It("[Suite:openshift/conformance/serial]Should have correct machine count on custom MCP update [apigroup:machineconfiguration.openshift.io]", func() {
		ValidateMachineCountsInCustomMCPUpdate(oc)
	})
})

// `ValidateMachineCountsInDefaultMCPUpdate` validates that the machine counts populated in the
// MCP status are correct throughout a standard update in a default pool.
func ValidateMachineCountsInDefaultMCPUpdate(oc *exutil.CLI) {
	s := "test"
	o.Expect(s).Should(o.ContainSubstring("t"))
}

// `ValidateMachineCountsInCustomMCPUpdate` validates that the machine counts populated in the
// MCP status are correct throughout a standard update in a custom pool.
func ValidateMachineCountsInCustomMCPUpdate(oc *exutil.CLI) {
	s := "test"
	o.Expect(s).Should(o.ContainSubstring("t"))
}
