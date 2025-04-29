package server

import (
	_itemHandler "github.com/Arismonx/MiniShop/modules/item/itemHandler"
	_itemRepository "github.com/Arismonx/MiniShop/modules/item/itemRepository"
	_itemUsecase "github.com/Arismonx/MiniShop/modules/item/itemUsecase"
)

func (s *server) itemService() {
	//initialize item
	repo := _itemRepository.NewItemRepository(s.db)
	usecase := _itemUsecase.NewItemUsecase(repo)
	httpHandler := _itemHandler.NewItemHttpHandler(s.cfg, usecase)
	gRPCHandler := _itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = gRPCHandler

	//Group path
	item := s.app.Group("/item_v1")

	//health check
	_ = item
}
