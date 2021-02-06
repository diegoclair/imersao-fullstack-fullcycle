package usecase

import (
	"github.com/diegoclair/imersao/codepix/domain/model"
)

// PixUseCase is the repository implementation
type PixUseCase struct {
	PixRepository     model.PixRepositoryInterface
	AccountRepository model.AccountRepositoryInterface
}

func (a *PixUseCase) RegisterKey(key, kind, accountID string) (*model.Pix, error) {

	account, err := a.AccountRepository.FindAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	pix, err := model.NewPix(kind, account, key)
	if err != nil {
		return nil, err
	}

	return a.PixRepository.AddPixKey(pix)
}

func (a *PixUseCase) FindKeyByID(key, kind string) (*model.Pix, error) {
	return a.PixRepository.FindKeyByID(key, kind)
}
