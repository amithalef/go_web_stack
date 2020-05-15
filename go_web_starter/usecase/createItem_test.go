package usecase_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Item", func() {
		Context("when Item Does Not Exist", func() {
			It("should create Item", func() {

				var itemStorage = ItemStorage{}

				createItem.execute(createItemInput)

				itemStorage.save.called.once()


				//Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
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