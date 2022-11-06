package dices

import (
	"gorm.io/gorm"
)

type DiceTable struct {
	Id         int   `json:"Dice_id,omitempty"     gorm:"column:id;primary_key"`
	GameId     int   `json:"game_id"               gorm:"column:game_id"`
	Status     int   `json:"status"                gorm:"column:status"`
	UserId     int   `json:"user_id"               gorm:"column:win_user_id"`
	DiceNum    int   `json:"dice_num"              gorm:"column:dice_num"`
	CreateTime int64 `json:"create_time"           gorm:"column:create_time"`
}

func (DiceTable) TableName() string {
	return "Dice_t"
}

func (u *DiceTable) Insert(tx *gorm.DB) error {
	err := tx.Table(u.TableName()).Create(u).Error
	if err == nil {
		err = tx.Table(u.TableName()).Last(u).Error
	}
	return err
}

func (u *DiceTable) QueryByFilter(tx *gorm.DB, data map[string]interface{}) (Dice DiceTable, err error) {
	err = tx.Table(u.TableName()).Where(data).First(&Dice).Error
	return
}

func (u *DiceTable) QueryListByFilter(tx *gorm.DB, data map[string]interface{}) (Dice []DiceTable, err error) {
	err = tx.Table(u.TableName()).Where(data).Find(&Dice).Error
	return
}
