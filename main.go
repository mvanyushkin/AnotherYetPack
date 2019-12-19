package main

import (
	"AnotherYetPack/domain"
	"AnotherYetPack/services"
)

func main() {
	e := domain.NewEntity()
	services.ProcessEntity(e)
}
