package dice

import (
	"dice/common/global"
	"dice/common/ico"
	"dice/model/dices"
	"dice/model/games"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Replay struct {
	GameId int `json:"game_id"  binding:"required"`
}

type User struct {
	BlueUserId int `json:"blue_user_id"`
	RedUserId  int `json:"red_user_id"`
	GameStatus int `json:"game_status"`
	WinUserId  int `json:"win_user_id"`
}

type Game struct {
	DiceNum int `json:"dice_num"`
	UserId  int `json:"user_id"`
}

type ReplayRes struct {
	User User   `json:"user"`
	Game []Game `json:"game"`
}

func (This Replay) DoHandle(c *gin.Context) *ico.Result {
	if err := c.ShouldBindJSON(&This); err != nil {
		return ico.Err(2099, "参数异常")
	}
	fmt.Printf("获取本场游戏记录 game_id: %d", This.GameId)
	// 1.校验
	err := This.check()
	if err != nil {
		return ico.Err(2001, "校验失败")
	}
	// 2.获取本场游戏信息
	game, err := This.getGame()
	if err != nil {
		return ico.Err(2004, "获取失败")
	}
	// 3.获取本场游戏全部记录
	gameList, err := This.getGameList()
	if err != nil {
		return ico.Err(2005, "获取失败")
	}
	Res := ReplayRes{User: game, Game: gameList}
	return ico.Succ(Res)
}

func (This Replay) check() (err error) {
	// 1.校验本局游戏是否存在
	gameT := &games.GameTable{}
	_, err = gameT.QueryByFilter(global.DBMysql, map[string]interface{}{"game_id": This.GameId})
	if err != nil {
		return
	}
	return
}

func (This Replay) getGame() (user User, err error) {
	gameT := &games.GameTable{}
	game, err := gameT.QueryByFilter(global.DBMysql, map[string]interface{}{"game_id": This.GameId})
	if err != nil {
		return
	}
	user.GameStatus = game.Status
	user.RedUserId = game.RedUserId
	user.BlueUserId = game.BlueUserId
	user.WinUserId = game.WinUserId
	return
}

func (This Replay) getGameList() (gameList []Game, err error) {
	diceT := &dices.DiceTable{}
	diceList, err := diceT.QueryListByFilter(global.DBMysql, map[string]interface{}{"game_id": This.GameId})
	if err != nil {
		return
	}
	for _, dice := range diceList {
		gameList = append(gameList, Game{
			UserId:  dice.UserId,
			DiceNum: dice.DiceNum,
		})
	}
	return
}
