package database

import (
	"github.com/huage66/zhihu_go/zhihu_creator/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Database interface {
	Open(conn string) (*gorm.DB, error)
}

type Mysql struct {
}

func (m *Mysql) Open() (*gorm.DB, error) {
	newLogger := glog.New(log.New(os.Stdout, "/r/n", log.LstdFlags), glog.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      glog.Info,
	})
	return gorm.Open(mysql.Open(config.Setting.Database.Tidb), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: false, // 禁用cache
	})
}
