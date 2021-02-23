package interfacer

// IContains 判断A是否包含B
type IContains interface {
	Contains(A interface{}, B interface{}) bool
}
