package interfacer

// ISRID 获取几何对象的WKT描述
type ISRID interface {
	SRID(geom interface{}) int64
}
