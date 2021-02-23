package interfacer

// IOverlaps 判断两个几何对象是否是重叠
type IOverlaps interface {
	Overlaps(geom1 interface{}, geom2 interface{}) bool
}
