package model

import "FlickServer/common"

type Account struct {
	Id       int64  `json:"id"`
	Username string `json:"-" orm:"size(80)"`
	Password string `json:"-" orm:"size(64)"`
	Nickname string `json:"nickname" orm:"size(80)"`
}

func (self *Account) QueryWithId(id int64) (*Account, error) {
	db := common.NewOrm()
	r := &Account{}
	if err := db.QueryTable("Account").Filter("id", id).One(r); err != nil {
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

func (self *Account) Add(username string, password string, nickname string) (*Account, error) {
	db := common.NewOrm()
	r := &Account{
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	if id, err := db.Insert(r); err != nil {
		return nil, err
	} else {
		r.Id = id
		return r, nil
	}
}

func (self *Account) HasUsername(username string) (bool, error) {
	db := common.NewOrm()
	if num, err := db.QueryTable("account").Filter("username", username).Count(); err != nil {
		return false, err
	} else {
		return num >= 1, nil
	}
}
