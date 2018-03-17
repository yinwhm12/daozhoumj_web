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

func TestDiKui(t *testing.T)  {
	 testM := make(map[string]int,15)
	 //原始数据
	 testM["a"] = 1
	 testM["b"] = 2
	 testM["c"] = 3
	 testM["d"] = 10
	 testM["e"] = 11
	 testM["f"] = 2
	 testM["g"] = 3
	 testM["h"] = 4
	 testM["w"] = 6
	 testM["z"] = 7
	 testM["x"] = 5
	 testM["y"] = 9
	 testM["o"] = 10
	 testM["m"] = 12
	 testM["n"] = 5
	 fmt.Println("-----------ff:",testM["ww"])

}
