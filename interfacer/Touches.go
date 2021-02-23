package interfacer

// ITouches 判断两个几何对象的边缘是否接触
type ITouches interface {
	Touches(geom1 interface{}, geom2 interface{}) bool
}
