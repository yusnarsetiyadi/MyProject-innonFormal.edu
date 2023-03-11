package feedback

/*

	Interface Application

*/

type ServiceInterface interface {
	GetAll() (data []FeedbackCore, err error)
	Create(input FeedbackCore) error
	GetById(id uint) (data FeedbackCore, err error)
	Update(input FeedbackCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []FeedbackCore, err error)
	Create(input FeedbackCore) error
	GetById(id uint) (data FeedbackCore, err error)
	Update(input FeedbackCore, id uint) error
	Delete(id uint) error
}
