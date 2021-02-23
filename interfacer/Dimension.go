package interfacer

// IDimension 获取几何对象的维数
type IDimension interface {
	Dimension(geom interface{}) int64
}
