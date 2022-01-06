package model

import (
	"time"
)

// ZhihuArticle 文章表
type ZhihuArticle struct {
	ID        int64     `gorm:"primaryKey;column:id"` // 主键ID
	Title     string    `gorm:"column:title"`         // 文章标题
	Content   string    `gorm:"column:content"`       // 文章内容
	UserID    int64     `gorm:"column:user_id"`       // 用户id
	Delete    *int8     `gorm:"column:delete"`        // 是否删除, 默认0-未删除 1-删除
	CreatedAt time.Time `gorm:"column:created_at"`    // 记录创建时间
	UpdatedAt time.Time `gorm:"column:updated_at"`    // 更新时间
}

func (z ZhihuArticle) TableName() string {
	return "zhihu_article"
}

