package server

import (
	_inventoryHandler "github.com/Arismonx/MiniShop/modules/inventory/inventoryHandler"
	_inventoryRepository "github.com/Arismonx/MiniShop/modules/inventory/inventoryRepository"
	_inventoryUsecase "github.com/Arismonx/MiniShop/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	//initialize inventory
	repo := _inventoryRepository.NewInventoryRepository(s.db)
	usecase := _inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := _inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	queueHandler := _inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)
	gRPCHandler := _inventoryHandler.NewInventoryGrpcHandler(usecase)

	_ = httpHandler
	_ = gRPCHandler
	_ = queueHandler

	//Group path
	inventory := s.app.Group("/inventory_v1")

	//health check
	inventory.GET("", s.healthCheck)
}
