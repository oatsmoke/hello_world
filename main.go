package main

import (
	"fmt"
	"github.com/gofrs/uuid"
)

func main() {
	uid, _ := uuid.NewV4()
	fmt.Println(uid.String())
}
