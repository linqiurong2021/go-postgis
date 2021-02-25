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
	Geom string
	Type string
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
