package interfacer

// IIntersects 判断两个几何对象是否分离
type IIntersects interface {
	Intersects(geom1 interface{}, geom2 interface{}) bool
}
