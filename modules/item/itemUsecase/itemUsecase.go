package itemUsecase

import _itemRepository "github.com/Arismonx/MiniShop/modules/item/itemRepository"

// Create interface and struct
type (
	ItemUsecaseService interface{}

	itemUsecase struct {
		itemRepository _itemRepository.ItemRepositoryService
	}
)

// Create constructor
func NewItemUsecase(
	itemRepository _itemRepository.ItemRepositoryService,
) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}
