package models

import (
	"testing"
	"fmt"
)

func TestGetMyClientByGameId(t *testing.T)  {
	r, err := GetMyClientByGameId("15167118137");
	if err != nil{
		fmt.Println("err-----------:",err)
	}
	fmt.Println("Ok---:",r)
	fmt.Println("Ok--sons-:",r.Sons)
	fmt.Println("Ok--sosnl-:",len(r.Sons))
}
