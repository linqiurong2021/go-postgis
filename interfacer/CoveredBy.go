package interfacer

// ICoveredBy 判断A是否被B所覆盖
type ICoveredBy interface {
	CoveredBy(A interface{}, B interface{}) bool
}
