package interfacer

// IEnvelope 获取几何对象的边界范围
type IEnvelope interface {
	Envelope(geom interface{}) interface{}
}
