package itemHandler

import _itemUsecase "github.com/Arismonx/MiniShop/modules/item/itemUsecase"

// Create interface and struct
type (
	ItemGrpcHandlerService interface{}

	itemGrpcHandler struct {
		itemUsecase _itemUsecase.ItemUsecaseService
	}
)

// Create constructor
func NewItemGrpcHandler(
	itemUsecase _itemUsecase.ItemUsecaseService,
) ItemGrpcHandlerService {
	return &itemGrpcHandler{itemUsecase}
}
