package interfacer

// IAsText 获取几何对象的WKT描述
type IAsText interface {
	AsText(geom interface{}) interface{}
}
