package interfacer

// IDWithin 如果两个几何对象间距离在给定值范围内，则返回TRUE
type IDWithin interface {
	DWithin(geom1 interface{}, geom2 interface{}, distance float64) bool
}
