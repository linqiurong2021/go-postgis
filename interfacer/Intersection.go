package interfacer

// IIntersection 获取两个几何对象相交的部分
type IIntersection interface {
	Intersection(geom1 interface{}, geom2 interface{}) interface{}
}
