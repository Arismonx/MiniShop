package inventoryUsecase

import _inventoryRepository "github.com/Arismonx/MiniShop/modules/inventory/inventoryRepository"

// *essential* Create interface and struct
type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository _inventoryRepository.InventoryRepositoryService
	}
)

// *essential* Create constructor of inventoryUsecase
func NewInventoryRepository(
	inventoryRepository _inventoryRepository.InventoryRepositoryService,
) InventoryUsecaseService {
	return &inventoryUsecase{inventoryRepository}
}
