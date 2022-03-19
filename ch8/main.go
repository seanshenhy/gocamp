package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)
var rdb = redis.NewClient(
	&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
})
var ctx = context.Background()
var n int64 = 20000
var isString bool
func main() {
	write2wKV()
	getRdsInfoStat()
}
func getRdsInfoStat()  {
	arr := strings.Split(rdb.Info(ctx, "memory").Val(), "\r\n")

	for i,v :=range arr{
		if i == 0 {
			continue
		}
		vv := strings.Split(v, ":")
		if len(vv) != 2 {
			continue
		}
		if vv[0] == "used_memory" {
			value, _ := strconv.Atoi(vv[1])
			fmt.Printf("总共有%d个key，%s字节（%d k），平均每个key占%d字节\n", n, vv[1], value / 1024, int64(value) / n)
		}
	}
}
func write2wKV()  {
	rand.Seed(time.Now().UnixNano())
	n := 20000
	min := math.MaxInt64
	max := math.MinInt64
	//n := 10
	for i := 1; i <= n; i++{
		key := fmt.Sprintf("key_%d", i)
		tv := getValue()
		switch tv.(type) {
		case string:
			value := tv.(string)
			if len(value) < min {
				min = len(value)
			}
			if len(value) > max {
				max = len(value)
			}
		}
		//value := rand.Int63n(100000000000)
		//fmt.Println(key,value)
		err := rdb.Set(ctx, key, tv, 0).Err()
		if err != nil {
			fmt.Println("写入失败",i)
			continue
		}
	}
	fmt.Println("min:",min,",max:",max)
	fmt.Println("写入成功")
}

func getValue()interface{} {
	if isString {
		return fmt.Sprintf("value_%d", rand.Int63n(100000000000))
	}
	return rand.Int63n(100000000000)
}
