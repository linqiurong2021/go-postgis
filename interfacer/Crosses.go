package interfacer

// ICrosses 判断两个几何对象是否互相穿过
type ICrosses interface {
	Crosses(geom1 interface{}, geom2 interface{}) bool
}
