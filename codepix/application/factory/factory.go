package factory

import (
	"log"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/imersao/codepix/domain/contract"
	"github.com/diegoclair/imersao/codepix/domain/service"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
	"github.com/diegoclair/imersao/codepix/infrastructure/data"
)

//Services is the factory to all serrvices
type Services struct {
	Cfg                *config.EnvironmentVariables
	Mapper             mapper.Mapper
	PixService         contract.PixService
	TransactionService contract.TransactionService
}

var (
	instance *Services
	once     sync.Once
)

//GetDomainServices to get instace of all services
func GetDomainServices() *Services {
	once.Do(func() {
		data, err := data.Connect()
		if err != nil {
			log.Fatalf("Error to connect data repositories: %v", err)
		}

		cfg := config.GetConfigEnvironment()
		mapper := mapper.New()
		svm := service.NewServiceManager()
		svc := service.New(data)

		instance = &Services{
			Cfg:                cfg,
			Mapper:             mapper,
			PixService:         svm.PixService(svc),
			TransactionService: svm.TransactionService(svc),
		}
	})
	return instance
}
