package interfacer

// ICentroid 获取几何对象的中心
type ICentroid interface {
	Centroid(geom interface{}) interface{}
}
