package interfacer

// IPointOnSurface 返回曲面上的一个点
type IPointOnSurface interface {
	PointOnSurface(geom interface{}) interface{}
}
