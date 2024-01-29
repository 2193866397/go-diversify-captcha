package utils

import (
	secureRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

/*
 *@author: 随风飘的云
 *@description: 随机数工具
 *@date: 2023-10-02 23:16
 */
// 获取伪随机数种子
func getRand() *rand.Rand {
	// 使用真随机种子解决高并发情况下设置的伪随机数时间戳种子一样的bug
	return rand.New(rand.NewSource(GetSecureRandInt(time.Now().UnixNano())))
}

// 返回一个真随机数生成[0, max)的整数
func GetSecureRandInt(max int64) int64 {
	secureMax := new(big.Int).SetInt64(int64(max))
	res, err := secureRand.Int(secureRand.Reader, secureMax)
	if err != nil {
		fmt.Printf("Can't generate random value: %v, %v", res, err)
	}
	return res.Int64()
}

// 返回一个真随机数按进制算len长度的整数，注意len必须要大于等于2，比如3为2位
func GetSecureLenRand(len int) int64 {
	res, err := secureRand.Prime(secureRand.Reader, len)
	if err != nil {
		fmt.Printf("Can't generate random value because the len is < 2")
	}
	return res.Int64()
}

// 切片生成随机数

func GetSecureRandIntValue(n int) byte {
	val := make([]byte, n)
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	n, err := secureRand.Read(val)
	if err != nil {
		fmt.Print("Can't generate securerandom value")
	}
	return val[n]
}

// 取值范围在[0,n)的伪随机int值， 如果n<=0会panic
func GetRandomInt(n int) int {
	return getRand().Intn(n)
}

// 取值范围在[0,n)的伪随机int值， 如果n<=0会panic
func GetRandomInt31n(n int32) int32 {
	return getRand().Int31n(n)
}

/**
 * @description: 返回一个取值范围在[0,n)的伪随机int64值， 如果n<=0会panic。
 * @param {int64} n
 * @return {*}
 */
func GetRandomInt63n(n int64) int64 {
	return getRand().Int63n(n)
}

/**
 * @description: 返回一个取值范围在[0.0, 1.0)的伪随机float32值。
 * @return {*}
 */
func GetRandomFloat() float32 {
	return getRand().Float32()
}

/**
 * @description: 返回一个取值范围在[0.0, 1.0)的伪随机float64值。
 * @return {*}
 */
func GetRandomFloat64() float64 {
	return getRand().Float64()
}

/**
 * @description: 返回一个有n个元素的， [0,n)范围内整数的伪随机排列的切片。
 * @param {int} n
 * @return {*}
 */
func GetContextLen(n int) []int {
	return getRand().Perm(n)
}

/**
 * @description: 返回[min, max)范围的随机数
 * @param {*} min
 * @param {int} max
 * @return {*}
 */
func GetRandBetween(min, max int) int {
	return getRand().Intn(max-min) + min
}
