package interfacer

// IDifference 从A去除和B相交的部分后返回
type IDifference interface {
	Difference(geom1 interface{}, geom2 interface{}) interface{}
}
