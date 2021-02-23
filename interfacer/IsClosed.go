package interfacer

// IIsClosed 判断几何对象是否闭合
type IIsClosed interface {
	IsClosed(geom interface{}) bool
}
