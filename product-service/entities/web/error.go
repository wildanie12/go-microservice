package web

import "github.com/wildanie12/go-microservice/product-service/config"

// Error custom error handler struct
type Error struct {
	Code int
	DevMessage  string
	ProdMessage string
	Message     string
}

// Error returns error message based on 
// what environment type that currently in set
func (err Error) Error() string {
	env := config.New().App.Env
	if env == "development" {
		return map[bool]string{ true: err.DevMessage, false: err.Message }[err.DevMessage != ""]
	}
	return map[bool]string{ true: err.ProdMessage, false: err.Message }[err.ProdMessage != ""]
}