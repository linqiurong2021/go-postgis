package interfacer

// IInteriorRingN 获取多边形的第N个内边界
type IInteriorRingN interface {
	InteriorRingN(geom interface{}, index int64) interface{}
}
