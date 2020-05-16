package createItem_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/usecase/createItem"
	mockstorage "github.com/amithnair91/go_web_stack/go_web_starter/usecase/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Create Item", t, func() {
		var mockCtrl *gomock.Controller
		//defer mockCtrl.Finish()

		Convey("With Valid Inputs", func() {
			input := createItem.Input{Name: "bag"}

			Convey("And Item Does Not Exists", func() {
				var capturedItem domain.Item
				mockCtrl = gomock.NewController(GinkgoT())
				mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
				mockItemStorage.
					EXPECT().
					Save(gomock.Any()).
					Do(func(arg domain.Item) {
						capturedItem = arg
					}).MaxTimes(1)
				usecase := createItem.Usecase{ItemStorage: mockItemStorage}

				Convey("should save to Storage", func() {
					usecase.Execute(input)
					So(capturedItem.Name, ShouldEqual, input.Name)
				})
			})

			Convey("And Item Already Exists", func() {

				Convey("should return error", func() {

				})

				Convey("should not save to storage", func() {})
			})
		})

		Convey("With InValid Inputs", func() {

		})

	})
}
