package models

import (
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
)

const _DB_NAME = "data/beelog.db"
const _SQLITE3_DRIVE = "sqlite3"

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.Mkdir(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVE, orm.DRSqlite)
	// 注册数据库，必须要有一个名称为"default"的DB
	orm.RegisterDataBase("default", _SQLITE3_DRIVE, _DB_NAME, 10)
}
