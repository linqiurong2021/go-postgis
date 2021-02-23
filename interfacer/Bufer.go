package interfacer

// IBuffer 缓冲区计算
type IBuffer interface {
	Buffer(geom1 interface{}, distance float64) interface{}
}
