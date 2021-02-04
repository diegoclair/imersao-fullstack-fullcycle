package usecase

import "github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"

// pixUseCase is the repository implementation
type pixUseCase struct {
	pixRepository     model.PixRepositoryInterface
	accountRepository model.AccountRepositoryInterface
}

func (a *pixUseCase) RegisterKey(key, kind, accountID string) (*model.Pix, error) {
	account, err := a.accountRepository.FindAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	pix, err := model.NewPix(kind, account, key)
	if err != nil {
		return nil, err
	}

	return a.pixRepository.AddPixKey(pix)
}

func (a *pixUseCase) FindKeyByID(key, kind string) (*model.Pix, error) {
	return a.pixRepository.FindKeyByID(key, kind)
}
