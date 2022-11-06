package users

type UserTable struct {
	UserId     int    `json:"user_id,omitempty"     gorm:"column:id;primary_key"`
	UserName   string `json:"user_name"             gorm:"column:user_name"`
	Password   string `json:"password"              gorm:"column:password"`
	CreateTime int64  `json:"create_time"           gorm:"column:create_time"`
}

func (UserTable) TableName() string {
	return "user_t"
}
