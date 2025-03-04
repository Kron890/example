package usecase

import (
	"fmt"
	"template/internal/example_module"
)

type UseCase struct {
	dbRepo          example_module.ModuleDatabase
	externalApiRepo example_module.ModuleExternalService
}
// отправляем все в структуру для работы с ней 
func NewUseCase(dbRepo example_module.ModuleDatabase, externalApiRepo example_module.ModuleExternalService) *UseCase {
	return &UseCase{
		dbRepo:          dbRepo,
		externalApiRepo: externalApiRepo,
	}
}

func (u *UseCase) ProcessTicker(ticker string) (string, error) {
	exists, err := u.dbRepo.CheckTicker(ticker)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", example_module.ErrTickerNotFound
	}

	price, err := u.externalApiRepo.GetPrice(ticker)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", price), nil
}
