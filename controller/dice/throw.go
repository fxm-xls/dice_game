package dice

import (
	"dice/common/global"
	"dice/common/ico"
	"dice/controller/public"
	"dice/model/dices"
	"dice/model/games"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Throw struct {
	UserId  int    `json:"user_id"  binding:"required"`
	GameId  int    `json:"game_id"  binding:"required"`
	DiceNum string `json:"dice_num" binding:"required"`
}

type ThrowRes struct {
	DiceNum int `json:"dice_num"`
}

func (This Throw) DoHandle(c *gin.Context) *ico.Result {
	if err := c.ShouldBindJSON(&This); err != nil {
		return ico.Err(2099, "参数异常")
	}
	fmt.Printf("用户投掷骰子 user_id: %d game_id: %d", This.UserId, This.GameId)
	// 1.校验
	diceNumInt, err := This.check()
	if err != nil {
		return ico.Err(2001, "校验失败")
	}
	// 2.存入数据库
	diceT := &dices.DiceTable{
		GameId:     This.GameId,
		UserId:     This.UserId,
		DiceNum:    diceNumInt,
		CreateTime: time.Now().Unix(),
	}
	if err = diceT.Insert(global.DBMysql); err != nil {
		return ico.Err(2002, "存入失败")
	}
	// 3.获取user当前骰子点数
	userDiceNum, err := This.getUserDiceNum(diceNumInt)
	if err != nil {
		return ico.Err(2003, "获取失败")
	}
	Res := ThrowRes{DiceNum: userDiceNum}
	return ico.Succ(Res)
}

func (This Throw) check() (diceNumInt int, err error) {
	// 0.解析骰子点数
	diceNumInt, diceTime, key, err := public.SplitDiceNum(This.DiceNum)
	if err != nil {
		fmt.Println()
		return
	}
	// 1.验证骰子点数
	// 1.1 时效性校验
	if err = CheckDiceTime(diceTime); err != nil {
		fmt.Println()
		return
	}
	// 1.2 点数校验
	if err = CheckDiceNum(diceNumInt); err != nil {
		fmt.Println()
		return
	}
	// 1.3 加入验证字段
	if err = CheckDiceKey(key); err != nil {
		fmt.Println()
		return
	}
	// 2.校验游戏是否结束
	gameT := &games.GameTable{}
	game, err := gameT.QueryByFilter(global.DBMysql, map[string]interface{}{"game_id": This.GameId})
	if err != nil {
		fmt.Println()
		return
	}
	if game.Status != 1 {
		// 1：游戏中 2：游戏结束 3：平局
		err = errors.New("游戏已结束")
		return
	}
	// 3.校验用户是否在该游戏中
	if This.UserId != game.BlueUserId && This.UserId != game.RedUserId {
		err = errors.New("参数异常")
		return
	}
	return
}

func (This Throw) getUserDiceNum(diceNumInt int) (userDiceNum int, err error) {
	gameT := &games.GameTable{}
	game, err := gameT.QueryByFilter(global.DBMysql, map[string]interface{}{"game_id": This.GameId})
	if err != nil {
		return
	}
	if This.UserId == game.BlueUserId {
		userDiceNum = diceNumInt + game.BlueDiceNum
	} else {
		userDiceNum = diceNumInt + game.RedDiceNum
	}
	return
}

func CheckDiceTime(diceTime int64) (err error) {
	nowTime := time.Now().Unix()
	if nowTime-5 > diceTime {
		err = errors.New("err")
		return
	}
	return
}

func CheckDiceNum(diceNumInt int) (err error) {
	if diceNumInt > 6 || diceNumInt < 1 {
		err = errors.New("err")
		return
	}
	return
}

func CheckDiceKey(key string) (err error) {
	// 可与时间戳base64解码之后 与设置的key比较
	if key != "fxm" {
		err = errors.New("err")
		return
	}
	return
}
