package main

import (
	"fmt"

	"github.com/linuxops-br/instachecker/pkg/instachecker"
)

func main() {
	ic := instachecker.NewInstagramUser("therock")

	fmt.Printf(
		"Nome: %s\nUsuário: %s",
		ic.GetName(),
		ic.GetPicture().Large,
	)
}
