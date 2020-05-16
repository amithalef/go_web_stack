package createItem_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/usecase/createItem"
	mockstorage "github.com/amithnair91/go_web_stack/go_web_starter/usecase/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Item", func() {

	Context("With Valid Input", func() {
		input := createItem.Input{Name: "bag"}

		Context("And Item Does Not exists", func() {
			var mockCtrl = gomock.NewController(GinkgoT())

			AfterEach(func() {
				mockCtrl.Finish()
			})
			var capturedItem domain.Item
			mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
			mockItemStorage.
				EXPECT().
				Save(gomock.Any()).
				Do(func(arg domain.Item) {
					capturedItem = arg
				}).MaxTimes(1)
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}

			It("should save item to storage", func() {
				usecase.Execute(input)
				Expect(capturedItem.Name).To(Equal(input.Name))
			})
		})

		Context("And Item exists", func() {
			It("should return error with message Item already exists", func() {})

			It("should not save item to storage", func() {})
		})

	})

	Context("With InValid Input", func() {

		It("should return error with message invalid input", func() {})

	})
})
