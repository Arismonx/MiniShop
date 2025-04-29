package inventoryHandler

import (
	_inventoryUsecase "github.com/Arismonx/MiniShop/modules/inventory/inventoryUsecase"
)

// *essentail* Create interface and struct
type (
	inventoryGrpcHandler struct {
		inventoryUsecase _inventoryUsecase.InventoryUsecaseService
	}
)

// *essntail* Create constructor of inventoryGrpcHandler
func NewInventoryGrpcHandler(
	inventoryUsecase _inventoryUsecase.InventoryUsecaseService,
) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{inventoryUsecase}
}
