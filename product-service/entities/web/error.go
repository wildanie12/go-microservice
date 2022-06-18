package web

import "github.com/wildanie12/go-microservice/product-service/config"

type WebError struct {
	Code int
	DevMessage  string
	ProdMessage string
	Message     string
}

func (err WebError) Error() string {
	env := config.New().App.Env
	if env == "development" {
		return map[bool]string{ true: err.DevMessage, false: err.Message }[err.DevMessage != ""]
	}
	return map[bool]string{ true: err.ProdMessage, false: err.Message }[err.ProdMessage != ""]
}