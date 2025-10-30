package domain

type Repository interface {
	Create(participant *Participant) error
	FindByID(id string) (*Participant, error)
	FindByUser(userID string) ([]Participant, error)
	FindByEvent(eventID string) ([]Participant, error)
	FindByEventAndUser(eventID, userID string) (*Participant, error)
	Update(participant *Participant) error
	Delete(id string) error

	CreatePermission(name, key string) error
	GetPermissionByKey(key string) (*Permission, error)
}
