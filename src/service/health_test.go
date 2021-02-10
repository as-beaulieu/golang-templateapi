package service

import (
	"TemplateApi/src/mocks"
	"TemplateApi/src/models"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("Heartbeat", func() {
	defer GinkgoRecover()
	positiveResponse := models.HeartbeatResponse{
		Message: "API Running",
		Time:    time.Now(),
	}
	mock := mocks.NewMockHealthReporter(gomock.NewController(GinkgoT()))
	Context("successful operation of Heartbeat()", func() {
		mock.EXPECT().Heartbeat().Return(&positiveResponse, nil)
		When("when Heartbeat is called", func() {
			result, err := mock.Heartbeat()
			It("completes with no errors", func() {
				Expect(err).To(BeNil())
			})
			It("and returns the positive response", func() {
				Expect(result).Should(Equal(&positiveResponse))
			})
		})
	})
})
