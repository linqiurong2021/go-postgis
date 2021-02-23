package interfacer

// IEquals 如果两个几何对象间距离在给定值范围内，则返回TRUE
type IEquals interface {
	Equals(geom1 interface{}, geom2 interface{}) bool
}
