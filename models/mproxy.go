package models

//代理

type Proxy struct {
	ID	string	`bson:"_id" json:"id,omitempty"`
	NickName	string	`bson:"nick_name" json:"nick_name,omitempty"`
	CardCounts	int	`bson:"card_counts" json:"card_counts,omitempty"` //现持卡数
	OldCardCounts	int	`bson:"old_card_counts" json:"old_card_counts,omitempty"` //历史交易卡数
	ProxyClass	int	`bson:"proxy_class" json:"proxy_class,omitempty"` //代理级别
	JoinTime	int	`bson:"join_time" json:"join_time,omitempty"`//加入代理时间

	BranchProxys	[]Proxy `bson:"branch_proxys" json:"branch_proxys,omitempty"`
	DealRecords	[]DealRecord `json:"deal_records,omitempty"`
}

//历史交易卡 数据 记录
type DealRecord struct {
	ID	string	`bson:"id" json:"id,omitempty"`
	NickName	string	`bson:"nick_name" json:"nick_name,omitempty"`
	DealType	int	`bson:"deal_type" json:"deal_type,omitempty"` //交易类型 1--为 // 出售 0为购入
	Counts	int	`bson:"counts" json:"counts,omitempty"`//交易数量
	DealState	int	`bson:"deal_state" json:"deal_state,omitempty"` //交易状态 1为成功 0为失败
	CreatedTime	int	`bson:"created_time" json:"created_time,omitempty"`//交易时间
}