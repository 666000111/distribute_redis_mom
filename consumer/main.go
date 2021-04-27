package main

import (
	"fmt"
	"github.com/luci/go-render/render"
	"math"
	"strconv"
)

func Init() {
	InitClient()
}

var (
	TIMES    = 100
	N        = 50
	EXCEPT   = float64(0)
	STANDARD = float64(1)
)

type Statistics struct {
	Sum     float64
	Average float64
	Data    []float64
}

func main() {
	Init()
	statistics := &Statistics{
		Sum:     0,
		Average: 0,
		Data:    make([]float64, 0),
	}
	for i := 0; i < TIMES; i++ {
		lenth := redisClient.LLen("norm:queue").Val()
		if lenth == 0 {
			continue
		}
		f64, err := consumeNorm()
		if err != nil {
			continue
		}
		statistics.Sum += f64
		statistics.Data = append(statistics.Data, f64)
		statistics.Average = statistics.Sum / float64(len(statistics.Data))
		fmt.Printf("add %f current statistics: sum:%f, average:%f count:%d\n", f64, statistics.Sum, statistics.Average, len(statistics.Data))
	}
	caculateAndPrint(statistics)
}

func consumeNorm() (float64, error) {
	str := redisClient.RPop("norm:queue").Val()
	f64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("convert string to float64 failed str = %s, err = %+v\n",
			str, render.Render(err))
		return 0, err
	}
	return f64, err
}

func caculateAndPrint(s *Statistics) {
	array := s.Data[TIMES-N:]
	sum := float64(0)
	for _, v := range array {
		sum += v
	}
	aver := sum / float64(N)
	s2 := float64(0)
	for _, v := range array {
		s2 += math.Pow(v-aver, 2)
	}
	fmt.Printf("s2 = %f\n", s2)
	abnormal := make([]float64, 0)
	for _, v := range s.Data {
		if math.Abs(v-EXCEPT) > 3*STANDARD {
			abnormal = append(abnormal, v)
		}
	}
	fmt.Printf("abnormal list = %+v", render.Render(abnormal))

}
