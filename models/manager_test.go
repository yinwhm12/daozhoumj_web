package models

import (
	"testing"
	"crypto"
	"io"
	"fmt"
	"time"
)

func TestMd5(t *testing.T)  {
	data := "333"
	md5 := crypto.MD5.New()
	io.WriteString(md5,data)
	fmt.Printf("%x\n",md5.Sum([]byte("yin")))

	str := md5.Sum(nil)
	tt := string(str)
	fmt.Println(tt)

}

func TestSaveM(t *testing.T)  {
	data := "yinwhm12"
	md5 := crypto.MD5.New()
	io.WriteString(md5,data)
	//str := md5.Sum([]byte("yin"))
	str := md5.Sum(nil)
	tt := string(str)

	m := Manager{
		Id:6,
		NickName:"admin",
		LastTime:int(time.Now().Unix()),
		Token:"dfjaldfjaldf",
		MDPwd:tt,
	}
	err := SaveM(&m)
	if err != nil{
		fmt.Println("err---",err)
		return
	}
	fmt.Println("ok")

}

func TestValidte(t *testing.T)  {
	nickName := "yin"
	data := "yin"
	md5 := crypto.MD5.New()
	io.WriteString(md5,data)
	//str := md5.Sum([]byte("yin"))
	str := md5.Sum(nil)
	tt := string(str)
	fmt.Println(tt)
	user, err := ValidateAndLogin(nickName,data)
	if err != nil{
		fmt.Println("---err find---",err)
		return
	}
	fmt.Println("user---",user)

}

func TestLastTime(t *testing.T)  {
	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,-1)
	fmt.Println("----------------y",yesTime.Unix())
	logDay := yesTime.Format("20060102")
	fmt.Println("-----------log:",logDay)
}
