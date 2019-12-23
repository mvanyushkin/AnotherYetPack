package main

import (
	"github.com/mvanyushkin/AnotherYetPack/domain"
	"github.com/mvanyushkin/AnotherYetPack/services"
)

func main() {
	e := domain.NewEntity()
	services.ProcessEntity(e)
}
