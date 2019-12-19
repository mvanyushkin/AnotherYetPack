package domain


type Entity struct {
	id int
}

func (e Entity) Do() {
	println("Do something")
}

func NewEntity() Entity  {
	return Entity{id:1}
}
