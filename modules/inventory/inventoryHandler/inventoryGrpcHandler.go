package inventoryHandler

import (
	"github.com/Arismonx/MiniShop/config"
	_inventoryUsecase "github.com/Arismonx/MiniShop/modules/inventory/inventoryUsecase"
)

// *essentail* Create interface and struct
type (
	InventoryGrpcHandler interface{}

	inventoryGrpcHandler struct {
		cfg              *config.Config
		inventoryUsecase _inventoryUsecase.InventoryUsecaseService
	}
)

//*essntail* Create constructor of inventoryGrpcHandler

func NewInventoryGrpcHandler(
	cfg *config.Config,
	inventoryUsecase _inventoryUsecase.InventoryUsecaseService,
) InventoryGrpcHandler {
	return &inventoryGrpcHandler{cfg, inventoryUsecase}
}
