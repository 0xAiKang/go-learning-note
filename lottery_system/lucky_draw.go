package main

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"time"
)

/**
	GetAllAwardBatch 获取所有的awardBatch
	返回类型为结构
 */
func GetAllAwardBatch() []AwardBatch {
	// 声明一个切片
	var awardBatches []AwardBatch

	// 获取Redis 连接
	conn, err := getConn()
	if err != nil {
		fmt.Println("get redis conn error: ", err)
		return awardBatches
	}
	defer conn.Close()

	awardInfoKey := getAwardInfoKey()
	// 获取Sorted Set 所有键值
	values, err := redis.Values(conn.Do("ZRANGE", awardInfoKey, 0, -1, "WITHSCORES"))
	if err != nil || len(values) == 0{
		fmt.Println("get all award error: ", err)
		return awardBatches
	}

	// 遍历结果
	for index, value := range values{
		// 根据索引进行取余
		// 0%2 = 0、1%2 = 1、2%2 = 0
		if index%2 ==0  {
			// 使用断言的写法，判断数据类型
			awardName, ok := value.([]byte)
			if !ok {
				fmt.Println("value type error: ", value)
				continue
			}
			// 向切片中追加内容
			// A、B、C
			awardBatches = append(awardBatches, AwardBatch{name: string(awardName),})
		}else {
			lastUpdateTimeStr, ok := value.([]byte)
			if !ok {
				fmt.Println("time type error: ", value)
				continue
			}
			lastUpdateTime, err := strconv.ParseInt(string(lastUpdateTimeStr), 10, 64)
			if err != nil {
				fmt.Println("time type error: ", err)
			}
			awardBatches[index/2].SetUpdateTime(lastUpdateTime)
		}
	}

	// 遍历奖励池
	// 0、1、2
	for index, awardBatch := range awardBatches {
		// 添加总奖励金额
		awardBatches[index].SetTotalAmount(Conf.AwardMap[awardBatch.GetName()])
	}

	return awardBatches
}

/**
	RandomGetAwardBatch 从奖励池中选择一个随机奖励
 */
func RandomGetAwardBatch() (*AwardBatch, error)  {
	// 获取Redis 连接
	conn, err := getConn()
	if err != nil {
		fmt.Println("get conn is nil：", err)
		return nil, err
	}
	// 关闭连接
	defer conn.Close()
	
	// 获取奖励池中剩余的奖励
	retMap, err := redis.Int64Map(conn.Do("HGETALL", getAwardBalanceKey()))
	if err != nil || retMap == nil {
		fmt.Println("redis HGETALL award error：", err)
		return nil, err
	}
	// 统计总奖励
	totalBalance := int64(0)
	// 遍历Map，直接累加value
	for _, balance := range retMap{
		totalBalance += balance 
	}
	
	fmt.Println("retMap: ", retMap)

	if totalBalance == 0{
		fmt.Println("total balance is 0")
		return nil, errors.New("total balance is 0")
	}
	
	fmt.Println("total balance is: ", totalBalance)

	// 获取所有奖励
	awardBatches := GetAllAwardBatch()

	// 统计剩余奖励
	for index, awardBatch := range awardBatches{
		awardBatches[index].totalBalance = retMap[awardBatch.GetName()]
	}

	fmt.Println("awardBatches: ", awardBatches)

	// 创建一个 random
	random := rand.New(rand.NewSource(totalBalance))
	// 生成一个[0, totalBalance) 的随机数
	num := random.Int63n(totalBalance)

	// 遍历切片
	for _, awardBatch := range awardBatches{
		// 奖品已经派走了
		if awardBatch.GetTotalBalance() > 0 {
			fmt.Println("the prizes have been drawn ...")
		}

		num = num - awardBatch.GetTotalBalance()
		if num < 0 {
			return &awardBatch, nil
		}
	}

	return nil, nil
}

/**
	GetAwardBatch 实现了一个特定的幸运抽奖算法
 */
func GetAwardBatch() *AwardBatch {
	awardBatch, err := RandomGetAwardBatch()

	if awardBatch == nil || err != nil {
		fmt.Println("sorry, you don't win the prize")
	}

	// 确认是否已经达到奖品发放时间点

	// 获取设定开始时间时间戳
	startTime, _ := ParseStringToTime(Conf.Award.StartTime)
	endTime, _ := ParseStringToTime(Conf.Award.EndTime)


	totalAmount := awardBatch.totalAmount
	totalBalance := awardBatch.totalBalance
	lastUpdateTime := awardBatch.GetUpdateTime()

	// 创建一个random
	random := rand.New(rand.NewSource(lastUpdateTime))

	// 计算时间（开始时间 - 结束时间）/剩余奖品数
	dateTime := (endTime - startTime) / awardBatch.totalAmount

	// 计算下一个派奖时间
	releaseTime := startTime + (totalAmount - totalBalance)*dateTime + int64(random.Int())%dateTime

	fmt.Println("releaseTime :", time.Unix(releaseTime, 0).Format("2006-01-02 15:04:05"))

	// 比较时间，如果当前时间小于派奖时间，则没有中奖
	if time.Now().Unix() < releaseTime {
		fmt.Println("sorry, you don't win the prize")
		return nil
	}
	return awardBatch
}

// 确认用户是否中奖
func GetAward(username string) *AwardBatch {
	awardBatch := GetAwardBatch()

	if awardBatch == nil {
		return awardBatch
	}

	// 更新最后更新时间和余额
	conn, err := getConn()

	if err != nil {
		fmt.Println("get conn is nil", err)
		return nil
	}
	defer conn.Close()

	// 保证操作原子性
	conn.Send("WATCH", getAwardBalanceKey())
	// 开启事务
	conn.Send("MULTI")
	// 更新最后一次中奖时间
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), awardBatch.GetName())
	// 更新奖品数量
	conn.Send("HSET", getAwardBalanceKey(),  awardBatch.GetName(), awardBatch.totalBalance-1)
	// 提交事务
	conn.Send("EXEC")
	// 执行
	err = conn.Flush()
	if err != nil {
		fmt.Println("redis error: ", err)
		return nil
	}

	// 中奖
	fmt.Println("congratulations, you won ", awardBatch.GetName())

	// 中奖时间
	awardTime := time.Unix(awardBatch.GetUpdateTime(), 0 ).Format("2006-01-02 15:04:05")

	// 创建一个协程
	go func() {
		if err := SaveRecords(awardBatch.GetName(), awardTime, username); err != nil {
			fmt.Printf("save records error: %v", err)
		}
	}()
	return awardBatch
}

/**
  初始化Redis
 */
func InitAwardRedis() error {
	// 获取连接
	conn, err := getConn()
	if err != nil {
		fmt.Printf("get redis conn is nil, %v \n", err)
		return err
	}
	// 关闭连接
	defer conn.Close()

	// 初始化Map
	// score，value
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "A")
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "B")
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "C")

	// key，value
	conn.Send("HSET", getAwardBalanceKey(), "A", Conf.AwardMap["A"])
	conn.Send("HSET", getAwardBalanceKey(), "B", Conf.AwardMap["B"])
	conn.Send("HSET", getAwardBalanceKey(), "C", Conf.AwardMap["C"])

	// 执行
	conn.Flush()

	for i := 0; i <6; i ++ {
		// 接收返回
		_, err := conn.Receive()
		if err != nil {
			fmt.Println("conn send error: ", err)
		}

	}
	return nil
}
