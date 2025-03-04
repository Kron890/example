package example_module

type UseCase interface {
	ProcessTicker(ticker string) (string, error)
}
