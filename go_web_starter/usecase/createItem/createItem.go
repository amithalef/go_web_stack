package createItem

import "github.com/amithnair91/go_web_stack/go_web_starter/usecase/storage"

type Usecase struct {
	ItemStorage storage.ItemStorage
}

type Input struct {
	Name string
}

func (c *Usecase) Execute(input Input) {

}
