package example_module

type ModuleDatabase interface {
	CheckTicker(ticker string) (bool, error)
}

type ModuleExternalService interface {
	GetPrice(ticker string) (float64, error)
}
