package requirements_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	testconfig "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TargetedSpaceRequirement", func() {
	var (
		config coreconfig.ReadWriter
	)

	BeforeEach(func() {
		config = testconfig.NewRepositoryWithDefaults()
	})

	Context("when the user has targeted a space", func() {
		It("succeeds", func() {
			req := NewTargetedSpaceRequirement(config)
			err := req.Execute()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when the user does not have a space targeted", func() {
		It("errors", func() {
			config.SetSpaceFields(models.SpaceFields{})

			err := NewTargetedSpaceRequirement(config).Execute()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No space targeted"))
		})
	})
})
