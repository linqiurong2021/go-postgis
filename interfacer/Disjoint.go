package interfacer

// IDisjoint 判断两个几何对象是否分离
type IDisjoint interface {
	Disjoint(geom1 interface{}, geom2 interface{}) bool
}
