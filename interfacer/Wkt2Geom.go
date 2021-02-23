package interfacer

// IGeom2Wkt geometry转wkt
type IGeom2Wkt interface {
	Geom2Wkt(wkt interface{}) interface{}
}
