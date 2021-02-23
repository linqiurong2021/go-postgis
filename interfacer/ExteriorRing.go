package interfacer

// IExteriorRing 获取多几何对象中的对象个数
type IExteriorRing interface {
	ExteriorRing(geom interface{}) int64
}
