package interfacer

// IGeometryN 获取多几何对象中第N个对象
type IGeometryN interface {
	GeometryN(geom interface{}, index int64) interface{}
}
