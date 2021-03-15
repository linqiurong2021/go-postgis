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
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Centerid(params.Geom, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// GetArea GetArea
func (p *GisTool) GetArea(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeomRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Area(params.Geom, params.Type, params.FromSID, params.ToSID)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// GetLength GetLength
func (p *GisTool) GetLength(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeomRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Length(params.Geom, params.Type, params.FromSID, params.ToSID)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// PointOnSurface PointOnSurface
func (p *GisTool) PointOnSurface(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.PointOnSurface(params.Geom, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// Boundary Boundary
func (p *GisTool) Boundary(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Boundary(params.Geom, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// Buffer Buffer
func (p *GisTool) Buffer(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisBufferRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Buffer(params.Geom, params.Type, params.Distance)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// Intersection Intersection
func (p *GisTool) Intersection(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeom2Request(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Intersection(params.Geom, params.Geom2, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// ShiftLongitude ShiftLongitude
func (p *GisTool) ShiftLongitude(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeomRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.ShiftLongitude(params.Geom, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// SymDifference SymDifference
func (p *GisTool) SymDifference(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeom2Request(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.SymDifference(params.Geom, params.Geom2, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// Difference Difference
func (p *GisTool) Difference(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeom2Request(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Difference(params.Geom, params.Geom2, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// Union Union
func (p *GisTool) Union(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeom2Request(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.Union(params.Geom, params.Geom2, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}

// Union Union
func (p *GisTool) MemUnion(rw http.ResponseWriter, r *http.Request) {
	// geom
	// 获取请求过来的参数
	params, err := libs.NewPostGisGeomRequest(r)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	// 设置头部信息
	rw.Header().Set("Content-Type", "application/json")
	//
	result, err := p.Logic.PostGis.MemUnion(params.Geom, params.Type)
	if err != nil {
		json, _ := response.Error(err.Error(), nil)
		// 日志记录
		fmt.Fprint(rw, json)
	}
	json, _ := response.Success("Success", result)
	fmt.Fprint(rw, json)
}
