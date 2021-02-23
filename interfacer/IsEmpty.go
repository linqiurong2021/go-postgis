package interfacer

// IIsEmpty 判断几何对象是否为空
type IIsEmpty interface {
	IsEmpty(geom interface{}) bool
}
