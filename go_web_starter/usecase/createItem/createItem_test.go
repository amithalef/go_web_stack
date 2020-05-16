package createItem_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/usecase/createItem"
	"github.com/amithnair91/go_web_stack/go_web_starter/usecase/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Item", func() {
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when Item Does Not Exist", func() {

		It("should create Item", func() {
			itemName := "bag"

			mockItemStorage := mock_storage.NewMockItemStorage(mockCtrl)

			var capturedItem domain.Item
			mockItemStorage.
				EXPECT().
				Save(gomock.Any()).
				Do(func(arg domain.Item) {
					capturedItem = arg
				}).MaxTimes(1)

			input := createItem.Input{Name: itemName}
			usecase := createItem.Usecase{ItemStorage: mockItemStorage}
			usecase.Execute(input)

			Expect(capturedItem.Name).To(Equal(itemName))
		})
	})

	Context("when Item Exists", func() {
		It("should return error with message item already exists", func() {

		})
	})
})

//Given("Create Item"){
//
//	When("Item Already exists"){
//
//		then("throws item already exists error"){}
//}
//When("Item does not exist"){
//
//	then("item is created successfully"){
//
//}
//}
//}
