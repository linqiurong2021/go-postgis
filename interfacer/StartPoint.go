package interfacer

// IStartPoint 获取线的起始点
type IStartPoint interface {
	StartPoint(geom interface{}) int64
}
