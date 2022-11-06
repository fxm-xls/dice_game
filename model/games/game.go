package games

import "gorm.io/gorm"

type GameTable struct {
	Id          int   `json:"Game_id,omitempty"     gorm:"column:id;primary_key"`
	GameId      int   `json:"game_id"               gorm:"column:game_id"`
	Status      int   `json:"status"                gorm:"column:status"`
	BlueUserId  int   `json:"blue_user_id"          gorm:"column:blue_user_id"`
	BlueDiceNum int   `json:"blue_dice_num"         gorm:"column:blue_dice_num"`
	RedUserId   int   `json:"red_user_id"           gorm:"column:red_user_id"`
	RedDiceNum  int   `json:"red_dice_num"          gorm:"column:red_dice_num"`
	WinUserId   int   `json:"win_user_id"           gorm:"column:win_user_id"`
	CreateTime  int64 `json:"create_time"           gorm:"column:create_time"`
}

func (GameTable) TableName() string {
	return "Game_t"
}

func (u *GameTable) Insert(tx *gorm.DB) error {
	err := tx.Table(u.TableName()).Create(u).Error
	if err == nil {
		err = tx.Table(u.TableName()).Last(u).Error
	}
	return err
}

func (u *GameTable) QueryByFilter(tx *gorm.DB, data map[string]interface{}) (Game GameTable, err error) {
	err = tx.Table(u.TableName()).Where(data).First(&Game).Error
	return
}

func (u *GameTable) QueryListByFilter(tx *gorm.DB, data map[string]interface{}) (Game []GameTable, err error) {
	err = tx.Table(u.TableName()).Where(data).Find(&Game).Error
	return
}
