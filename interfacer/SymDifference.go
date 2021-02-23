package interfacer

// ISymDifference 获取两个几何对象不相交的部分
type ISymDifference interface {
	SymDifference(geom1 interface{}, geom2 interface{}) interface{}
}
