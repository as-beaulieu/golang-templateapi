package service

import (
	"TemplateApi/src/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestService_CreateSimpleMessage(t *testing.T) {
	Describe("testing CreateSimpleMessage", func() {
		service := ServiceBuilder{}.Build()
		Context("successful creation of simple messages", func() {
			When("when CreateSimpleMessage is called with a valid message", func() {
				correctMessage := models.SimpleMessage{
					ID:      "12345-2231-12312-11",
					Message: "This is a valid message",
				}
				result, err := service.CreateSimpleMessage(correctMessage)
				It("returns a success with no errors", func() {
					Expect(err).To(nil)
					Expect(result.Error).To(nil)
				})
				It("returns the message in the response", func() {
					Expect(result.ID).To(Equal(correctMessage.ID))
					Expect(result.Message).To(Equal(correctMessage.Message))
				})
			})
		})
	})
}
