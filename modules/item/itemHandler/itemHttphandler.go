package itemHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_itemUsecase "github.com/Arismonx/MiniShop/modules/item/itemUsecase"
)

// Create interface and struct
type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase _itemUsecase.ItemUsecaseService
	}
)

// Create constructor
func NewItemHttpHandler(
	cfg *config.Config,
	itemUsecase _itemUsecase.ItemUsecaseService,
) ItemHttpHandlerService {
	return &itemHttpHandler{cfg, itemUsecase}
}
