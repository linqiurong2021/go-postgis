package interfacer

// IRelate 获得两个几何对象的关系（DE-9IM矩阵）
type IRelate interface {
	Relate(geom1 interface{}, geom2 interface{}) bool
}
