package service

import (
	"TemplateApi/src/mocks"
	"TemplateApi/src/models"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("User", func() {
	defer GinkgoRecover()
	Describe("testing CreateUser", func() {
		userMock := mocks.NewMockUserOperator(gomock.NewController(GinkgoT()))
		Context("successful creation of user", func() {
			userMock.
				EXPECT().
				CreateUser(gomock.AssignableToTypeOf(models.User{})).
				DoAndReturn(func(user models.User) (*models.User, error) {
					return &user, nil
				})
			When("when CreateUser is called with a valid request", func() {
				correctRequest := models.User{
					ID:   "1234",
					Name: "name",
				}
				result, err := userMock.CreateUser(correctRequest)
				It("returns a success with no errors", func() {
					Expect(err).To(BeNil())
				})
				It("returns the message in the response", func() {
					Expect(result.ID).To(Equal(correctRequest.ID))
					Expect(result.Name).To(Equal(correctRequest.Name))
				})
			})
		})
	})
})
