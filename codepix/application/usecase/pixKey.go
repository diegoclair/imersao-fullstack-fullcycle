package usecase

import "github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"

// pixUseCase is the repository implementation
type pixUseCase struct {
	pixKeyRepository  model.PixKeyRepositoryInterface
	accountRepository model.AccountRepositoryInterface
}

func (a *pixUseCase) RegisterKey(key, kind, accountID string) (*model.PixKey, error) {
	account, err := a.accountRepository.FindAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	return a.pixKeyRepository.AddPixKey(pixKey)
}

func (a *pixUseCase) FindKeyByID(key, kind string) (*model.PixKey, error) {
	return a.pixKeyRepository.FindKeyByID(key, kind)
}
