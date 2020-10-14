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

func TestService_CreateUser(t *testing.T) {
	Describe("testing CreateUser", func() {
		controller := gomock.NewController(t)
		serviceMock := mocks.NewMockService(controller)
		Context("successful creation of user", func() {
			serviceMock.
				EXPECT().
				CreateUser(gomock.Eq(models.User{})).
				DoAndReturn(func(user models.User) (*models.User, error) {
					time.Sleep(1 * time.Second)
					return &user, nil
				}).
				AnyTimes()
			When("when CreateUser is called with a valid request", func() {
				correctRequest := models.User{
					ID:   "1234",
					Name: "name",
				}
				result, err := serviceMock.CreateUser(correctRequest)
				It("returns a success with no errors", func() {
					Expect(err).To(nil)
				})
				It("returns the message in the response", func() {
					Expect(result.ID).To(Equal(correctRequest.ID))
					Expect(result.Name).To(Equal(correctRequest.Name))
				})
			})
		})
	})
}
