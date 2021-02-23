package interfacer

// IEndPoint 获取线的起终点
type IEndPoint interface {
	EndPoint(geom interface{}) int64
}
