package logic

import (
	"github.com/linqiurong2021/go-postgis/service"
)

// PostGis PostGis
type PostGis struct {
	Service *service.Service
}

// NewPostGis NewPostGis
func NewPostGis() *PostGis {
	return &PostGis{
		Service: service.NewService(),
	}
}

// Centerid Centerid
func (pg *PostGis) Centerid(geom interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Centroid(geom, returnType)
}
