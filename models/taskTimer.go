package models

import "sync"

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
			originData.Unlock()
			originData.userPayLoadMap[v.GameId] = sum
		}
		if len(users) < limit{
			break //没有数据了
		}
		i = (i+1) * limit
		
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
func (o *OriginData)SumPayLoad()  {
	
}

func DiKui(gameId string)(int)  {
	myclient, err := GetMyClientByGameId(gameId)
	if err != nil{

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