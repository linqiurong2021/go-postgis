package interfacer

// INumGeometries 获取多几何对象中的对象个数
type INumGeometries interface {
	NumGeometries(geom interface{}) int64
}
