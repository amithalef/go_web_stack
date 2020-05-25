package createItem_test

import (
	"errors"
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/createItem"
	mockstorage "github.com/amithnair91/go_web_stack/go_web_starter/app/commands/mocks"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smartystreets/assertions/should"
)

var _ = Describe("Create Item", func() {

	Context("With Valid Input", func() {
		input := createItem.NewInput("bag")
		Context("Then", func() {
			var mockCtrl = gomock.NewController(GinkgoT())
			var capturedItem *domain.Item
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			mockItemStorage.EXPECT().Save(gomock.Any()).
				Do(func(arg *domain.Item) {
					capturedItem = arg
				}).MaxTimes(1)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			It("should save item to storage", func() {
				defer mockCtrl.Finish()
				usecase.Execute(*input)
				Expect(capturedItem.Name).To(Equal(input.Name))
			})
		})

		Context("And save in storage fails", func() {
			databaseError := "database call failed"
			var mockCtrl = gomock.NewController(GinkgoT())
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			mockItemStorage.EXPECT().Save(gomock.Any()).Return(domain.Item{}, errors.New(databaseError)).MaxTimes(1)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			It("should return error with message storage is down", func() {
				defer mockCtrl.Finish()
				err := usecase.Execute(*input)

				Expect(err, should.NotBeNil)
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(`save to storage failed : %s`, databaseError)))
			})
		})
	})

	Context("With InValid Inputs", func() {
		It("should return error with message invalid input when name is empty", func() {
			invalidInput := createItem.NewInput("")
			var mockCtrl = gomock.NewController(GinkgoT())
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			err := usecase.Execute(*invalidInput)

			Expect(err.Error()).To(ContainSubstring("name cannot be empty"))
		})
	})
})
