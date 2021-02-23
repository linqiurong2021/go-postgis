package interfacer

// INumPoints 获取几何对象中的点个数
type INumPoints interface {
	NumPoints(geom interface{}) int64
}
