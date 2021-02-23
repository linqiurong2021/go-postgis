package interfacer

// IConvexHull 判断A是否包含B
type IConvexHull interface {
	ConvexHull(geom interface{}) interface{}
}
