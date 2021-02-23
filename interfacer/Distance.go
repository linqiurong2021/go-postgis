package interfacer

// IDistance 计算两个坐标点距离
type IDistance interface {
	Distance(geom1 interface{}, geom2 interface{}) interface{}
}
