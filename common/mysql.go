package common

import (
"fmt"
"net/url"
"time"

"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
)

type Orm struct {
	dbname string
}

func newOrm(aliasName string) *Orm {
	return &Orm{
		dbname: aliasName,
	}
}

func (self *Orm) ConnectMySQL(username string, password string, dbname string, mysql_ip string, param ...int) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	if len(mysql_ip) > 0 {
		connect := fmt.Sprintf(username+":"+password+"@tcp("+mysql_ip+")/"+dbname+"?charset=utf8&loc=%s", url.QueryEscape("Asia/Shanghai"))
		orm.RegisterDataBase(self.dbname, "mysql", connect, param...) // 最后2个参数: 最大空闲连接, 最大数据库连接
	} else {
		connect := fmt.Sprintf(username+":"+password+"@/"+dbname+"?charset=utf8&loc=%s", url.QueryEscape("Asia/Shanghai"))
		orm.RegisterDataBase(self.dbname, "mysql", connect, param...) // 最后2个参数: 最大空闲连接, 最大数据库连接
	}
}

func (self *Orm) NewOrm() orm.Ormer {
	o := orm.NewOrm()
	o.Using(self.dbname)
	return o
}

func (self *Orm) BuildDB(drop bool, log bool) {
	orm.RunSyncdb(self.dbname, drop, log)
}

var (
	sql *Orm
)

func InitOrm(aliasName string, print_log bool) {
	if sql != nil {
		return
	}
	sql = newOrm(aliasName)
	orm.Debug = print_log
}

func NewOrm() orm.Ormer {
	return sql.NewOrm()
}

func ConnectMySQL(username string, password string, dbname string, mysql_ip string, param ...int) {
	sql.ConnectMySQL(username, password, dbname, mysql_ip, param...)
}

func RegisterModel(model ...interface{}) {
	orm.RegisterModel(model...)
}

func Syncdb() {
	sql.BuildDB(false, true)
}

type DBHelper struct {
	Orm orm.Ormer `json:"-" orm:"-"`
}

func (self *DBHelper) NewOrm() orm.Ormer {
	if self.Orm == nil {
		self.Orm = NewOrm()
	}
	return self.Orm
}

