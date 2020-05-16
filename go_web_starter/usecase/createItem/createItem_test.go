package createItem_test

import (
	"../createItem"
	"../storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Item", func() {
	Context("when Item Does Not Exist", func() {
		It("should create Item", func() {
			var itemStorage = storage.ItemStorage{}
			input := createItem.Input{}
			usecase := createItem.Usecase{}
			usecase.Execute(input)

			itemStorage.save.called.once()
		})
	})

	Context("when Item Exists", func() {
		It("should return error with message item already exists", func() {
			Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
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
