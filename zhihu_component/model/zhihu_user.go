package model

import "time"

// ZhihuUser 话题评论表
type ZhihuUser struct {
	ID        int64     `gorm:"primaryKey;column:id"` // 主键ID
	Username  string    `gorm:"column:username"`      // 用户名称
	Phone     string    `gorm:"column:phone"`         // 电话号码
	Password  string    `gorm:"column:password"`      // 密码,加密形式
	Delete    *int8     `gorm:"column:delete"`        // 是否删除, 默认0-未删除 1-删除
	CreatedAt time.Time `gorm:"column:created_at"`    // 记录创建时间
	UpdatedAt time.Time `gorm:"column:updated_at"`    // 更新时间
}

func (z ZhihuUser) TableName() string {
	return "zhihu_user"
}
