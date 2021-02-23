package interfacer

// ICovers 判断A是否覆盖
type ICovers interface {
	Covers(A interface{}, B interface{}) bool
}
