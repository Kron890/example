package app

import (
	"template/config"
	"template/internal/example_module/delivery"
	"template/internal/example_module/repository/external_api"
	"template/internal/example_module/repository/postgres"
	"template/internal/example_module/usecase"
)

// Здесь мы инициализируем все наше приложение от коннекта к базе данных,
// до всех модулей твоего приложение
func (s *Server) InitApp(/* server *Server, */_ *config.Config) error {
	//подключения к базе данные
	// dbConnection, err := database.NewDbConnection("")
	//if err != nil {
	//	return err
	//}
	// новый репозиторий для работы с бд 
	repo := postgres.NewRepo(nil/* dbConnection */)
	// взаимодействия с внешним апи
	externalAPi := external_api.NewApi()
	// отправляем все в структуру для работы usecase 
	tickerUseCase := usecase.NewUseCase(repo, externalAPi)
	//отправляем все в структуру для работы в delivery
	handler := delivery.NewHandler(tickerUseCase)

	delivery.MapRoutes(s.AddHandler, handler)

	return nil
}
