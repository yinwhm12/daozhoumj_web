package models

import (
	"sync"
	"time"
)

type OriginData struct {
	//记录原始数据
	mux sync.RWMutex
	userPayLoadMap map[string]int
}

var originData = struct {
	sync.RWMutex
	userPayLoadMap map[string]int
}{userPayLoadMap:make(map[string]int)}

//init
func newOriginData(limit int)(*OriginData)  {
	originData := new(OriginData)
	originData.userPayLoadMap = make(map[string]int,limit)
	return originData
}

//定时任务
//每周一的2点开始 统计 每周的业务
func UpdateEveryWeekDayTask(t1,t2 int)  {
	//获取所有的玩家自己的贡献
	i := 0
	limit := 100
	//originData := new(OriginData)
	//originData.userPayLoadMap = make(map[string]int,100)
	for{
		users, err :=  GetLimitMoreUser(i,limit)
		if err != nil{
			//todo 失败
			break
		}
		for _, v:= range users{
			ps, err := GetPayLoadsByGameId(v.GameId,t1,t2)
			if err !=nil {
				//todo 失败 需要处理
				continue
			}
			sum := 0
			for _, vv := range ps {
				sum += vv.Asset //统计
			}
			originData.Lock()
			originData.userPayLoadMap[v.GameId] = sum
			originData.Unlock()
		}
		if len(users) < limit{
			break //没有数据了
		}
		i = (i+1) * limit
		
	}
	ComputeGamer()

	//创建每个人的数据

	for k,v := range originData.userPayLoadMap{
		ac := Achiment{
			Id:k,
			Achievements:0,
			CreateTime:int(time.Now().Unix()),
		}
		if c, ok := sumData.sum[k]; ok{
			ac.Commision = c
			if c < 10000{

			}
		}
	}


}

//type SumData struct {
//	//统计的数据
//	mux sync.RWMutex
//	payLoadMap map[string]int
//}

var sumData = struct {
	sync.RWMutex
	sum map[string]int
}{sum:make(map[string]int)}


//计算 总量
func ComputeGamer()  {
	sumData.sum = make(map[string]int,len(originData.userPayLoadMap))
	for k,_ := range originData.userPayLoadMap{
		if _, ok := sumData.sum[k]; !ok{
			sumData.Lock()
			sumData.sum[k] = DiKui(k)
			sumData.Unlock()
		}
	}
}


func DiKui(gameId string)(int)  {
	myclient, err := GetMyClientByGameId(gameId)
	if err != nil{
		return 0
	}
	if myclient == nil || len(myclient.Sons) == 0{
		return originData.userPayLoadMap[gameId]
	}else{
		for _, v := range myclient.Sons{
			sumData.sum[gameId]+=DiKui(v)
		}

		return sumData.sum[gameId]
		//return sumData[]
	}
}

func getBi(assert int)float32  {
	if assert < 10000{
		//return (bi - 0.006) * float32(assert)
		return 0

	}else if assert <=30 *10000{

		return 0.006 + 0.01
	}else if assert <=100 *10000{

		return 0.008 + 0.01
	}else if assert <=200 * 10000{

		return 0.01 + 0.01
	}else if assert <= 500 * 10000{

		return 0.013 + 0.01
	}else if assert <= 100 *10000 {
		return 0.016 + 0.01
	}else{
		return 0.02 + 0.01
	}
}

//按比例计算 直推开始
func ResultTheWeek(gameId string)(float32)  {
	var  sum float32
	bi := getBi(sumData.sum[gameId])
	if bi == 0{
		//一万的业绩没达到
		return 0
	}else{
		//说明达到要求
	}
	myclient, err := GetMyClientByGameId(gameId)
	if err != nil{
		return 0
	}
	if myclient == nil || len(myclient.Sons) == 0{
		return 0
	}else{

		for _, v := range myclient.Sons{
			//if vv, ok :=originData.userPayLoadMap[v]; ok{
				sum += float32(sumData.sum[v]) * bi
				sum += ResultOne(v,bi)
				/*

				代理制度表（每周返累计制）
				 代理级别 业绩（游戏流水） 佣金%
				1 主任（班长） 1万-30万 0.6
				2 高级主任（排长） 30万-100万 0.8
				2 经理（连长） 100万-200万 1.0
				4 高级经理（营长） 200万-500万 1.3
				5 总监（团长） 500万-1000万 1.6
				6 高级总监（师长） 1000万以上 2.0

				例如:本周你有600万成绩是总监（团长），其中直推30万成绩，你的团队会员小张10万成绩是主任（班长），小李60万成绩是高级主任（排长），小王150万成绩是经理（连长小刘350万成绩是高级经理（营长）
			  你本周的收入是 ?
			  你直推佣|金: 1.6x30万=4800
			  你到小张级差佣|金:(1.6-0.6)x10万=1000
			  你到小李级差佣|金:( 1.6-0.8)x60万=4800
			  你到小王级差佣|金:( 1.6-1.0)x150万=9000
			  你到小刘级差佣|金:( 1.6-1.3)x350万=10500
			  合计: 4800+1000+4800+9000+10500=30100
				 */
		}
		return sum
	}
}

//递归 算每个人 的贡献 //间接的 不是直推
func ResultOne(gameId string,bi float32)(float32)  {
	myclient, err := GetMyClientByGameId(gameId)
	if err != nil{
		return 0
	}
	if myclient == nil || len(myclient.Sons) == 0{
		return SwitCompute(bi,originData.userPayLoadMap[gameId])
	}else{
		var sum float32
		for _, v := range myclient.Sons{
			//sum += float32(sumData.sum[v])
			sum += SwitCompute(bi,originData.userPayLoadMap[v])
			sum += ResultOne(v,bi)
		}
		return sum
	}
}

func SwitCompute(bi float32,assert int)(float32)  {
	if assert < 10000{
		//return (bi - 0.006) * float32(assert)
		return 0

	}else if assert <=30 *10000{

		return (bi - 0.006) * float32(assert)
	}else if assert <=100 *10000{

		return (bi - 0.008) * float32(assert)
	}else if assert <=200 * 10000{

		return (bi - 0.01) * float32(assert)
	}else if assert <= 500 * 10000{

		return (bi - 0.013) * float32(assert)
	}else if assert <= 100 *10000 {
		return (bi - 0.016) * float32(assert)
	}else{
		return (bi - 0.02) * float32(assert)
	}
}