package createItem_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Create Item", func() {
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("With Valid Input", func() {

		Context("And Item Does Not exists", func() {

			It("should save item to storage", func() {

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



	//Context("when Item Does Not Exist", func() {
	//
	//	var capturedItem domain.Item
	//
	//	//var mockItemStorage storage.ItemStorage
	//	//BeforeEach(func() {
	//	//	mockCtrl = gomock.NewController(GinkgoT())
	//	//	mockItemStorage := mock_storage.NewMockItemStorage(mockCtrl)
	//	//	mockItemStorage.
	//	//		EXPECT().
	//	//		Save(gomock.Any()).
	//	//		Do(func(arg domain.Item) {
	//	//			capturedItem = arg
	//	//		}).MaxTimes(1)
	//	//})
	//
	//	It("should create Item", func() {
	//		mockCtrl = gomock.NewController(GinkgoT())
	//		mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
	//		mockItemStorage.
	//			EXPECT().
	//			Save(gomock.Any()).
	//			Do(func(arg domain.Item) {
	//				capturedItem = arg
	//			}).MaxTimes(1)
	//
	//		itemName := "bag"
	//		input := createItem.Input{Name: itemName}
	//		usecase := createItem.Usecase{ItemStorage: mockItemStorage}
	//		usecase.Execute(input)
	//
	//		Expect(capturedItem.Name).To(Equal(itemName))
	//	})
	//})
	//
	//Context("when Item Exists", func() {
	//	It("should return error with message item already exists", func() {
	//
	//	})
	//})
})
