package interfacer

// IIsRing 判断曲线是否闭合并且不包含特殊点
type IIsRing interface {
	IsRing(geom interface{}) bool
}
