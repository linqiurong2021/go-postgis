package interfacer

// ILength 计算线的长度
type ILength interface {
	Length(geom1 interface{}) float64
}
