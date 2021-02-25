package api

import "fmt"

// RESTfulAPI RESTfulAPI
type RESTfulAPI struct {
	GisTool
}

// NewRESTfulAPI NewRESTfulAPI
// func NewRESTfulAPI() *RESTfulAPI {
// 	return &RESTfulAPI{
// 		Tool: NewGisTool(),
// 	}
// }

// Get Get
func (restful *RESTfulAPI) Get() {
	fmt.Print("API")
}
