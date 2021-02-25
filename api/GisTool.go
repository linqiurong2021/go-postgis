package api

import (
	"fmt"
	"net/http"

	"github.com/linqiurong2021/go-postgis/libs"
	"github.com/linqiurong2021/go-postgis/libs/response"
	"github.com/linqiurong2021/go-postgis/logic"
)

// GisTool GisTool
type GisTool struct {
	Logic *logic.Logic
}

// NewGisTool NewGisTool
func NewGisTool() *GisTool {
	return &GisTool{
		Logic: logic.NewLogic(),
	}
}

// Centroid Centroid
func (p *GisTool) Centroid(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		fmt.Fprint(rw, json)
	}
	//
	result, err := p.Logic.PostGis.Centerid(params.Geom, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}
