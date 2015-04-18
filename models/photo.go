package models

import (
//	_ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
    "github.com/astaxie/beego"
	"log"
	"time"
    "github.com/Unknwon/com"
    "os"
    "path"
)

const (
// 设置数据库路径
    _DB_NAME = "data/fphotos.db"
// 设置数据库名称
    _SQLITE3_DRIVER = "sqlite3"
)

type SmPhoto struct {
	Id         int64
	Name       string
	Type       string
	Size       int32
	Ext        string `xorm:"size(100)"`
	Md5        string `xorm:"size(100)"`
	Sha1       string `xorm:"size(100)"`
	Savename   string `xorm:"size(200)"`
	Savepath   string `xorm:"size(100)"`
	Url        string
	UserKey    string
	CreateTime time.Time `xorm:"created"`
	From       string `xorm:"size(50)"`
}

var engine *xorm.Engine
//var mysqlEngine *xorm.Engine

func init() {

    if !com.IsExist(_DB_NAME) {
        os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
        os.Create(_DB_NAME)
    }

	var err error
//	mysqlEngine, _ = xorm.NewEngine("mysql", "root:root@/photo?charset=utf8")
    engine, err = xorm.NewEngine("sqlite3", _DB_NAME)
	if err != nil {
		log.Fatalf("Database Connect Error, [%v]", err)
		return
	}


    if beego.RunMode == "dev" {
        engine.ShowSQL = true   //则会在控制台打印出生成的SQL语句；
        engine.ShowDebug = true //则会在控制台打印调试信息；
        engine.ShowErr = true   //则会在控制台打印错误信息；
        engine.ShowWarn = true  //则会在控制台打印警告信息；
    }

    syncErr := engine.Sync2(new(SmPhoto))

    if syncErr != nil {
        log.Fatalf("syncErr: %v", syncErr)

    }
}

// @Title GetPhotoCount
// @Description 获取数量
// @param err error 错误信息
// @param count int64 数量
func (p *SmPhoto) GetPhotoCount() (err error, count int64) {
    count, err = engine.Count(p)
    return err, count
}

// @Title GetAllPhotos
// @Description 获取分页图片
// @param pageOffset int 起始位置
// @param pageSize int 每页的数量
func (p *SmPhoto) GetAllPhotos(pageOffset int, pageSize int) (err error, s []SmPhoto) {

	if pageSize == 0 {
        pageSize = 20
	}

	err = engine.OrderBy("create_time desc").Limit(pageSize, pageOffset).Find(&s)
	if err == nil {
		return nil, s
	}
	return err, nil
}

func (p *SmPhoto) InsertOnPhoto() error {
    _, err := engine.InsertOne(p)
    return err
}