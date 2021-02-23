package interfacer

// IWkt2Geom wkt转geometry
type IWkt2Geom interface {
	Wkt2Geom(wkt interface{}, wkid int) interface{}
}
