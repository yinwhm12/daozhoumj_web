package models

import (
	"fmt"
	"testing"
	"log"
)

func TestUserI(t *testing.T)  {
	user, err := GetUserByGameId("1517722155")
	if err != nil{
		log.Fatal("db err ",err)
		return
	}
	fmt.Println(*user)

}
