package manager

import (
	"dysn/character/internal/config"
	"dysn/character/internal/helper"
	"dysn/character/internal/model"
	"dysn/character/internal/repository"
	"dysn/character/internal/service/logger"
	"errors"
)

var (
	defaultCountCharacters = 10
	errRepoError           = errors.New("internal server error")
	errCharacterNotFound   = errors.New("character not found")
)

type CharacterManagerInterface interface {
	GetCharacters(limit, offset int) ([]*model.Character, error)
	CreateCharacter(name, description string) (*model.Character, error)
	UpdateCharacter(id string, model *model.Character) error
	DeleteCharacter(id string) error
}

type CharacterManager struct {
	cfg           *config.Config
	lgr           *logger.Logger
	characterRepo repository.CharacterRepoInterface
}

func NewCharacterManager(cfg *config.Config,
	lgr *logger.Logger,
	characterRepo repository.CharacterRepoInterface) *CharacterManager {
	return &CharacterManager{
		cfg:           cfg,
		lgr:           lgr,
		characterRepo: characterRepo,
	}
}

func (charMan *CharacterManager) GetCharacters(limit, offset int) ([]*model.Character, error) {
	if limit == 0 {
		limit = defaultCountCharacters
	}

	list, err := charMan.characterRepo.List(limit, offset)
	if err != nil {
		charMan.lgr.ErrorLog.Println(err)

		return nil, err
	}

	return list, nil
}

func (charMan *CharacterManager) CreateCharacter(name, description string) (*model.Character, error) {
	mdl := model.NewCharacter(name, description, "", helper.StatusActive) //TODO:add filename
	character, err := charMan.characterRepo.Create(mdl)
	if err != nil {
		charMan.lgr.ErrorLog.Println(err)

		return nil, errRepoError
	}

	return character, nil
}

func (charMan *CharacterManager) UpdateCharacter(id string, model *model.Character) error {
	exist, err := charMan.characterRepo.Exist(id)
	if err != nil {
		charMan.lgr.ErrorLog.Println(err)
		return errRepoError
	}
	if !exist {
		return errCharacterNotFound
	}

	if err := charMan.characterRepo.Update(id, model); err != nil {
		charMan.lgr.ErrorLog.Println(err)

		return errRepoError
	}

	return nil
}

func (charMan *CharacterManager) DeleteCharacter(id string) error {
	exist, err := charMan.characterRepo.Exist(id)
	if err != nil {
		charMan.lgr.ErrorLog.Println(err)

		return errRepoError
	}
	if !exist {
		return errCharacterNotFound
	}

	deleted := charMan.characterRepo.Delete(id)
	if deleted != nil {
		charMan.lgr.ErrorLog.Println(err)

		return errRepoError
	}

	return nil
}
