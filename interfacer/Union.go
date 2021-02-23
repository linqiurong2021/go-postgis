package interfacer

// IUnion 返回两个几何对象的合并结果
type IUnion interface {
	// 返回两个几何对象的合并结果
	Union(geom1 interface{}, geom2 interface{}) interface{}
	// 返回两个几何对象的合并结果集
	UnionSet(geom interface{}) interface{}
	// 用较少的内存和较长的时间完成合并操作，结果和Union相同
	MemUnion(geom interface{}) interface{}
}
