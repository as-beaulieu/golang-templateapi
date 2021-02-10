package service

import (
	"TemplateApi/src/mocks"
	"TemplateApi/src/models"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSimpleMessage(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("Simple Message", func() {
	mock := mocks.NewMockMessenger(gomock.NewController(GinkgoT()))
	Context("successful creation of simple messages", func() {
		mock.
			EXPECT().
			CreateSimpleMessage(gomock.AssignableToTypeOf(models.SimpleMessage{})).
			DoAndReturn(func(message models.SimpleMessage) (*models.SimpleMessageResponse, error) {
				return &models.SimpleMessageResponse{
					ID:      message.ID,
					Message: message.Message,
					Error:   nil,
				}, nil
			})
		When("when CreateSimpleMessage is called with a valid message", func() {
			correctMessage := models.SimpleMessage{
				ID:      "12345-2231-12312-11",
				Message: "This is a valid message",
			}
			result, err := mock.CreateSimpleMessage(correctMessage)
			It("returns a success with no errors", func() {
				Expect(err).To(BeNil())
				Expect(result.Error).To(BeNil())
			})
			It("returns the message in the response", func() {
				Expect(result.ID).To(Equal(correctMessage.ID))
				Expect(result.Message).To(Equal(correctMessage.Message))
			})
		})
	})
})
