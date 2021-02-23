package interfacer

// IShiftLongitude 将经度小于0的值加360使所有经度值在0-360间
type IShiftLongitude interface {
	ShiftLongitude(geom interface{}) interface{}
}
