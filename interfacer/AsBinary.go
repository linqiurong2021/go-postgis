package interfacer

// IAsBinary 获取几何对象的WKB描述
type IAsBinary interface {
	AsBinary(geom interface{}) interface{}
}
