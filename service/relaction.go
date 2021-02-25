package service

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/linqiurong2021/go-postgis/db"
)

// Relaction 几何对象关系
type Relaction struct {
	Pool    *pgxpool.Pool
	PostGis *db.PostGis
}

//NewRelaction NewRelaction
func NewRelaction() *Relaction {

	return &Relaction{
		PostGis: db.NewPostGis(),
	}
}

// Centroid Centroid
func (r *Relaction) Centroid(geometry interface{}, returnType string) (result interface{}, err error) {
	return r.PostGis.Centroid(geometry,returnType)
}
