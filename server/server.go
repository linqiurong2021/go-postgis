package server

import "github.com/linqiurong2021/go-postgis/libs"

// NewServer NewServer
func NewServer(postgres *libs.Postgre) *RESTful {
	return &RESTful{
		postgres: postgres,
	}
}

// RESTful RESTful
type RESTful struct {
	postgres *libs.Postgre
}

// Run Run
func (r *RESTful) Run() {
	//
}
