package createItem_test

//
//func TestSpec(t *testing.T) {
//
//	Convey("Create Item", t, func() {
//		var mockCtrl *gomock.Controller
//		//defer mockCtrl.Finish()
//
//		Convey("With Valid Inputs", func() {
//			input := createItem.Input{Name: "bag"}
//
//			Convey("And Item Does Not Exists", func() {
//				var capturedItem domain.Item
//				mockCtrl = gomock.NewController(GinkgoT())
//				mockItemStorage := mockstorage.NewMockItemStorage(mockCtrl)
//				mockItemStorage.
//					EXPECT().
//					Save(gomock.Any()).
//					Do(func(arg domain.Item) {
//						capturedItem = arg
//					}).MaxTimes(1)
//				usecase := createItem.Usecase{ItemStorage: mockItemStorage}
//
//				Convey("should save to Storage", func() {
//					usecase.Execute(input)
//					So(capturedItem.Name, ShouldEqual, input.Name)
//				})
//			})
//
//			Convey("And Item Already Exists", func() {
//
//				Convey("should return error", func() {
//
//				})
//
//				Convey("should not save to storage", func() {})
//			})
//		})
//
//		Convey("With InValid Inputs", func() {
//
//		})
//
//	})
//}
