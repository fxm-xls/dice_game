package public

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitDiceNum 切分字符串骰子点数 分为 int点数、时间、key
func SplitDiceNum(diceNum string) (diceNumInt int, diceTime int64, key string, err error) {
	//1.按照标识符切分 前面是骰子点数，中间是时间戳，后面是标识字段
	diceNumList := strings.Split(diceNum, "fxm")
	if len(diceNumList) != 3 {
		fmt.Println("err")
		return
	}
	diceNumInt, err = strconv.Atoi(diceNumList[0])
	if err != nil {
		fmt.Println("err")
		return
	}
	diceTimeInt, err := strconv.Atoi(diceNumList[1])
	if err != nil {
		fmt.Println("err")
		return
	}
	diceTime = int64(diceTimeInt)
	key = diceNumList[2]
	//2.按照长度切分 前面4位是骰子点数，中间16位是时间戳，后面是标识字段
	//...
	return
}

func CheckDiceNum(diceNum string) {

}
