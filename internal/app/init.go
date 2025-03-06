package app

import (
	"template/config"
	"template/infrastructure/database"
	"template/internal/example_module/delivery"
	"template/internal/example_module/repository/external_api"
	"template/internal/example_module/repository/postgres"
	"template/internal/example_module/usecase"
)

// Здесь мы инициализируем все наше приложение от коннекта к базе данных,
// до всех модулей твоего приложение
func (s *Server) InitApp( /* server *Server, */ _ *config.Config) error {
	//подключения к базе данные
	dbConnection, err := database.NewDbConnection("")
	if err != nil {
		return err
	}
	// Создание репозитория для работы с тикерами через PostgreSQL
	repo := postgres.NewRepo(dbConnection)

	// Инициализация клиента для взаимодействия с внешним API (заглушка)
	externalAPi := external_api.NewApi()

	// Создание слоя бизнес-логики (usecase) и передача в него зависимостей
	tickerUseCase := usecase.NewUseCase(repo, externalAPi)

	//отправляем все в структуру для работы в delivery
	// Инициализация обработчиков HTTP-запросов (delivery) и передача в них usecase
	handler := delivery.NewHandler(tickerUseCase)

	// Регистрация маршрутов в сервере
	delivery.MapRoutes(s.AddHandler, handler)

	return nil
}

/* Итог:
repo → работа с БД.
externalAPi → работа с внешним API.
tickerUseCase → связывает БД и API, обрабатывает данные.
handler → отвечает за HTTP-запросы.
MapRoutes → регистрирует маршруты сервера. */
