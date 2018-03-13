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


func TestUpdateEveryWeekDayTask(t *testing.T)  {
	timestr := time.Now().Format("2006-01-02")
	tt, _ := time.Parse("2006-01-02",timestr)
	timeUnix := tt.Unix()
	t2 := timeUnix - 60*60*8
	t1 := t2 - 60*60*24*7
	fmt.Println("------------t1,",t1,"----t2",t2)
	UpdateEveryWeekDayTask(int(t1),int(t2))
}

func TestGetPayLoadsByGameId(t *testing.T)  {
	timestr := time.Now().Format("2006-01-02")
	tt, _ := time.Parse("2006-01-02",timestr)
	timeUnix := tt.Unix()
	t2 := timeUnix - 60*60*8
	t1 := t2 - 60*60*24*7
	fmt.Println("------------t1,",t1,"----t2",t2)
	ps , err := GetPayLoadsByGameId("1517931441",int(t1),int(t2))
	if err != nil{
		fmt.Println("------------err",err)
		return
	}
	fmt.Println("------------ps",ps)
}

func TestGetPayLoadsByTime(t *testing.T)  {
	timestr := time.Now().Format("2006-01-02")
	tt, _ := time.Parse("2006-01-02",timestr)
	timeUnix := tt.Unix()
	t2 := timeUnix - 60*60*8
	t1 := t2 - 60*60*24*7
	fmt.Println("------------t1,",t1,"----t2",t2)
	ps , err := GetPayLoadsByTime(int(t1),int(t2))
	if err != nil{
		fmt.Println("------------err",err)
		return
	}
	fmt.Println("------------ps",ps)
	for i, v := range ps {
		fmt.Println("------------i:",i,"=======v:",v)
	}
}
