package interfacer

// IPointN 获取几何对象的第N个点
type IPointN interface {
	PointN(geom interface{}, index int64) interface{}
}
