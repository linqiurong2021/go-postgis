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

	http.HandleFunc("/gis/getCenter", gisAPI.Centroid)
	//
	http.HandleFunc("/gis/getArea", gisAPI.GetArea)
	//
	http.HandleFunc("/gis/getLength", gisAPI.GetLength)
	//
	http.HandleFunc("/gis/pointOnSurface", gisAPI.PointOnSurface)
	//
	http.HandleFunc("/gis/boundary", gisAPI.Boundary)
	//
	http.HandleFunc("/gis/buffer", gisAPI.Buffer)
	// Intersection
	http.HandleFunc("/gis/intersection", gisAPI.Intersection)
	// ShiftLongitude
	http.HandleFunc("/gis/shiftLongitude", gisAPI.ShiftLongitude)
	// SymDifference
	http.HandleFunc("/gis/symDifference", gisAPI.SymDifference)
	// Difference
	http.HandleFunc("/gis/difference", gisAPI.Difference)
	// Union
	http.HandleFunc("/gis/union", gisAPI.Union)
	// MemUnion
	http.HandleFunc("/gis/memUnion", gisAPI.MemUnion)

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
