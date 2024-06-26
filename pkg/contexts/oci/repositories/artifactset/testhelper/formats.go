package testhelper

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"

	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/artifactset"
)

func TestForAllFormats(msg string, f func(fmt string)) {
	DescribeTable(fmt.Sprintf("%s: structure format handling", msg), f,
		Entry("OCM format", artifactset.FORMAT_OCM),
		Entry("OCI format", artifactset.FORMAT_OCI),
	)
}

func FTestForAllFormats(msg string, f func(fmt string)) {
	FDescribeTable(fmt.Sprintf("%s: structure format handling", msg), f,
		Entry("OCM format", artifactset.FORMAT_OCM),
		Entry("OCI format", artifactset.FORMAT_OCI),
	)
}
