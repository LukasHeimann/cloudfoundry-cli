package copyapplicationsource_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/apifakes"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/copyapplicationsource"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/net"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal/terminalfakes"
	testconfig "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	testnet "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/net"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/trace/tracefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CopyApplicationSource", func() {
	var (
		repo       Repository
		testServer *httptest.Server
		configRepo coreconfig.ReadWriter
	)

	setupTestServer := func(reqs ...testnet.TestRequest) {
		testServer, _ = testnet.NewServer(reqs)
		configRepo.SetAPIEndpoint(testServer.URL)
	}

	BeforeEach(func() {
		configRepo = testconfig.NewRepositoryWithDefaults()
		gateway := net.NewCloudControllerGateway(configRepo, time.Now, new(terminalfakes.FakeUI), new(tracefakes.FakePrinter), "")
		repo = NewCloudControllerCopyApplicationSourceRepository(configRepo, gateway)
	})

	AfterEach(func() {
		testServer.Close()
	})

	Describe(".CopyApplication", func() {
		BeforeEach(func() {
			setupTestServer(apifakes.NewCloudControllerTestRequest(testnet.TestRequest{
				Method: "POST",
				Path:   "/v2/apps/target-app-guid/copy_bits",
				Matcher: testnet.RequestBodyMatcher(`{
					"source_app_guid": "source-app-guid"
				}`),
				Response: testnet.TestResponse{
					Status: http.StatusCreated,
				},
			}))
		})

		It("should return a CopyApplicationModel", func() {
			err := repo.CopyApplication("source-app-guid", "target-app-guid")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
