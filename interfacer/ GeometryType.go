package interfacer

// IGeometryType 获取几何对象的类型
type IGeometryType interface {
	GeometryType(geom1 interface{}) float64
}
