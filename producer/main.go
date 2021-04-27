package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Init() {
	InitClient()
}

func main() {
	Init()
	for i := 0; i < 100; i++ {
		product, count := produceNorm()
		fmt.Printf("add %.3f to queue,current count = %d\n", product, count)
		time.Sleep(time.Millisecond * 10)
	}
}

func produceNorm() (float64, int64) {
	num := rand.NormFloat64()
	count := redisClient.LPush("norm:queue", num).Val()
	return num, count
}
