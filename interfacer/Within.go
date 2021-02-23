package interfacer

// IWithin 判断A是否被B包含
type IWithin interface {
	Within(A interface{}, B interface{}) bool
}
