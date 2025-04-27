package inventoryHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_inventoryUsecase "github.com/Arismonx/MiniShop/modules/inventory/inventoryUsecase"
)

// *essential* Create interface and struct
type (
	InventoryQueueHandlerService interface{}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase _inventoryUsecase.InventoryUsecaseService
	}
)

// *essentail* Create contructor of inventoryHandler
func NewInventoryQueueHandler(
	cfg *config.Config,
	inventoryUsecase _inventoryUsecase.InventoryUsecaseService,
) InventoryQueueHandlerService {
	return &inventoryQueueHandler{cfg, inventoryUsecase}
}
