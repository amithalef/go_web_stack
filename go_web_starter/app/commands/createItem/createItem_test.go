package createItem_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smartystreets/assertions/should"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/createItem"
	mockstorage "github.com/amithnair91/go_web_stack/go_web_starter/app/commands/mocks"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

var _ = Describe("Create Item", func() {

	Context("With Valid Input", func() {
		input := createItem.Input{Name: "bag"}
		Context("And Item Does Not exists", func() {
			var mockCtrl = gomock.NewController(GinkgoT())
			var capturedItem *domain.Item
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			mockItemStorage.EXPECT().Exists(gomock.Any()).Return(false, nil).MaxTimes(1)
			mockItemStorage.EXPECT().Save(gomock.Any()).
				Do(func(arg *domain.Item) {
					capturedItem = arg
				}).MaxTimes(1)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			It("should save item to storage", func() {
				defer mockCtrl.Finish()
				usecase.Execute(input)
				Expect(capturedItem.Name).To(Equal(input.Name))
			})
		})

		Context("And Item exists", func() {
			var mockCtrl = gomock.NewController(GinkgoT())
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			mockItemStorage.EXPECT().Exists(gomock.Any()).Return(true, nil).MaxTimes(1)
			mockItemStorage.EXPECT().Save(gomock.Any()).MaxTimes(0)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			It("should return error with message Item already exists", func() {
				defer mockCtrl.Finish()
				err := usecase.Execute(input)

				Expect(err, should.NotBeNil)
				Expect(err.Error()).To(Equal("Item Already exists"))
			})
		})

		Context("And exists in storage fails", func() {
			databaseError := "database call failed"
			var mockCtrl = gomock.NewController(GinkgoT())
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			mockItemStorage.EXPECT().Exists(gomock.Any()).Return(false, errors.New(databaseError)).MaxTimes(1)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			It("should return error with message storage is down", func() {
				defer mockCtrl.Finish()
				err := usecase.Execute(input)

				Expect(err, should.NotBeNil)
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(`storage is down : %s`, databaseError)))
			})
		})
	})

	Context("With InValid Input", func() {
		invalidInput := createItem.Input{Name: ""}

		It("should return error with message invalid input", func() {
			var mockCtrl = gomock.NewController(GinkgoT())
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			err := usecase.Execute(invalidInput)

			Expect(err.Error()).To(Equal("name cannot be empty"))
		})
	})
})
