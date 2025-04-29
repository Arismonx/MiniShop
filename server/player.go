package server

import (
	_playerHandler "github.com/Arismonx/MiniShop/modules/player/playerHandler"
	_playerRepository "github.com/Arismonx/MiniShop/modules/player/playerRepository"
	_playerUsecase "github.com/Arismonx/MiniShop/modules/player/playerUsecase"
)

func (s *server) playerService() {
	//initialize player
	repo := _playerRepository.NewPlayerRepository(s.db)
	usecase := _playerUsecase.NewPlayerUsecase(repo)
	httpHandler := _playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	gRPCHandler := _playerHandler.NewPlayerGrpcHandler(usecase)

	_ = httpHandler
	_ = gRPCHandler

	//Group path
	player := s.app.Group("/player_v1")

	//health check
	player.GET("", s.healthCheck)
}
