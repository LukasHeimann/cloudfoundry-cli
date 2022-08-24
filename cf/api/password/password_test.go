package password_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/apifakes"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/net"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal/terminalfakes"
	testconfig "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	testnet "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/net"

	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/password"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/trace/tracefakes"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CloudControllerPasswordRepository", func() {
	It("updates your password", func() {
		req := apifakes.NewCloudControllerTestRequest(testnet.TestRequest{
			Method:   "PUT",
			Path:     "/Users/my-user-guid/password",
			Matcher:  testnet.RequestBodyMatcher(`{"password":"new-password","oldPassword":"old-password"}`),
			Response: testnet.TestResponse{Status: http.StatusOK},
		})

		passwordUpdateServer, handler, repo := createPasswordRepo(req)
		defer passwordUpdateServer.Close()

		apiErr := repo.UpdatePassword("old-password", "new-password")
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(apiErr).NotTo(HaveOccurred())
	})

	When("the inputs contains special characters", func() {
		It("handles escaping", func() {
			req := apifakes.NewCloudControllerTestRequest(testnet.TestRequest{
				Method:   "PUT",
				Path:     "/Users/my-user-guid/password",
				Matcher:  testnet.RequestBodyMatcher(`{"password":"more-\\-\\b\\\\-crazy","oldPassword":"crazy-\\.\\b-password"}`),
				Response: testnet.TestResponse{Status: http.StatusOK},
			})

			passwordUpdateServer, handler, repo := createPasswordRepo(req)
			defer passwordUpdateServer.Close()

			apiErr := repo.UpdatePassword(`crazy-\.\b-password`, `more-\-\b\\-crazy`)
			Expect(handler).To(HaveAllRequestsCalled())
			Expect(apiErr).NotTo(HaveOccurred())
		})
	})
})

func createPasswordRepo(req testnet.TestRequest) (passwordServer *httptest.Server, handler *testnet.TestHandler, repo Repository) {
	passwordServer, handler = testnet.NewServer([]testnet.TestRequest{req})

	configRepo := testconfig.NewRepositoryWithDefaults()
	configRepo.SetUaaEndpoint(passwordServer.URL)
	gateway := net.NewCloudControllerGateway(configRepo, time.Now, new(terminalfakes.FakeUI), new(tracefakes.FakePrinter), "")
	repo = NewCloudControllerRepository(configRepo, gateway)
	return
}
