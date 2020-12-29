package main

//  定义结构
type AwardBatch struct {
	// 奖励名称
	name string

	// 总奖励金额
	totalBalance int64

	// 总金额
	totalAmount	int64

	// 更新时间
	updateTime int64
}

// 给结构定义一个行为/方法
func (award *AwardBatch) SetUpdateTime(updateTime int64) {
	award.updateTime = updateTime
}

func (award *AwardBatch) SetTotalAmount(totalAmount int64) {
	award.totalAmount = totalAmount
}

func (award *AwardBatch) GetName() string {
	return award.name
}

func (award *AwardBatch) GetTotalBalance() int64 {
	return award.totalBalance
}

func (award *AwardBatch) GetUpdateTime() int64 {
	return award.updateTime
}