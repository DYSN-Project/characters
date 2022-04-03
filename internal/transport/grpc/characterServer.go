package grpc

import (
	"context"
	"dysn/character/internal/manager"
	"dysn/character/internal/model"
	frmRequest "dysn/character/internal/model/request"
	"dysn/character/internal/service/logger"
	"dysn/character/internal/service/validation"
	"dysn/character/internal/transport/grpc/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CharacterServer struct {
	logger     *logger.Logger
	validation validation.ValidationInterface
	manager    manager.CharacterManagerInterface
	pb.UnimplementedCharacterServer
}

func NewCharacterServer(logger *logger.Logger,
	validation validation.ValidationInterface,
	manager manager.CharacterManagerInterface) *CharacterServer {
	return &CharacterServer{
		logger:     logger,
		validation: validation,
		manager:    manager,
	}
}

func (srv *CharacterServer) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	srv.logger.InfoLog.Println("pong")

	return &emptypb.Empty{}, nil
}

func (srv *CharacterServer) List(_ context.Context, request *pb.ListRequest) (*pb.CharacterList, error) {
	frm := frmRequest.NewListForm(int(request.GetLimit()), int(request.GetOffset()))
	if err := srv.validation.ValidateList(frm); err != nil {
		srv.logger.ErrorLog.Println(err)
		return nil, err
	}
	list, err := srv.manager.GetCharacters(frm.Limit, frm.Offset)
	if err != nil {
		return nil, err
	}

	rsp := pb.CharacterList{}
	rsp.List = make([]*pb.CharacterFull, len(list))
	for i, v := range list {
		rsp.List[i] = &pb.CharacterFull{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Image:       v.Image,
		}
	}

	return &rsp, nil
}

func (srv *CharacterServer) CreateCharacter(_ context.Context, request *pb.CreateCharacterRequest) (*pb.CharacterFull, error) {
	frm := frmRequest.NewCreateForm(request.GetName(), request.GetDescription())
	if err := srv.validation.ValidateCreate(frm); err != nil {
		srv.logger.ErrorLog.Println(err)
		return nil, err
	}
	character, err := srv.manager.CreateCharacter(frm.Name, frm.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CharacterFull{
		Id:          character.ID,
		Name:        character.Name,
		Description: character.Description,
		Image:       "",
	}, nil
}

func (srv *CharacterServer) UpdateCharacter(_ context.Context, request *pb.UpdateCharacterRequest) (*emptypb.Empty, error) {
	frm := frmRequest.NewUpdateForm(request.GetName(), request.GetDescription(), int(request.GetStatus()))
	if err := srv.validation.ValidateUpdate(frm); err != nil {
		srv.logger.ErrorLog.Println(err)
		return nil, err
	}

	updateFilters := model.NewCharacter(frm.Name,frm.Description,"",frm.Status)
	 err := srv.manager.UpdateCharacter(request.GetId(),updateFilters)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (srv *CharacterServer) DeleteCharacter(_ context.Context, request *pb.DeleteCharacterRequest) (*emptypb.Empty, error) {
	frm := frmRequest.NewDeleteForm(request.GetId())
	if err := srv.validation.ValidateDelete(frm); err != nil {
		srv.logger.ErrorLog.Println(err)
		return nil, err
	}

	if err := srv.manager.DeleteCharacter(request.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}