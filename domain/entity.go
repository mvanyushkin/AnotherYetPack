package domain

type Entity struct {
	Id int
}

func NewEntity() Entity {
	return Entity{Id: 1}
}
