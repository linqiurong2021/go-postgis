package interfacer

// IBoundary 获取边界
type IBoundary interface {
	Boundary(geom interface{}) interface{}
}
