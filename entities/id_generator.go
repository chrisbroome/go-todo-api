package entities

type IdGenerator interface {
	GenerateID() string
}
