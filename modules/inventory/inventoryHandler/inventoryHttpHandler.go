package inventoryHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_inventoryUsecase "github.com/Arismonx/MiniShop/modules/inventory/inventoryUsecase"
)

// *essential* Create interface and struct
type (
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase _inventoryUsecase.InventoryUsecaseService
	}
)

// *essentail* Create contructor of inventoryHandler
func NewInventoryHttpHandler(
	cfg *config.Config,
	inventoryUsecase _inventoryUsecase.InventoryUsecaseService,
) InventoryHttpHandlerService {
	return &inventoryHttpHandler{cfg, inventoryUsecase}
}
