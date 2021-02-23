package interfacer

// INumInteriorRings 获取多边形内边界个数
type INumInteriorRings interface {
	NumInteriorRings(geom interface{}) int64
}
