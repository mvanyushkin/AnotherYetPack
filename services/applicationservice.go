package services

import (
	"AnotherYetPack/domain"
	"strconv"
)

func ProcessEntity(e domain.Entity) {
	println("Process entity " + strconv.Itoa(e.Id))
}
