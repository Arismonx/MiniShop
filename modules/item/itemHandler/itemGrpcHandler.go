package itemHandler

import _itemUsecase "github.com/Arismonx/MiniShop/modules/item/itemUsecase"

// Create interface and struct
type (
	itemGrpcHandler struct {
		itemUsecase _itemUsecase.ItemUsecaseService
	}
)

// Create constructor
func NewItemGrpcHandler(
	itemUsecase _itemUsecase.ItemUsecaseService,
) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase}
}
