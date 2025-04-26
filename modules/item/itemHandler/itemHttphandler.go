package itemHandler

import _itemUsecase "github.com/Arismonx/MiniShop/modules/item/itemUsecase"

// Create interface and struct
type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		itemUsecase _itemUsecase.ItemUsecaseService
	}
)

// Create constructor
func NewItemHttpHandler(
	itemUsecase _itemUsecase.ItemUsecaseService,
) ItemHttpHandlerService {
	return &itemHttpHandler{itemUsecase}
}
