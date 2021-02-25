package logic

// logic => service =>

// Logic Logic
type Logic struct {
	PostGis *PostGis
}

// NewLogic NewLogic
func NewLogic() *Logic {
	return &Logic{
		PostGis: NewPostGis(),
	}
}
