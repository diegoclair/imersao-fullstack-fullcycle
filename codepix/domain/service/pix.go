package service

import (
	"github.com/diegoclair/imersao/codepix/contract"
	"github.com/diegoclair/imersao/codepix/domain/entity"
)

type pixService struct {
	svc *Service
}

//newPixService return a new instance of the service
func newPixService(svc *Service) contract.PixService {
	return &pixService{
		svc: svc,
	}
}

func (s *pixService) RegisterKey(key, kind, accountID string) (*entity.Pix, error) {

	account, err := s.svc.db.Postgres().Account().FindAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	pix, err := entity.NewPix(kind, account, key)
	if err != nil {
		return nil, err
	}

	return s.svc.db.Postgres().Pix().AddPixKey(pix)
}

func (s *pixService) FindKeyByID(key, kind string) (*entity.Pix, error) {
	return s.svc.db.Postgres().Pix().FindPixByKey(key, kind)
}
