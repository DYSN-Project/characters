package app

import (
	"context"
	"dysn/character/internal/config"
	"dysn/character/internal/manager"
	"dysn/character/internal/repository"
	"dysn/character/internal/service/database"
	"dysn/character/internal/service/logger"
	"dysn/character/internal/service/validation"
	"dysn/character/internal/transport/grpc"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) {
	cfg := config.NewConfig()
	log := logger.NewLogger()

	db := database.Init(ctx,cfg,log)
	characterRepo := repository.NewCharacterRepo(db)
	vld := validation.NewValidation()
	mng := manager.NewCharacterManager(cfg,log,characterRepo)

	srv := grpc.NewGrpcServer(cfg.GetSelfPort(), log, vld,mng)
	go srv.StartServer()
	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-sgn:
	}
}