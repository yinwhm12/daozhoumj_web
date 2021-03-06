package mongodb

import (
	"testing"
	"time"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

type UserTest struct {
	Id int "_id"
	Name string `json:"name"`
	CreateTime int `json:"create_time"`
}


func TestUserTest(t *testing.T)  {
	conn := Conn()
	defer conn.Close()

	user := UserTest{
		Id:2,
		Name:"yn",
		CreateTime:int(time.Now().Unix()),
	}
	c := conn.DB("ddzhu").C("test")
	c.Insert(&user)

}

func TestFindUser(t *testing.T)  {
	conn := Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("test")
	user := UserTest{}
	_ = c.Find(bson.M{"_id":1}).One(&user)
	fmt.Println(user)
}

func TestLast1Time(t *testing.T)  {
	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,-1)
	fmt.Println("----------------y",yesTime.Unix())
	logDay := yesTime.Format("20060102")
	fmt.Println("-----------log:",logDay)


}

//获取当天0点的时间戳
func TestZeroCurDay(t *testing.T)  {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	tt, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := tt.Unix()
	fmt.Println("timeNumber:", timeNumber)
	fmt.Println("------------timed",int(timeNumber))
	
}
// 对应的星期几 每3s运行一次
func TestCron(t *testing.T)  {
	i := 0
	c := cron.New()
	spec := "*/3 * * * * 2"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	time.Sleep(10 *time.Second)
	//select{}
}

