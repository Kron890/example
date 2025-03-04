package postgres

import "template/infrastructure/database"

type TickerRepository struct {
	db *database.DBConn
}
//dbConnection попадает сюда и переобразуется в TickerRepository
//следовательно мы работаем не на прямую с бд 
func NewRepo(db *database.DBConn) *TickerRepository {
	return &TickerRepository{
		db,
	}
}

func (r *TickerRepository) CheckTicker(ticker string) (bool, error) {
	var exists bool
	return exists, r.db.QueryRow("select exists(id) from ticker_storage where ticker = $1", ticker).Scan(&exists)
}
