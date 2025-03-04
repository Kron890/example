package external_api

//внешний апи, обращаемся к примеру бинанса для получения инофрмации

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (a Api) GetPrice(_ string) (float64, error) {
	return 10.0, nil
}
