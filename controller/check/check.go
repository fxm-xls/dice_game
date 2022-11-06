package check

import (
	"dice/common/ico"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Check struct {
	DiceNum string `json:"dice_num" binding:"required"`
}

type CheckRes struct {
	DiceNum int `json:"dice_num"`
}

func (This Check) DoHandle(c *gin.Context) *ico.Result {
	if err := c.ShouldBindJSON(&This); err != nil {
		return ico.Err(2099, "")
	}
	fmt.Printf("校验骰子点数 dice_num: %s", This.DiceNum)

	Res := CheckRes{DiceNum: 1}
	return ico.Succ(Res)
}
