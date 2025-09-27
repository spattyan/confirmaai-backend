package domain

type Repository interface {
	Create(event *Event) error
	FindByID(id string) (*Event, error)
	Update(event *Event) error
	Delete(id string) error
	List() ([]Event, error)
}
