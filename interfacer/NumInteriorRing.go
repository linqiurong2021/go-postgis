package interfacer

// INumInteriorRing 获取多边形内边界个数
type INumInteriorRing interface {
	NumInteriorRing(geom interface{}) int64
}
