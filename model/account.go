package model

import "FlickServer/common"

type Account struct {
	Id       int64  `json:"id"`
	Username string `json:"name" orm:"size(80)"`
	Password string `json:"-" orm:"size(64)"`
	UserId   string `json:"userID" orm:size(200)"`
	//CommentData  *CommentData  `orm:"null;rel(fk)"` // fk表示外键
	//ScoreData  *ScoreData `orm:"null;rel(fk)"`
	UpdateTime int64 `json:"updateTime"`
	UserLevel  int16 `json:"userLevel"`
}

func (self *Account) QueryWithId(id int64) (*Account, error) {
	db := common.NewOrm()
	r := &Account{}
	// RelatedSel 即为关联查询所有相关表
	if err := db.QueryTable("Account").Filter("id", id).RelatedSel().One(r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (self *Account) QueryAllUsers() ([]*Account, error) {
	db := common.NewOrm()
	r := make([]*Account, 0, 1000)
	if _, err := db.QueryTable("Account").Limit(1000).All(&r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (self *Account) QueryWithUsername(username string) (*Account, error) {
	db := common.NewOrm()
	r := &Account{}
	if err := db.QueryTable("Account").Filter("username", username).One(r); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (self *Account) Add(username string, password string) (*Account, error) {
	db := common.NewOrm()
	r := &Account{
		Username: username,
		Password: password,
	}
	if id, err := db.Insert(r); err != nil {
		return nil, err
	} else {
		r.Id = id
		return r, nil
	}
}

//func (self *Account) Add(username string, password string) (*Account, error) {
//	db := common.NewOrm()
//	r := &Account{
//		Username: username,
//		Password: password,
//		//Bank: &Bank{},
//	}
//	// 开启事务
//	if err := db.Begin(); err != nil {
//		return nil, err
//	}
//	if id, err := db.Insert(r); err != nil {
//		db.Rollback() // 回滚
//		return nil, err
//	} else {
//
//		if bank_id, err := db.Insert(&Bank{}); err != nil {
//			db.Rollback() // 回滚
//			return nil, err
//		} else {
//			r.Bank.Id = bank_id
//			if _, err := db.Update(r, "bank"); err != nil {
//				db.Rollback() // 回滚
//				return nil, err
//			}
//		}
//
//		// 事务提交
//		if err := db.Commit(); err != nil {
//			return nil, err
//		}
//
//		r.Id = id
//		return r, nil
//	}
//}

func (self *Account) HasUsername(username string) (bool, error) {
	db := common.NewOrm()
	if num, err := db.QueryTable("account").Filter("username", username).Count(); err != nil {
		return false, err
	} else {
		return num >= 1, nil
	}
}

//type Bank struct {
//	Id   int64 `json:"id"`
//	Cash int64 `json:"cash"`
//}
