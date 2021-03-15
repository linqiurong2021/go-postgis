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

// Area Area
func (pg *PostGis) Area(geom interface{}, coordinateType string, fromSID int64, toSID int64) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Area(geom, coordinateType, fromSID, toSID)
}

// Length Length
func (pg *PostGis) Length(geom interface{}, coordinateType string, fromSID int64, toSID int64) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Length(geom, coordinateType, fromSID, toSID)
}

// PointOnSurface PointOnSurface
func (pg *PostGis) PointOnSurface(geom interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.PointOnSurface(geom, returnType)
}

// Boundary Boundary
func (pg *PostGis) Boundary(geom interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Boundary(geom, returnType)
}

// Buffer Buffer
func (pg *PostGis) Buffer(geom interface{}, returnType string, distance float64) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Buffer(geom, returnType, distance)
}

// Intersection Intersection
func (pg *PostGis) Intersection(geom interface{}, geom2 interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Intersection(geom, geom2, returnType)
}

// ShiftLongitude ShiftLongitude
func (pg *PostGis) ShiftLongitude(geom interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.ShiftLongitude(geom, returnType)
}

// SymDifference SymDifference
func (pg *PostGis) SymDifference(geom interface{}, geom2 interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.SymDifference(geom, geom2, returnType)
}

// Difference Difference
func (pg *PostGis) Difference(geom interface{}, geom2 interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Difference(geom, geom2, returnType)
}

// Union Union
func (pg *PostGis) Union(geom interface{}, geom2 interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.Union(geom, geom2, returnType)
}

// Union Union
func (pg *PostGis) MemUnion(geom interface{}, returnType string) (result interface{}, err error) {
	return pg.Service.Relaction.PostGis.MemUnion(geom, returnType)
}
