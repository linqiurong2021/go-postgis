package interfacer

// IIsSimple 判断几何对象是否不包含特殊点（比如自相交
type IIsSimple interface {
	IsSimple(geom interface{}) bool
}
