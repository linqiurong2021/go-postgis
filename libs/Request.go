package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request Request
type Request struct{}

// PostGisRequest PostGisRequest
type PostGisRequest struct {
	Geom string `json:"geom"`
	Type string `json:"type"`
}

// PostGisGeomRequest 地理坐标系请求
type PostGisGeomRequest struct {
	PostGisRequest
	FromSID int64 `json:"from_sid"`
	ToSID   int64 `json:"to_sid"`
}

// PostGisBuffRequest 缓冲请求
type PostGisBuffRequest struct {
	PostGisRequest
	Distance float64 // 缓冲长度 单位m
}

// PostGisRequest PostGisRequest
type PostGisGeom2Request struct {
	PostGisRequest
	Geom2 string `json:"geom2"`
}

// NewPostGisRequest NewPostGisRequest 转 MAP
func NewPostGisRequest(r *http.Request) (request *PostGisRequest, err error) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return
}

// NewPostGisGeomRequest NewPostGisGeomRequest 转 MAP
func NewPostGisGeomRequest(r *http.Request) (request *PostGisGeomRequest, err error) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return
}

// NewPostGisBufferRequest NewPostGisBufferRequest
func NewPostGisBufferRequest(r *http.Request) (request *PostGisBuffRequest, err error) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return
}

// NewPostGisGeom2Request NewPostGisGeom2Request
func NewPostGisGeom2Request(r *http.Request) (request *PostGisGeom2Request, err error) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return
}

// NewRequestToMap Request 转 MAP
func NewRequestToMap(r *http.Request) (map[string]interface{}, error) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}
	// 参数
	params := make(map[string]interface{})
	//
	if err := json.Unmarshal(body, &params); err != nil {
		return nil, err
	}
	return params, nil
}
