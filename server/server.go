package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/linqiurong2021/go-postgis/api"
	"github.com/linqiurong2021/go-postgis/logic"
)

// RESTful RESTful
type RESTful struct {
	logic *logic.Logic
}

// NewServer NewServer
func NewServer() *RESTful {
	return &RESTful{}
}

// Run Run
func (r *RESTful) Run() {
	// api => logic => service =>
	gisAPI := api.NewGisTool()
	//
	http.HandleFunc("/gis/getCenter", gisAPI.Centroid)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello World!!!")
	})

	addr := ":8080"
	fmt.Printf("\n server start at %s \n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
}
