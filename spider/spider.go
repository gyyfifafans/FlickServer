package spider

import (
	"FlickServer/common"
	"FlickServer/model"
	"encoding/json"
	"github.com/seefan/to"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"strings"
)

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func u16ToCh(u16 string) string {
	if r, err := UnescapeUnicode([]byte(u16)); err != nil {
		panic(err)
	} else {
		return string(r)
	}
}

func get(url string) (string, error) {
	if _, resp, errs := gorequest.New().Get(url).End(); len(errs) != 0 {
		return "", errs[0]
	} else {
		return resp, nil
	}
}

func Capture() {
	if err := model.Config.Load("config/app.ini"); err != nil {
		panic("配置文件加载失败.")
	}

	{
		db := common.NewOrm()
		db.Raw("delete from music_data;").Exec() // 移除原music_data数据
		db.Raw("truncate table music_data;").Exec() // 设置表id从1开始
		url := "http://timetag.main.jp/nicoflick/nicoflick.php?req=music&time=0"
		captureMusicData(url)
	}
}

func captureMusicData(url string) {
	db := common.NewOrm()
	if r, err := get(url); err != nil {
		fmt.Printf("错误: %s\n", err.Error())
	} else {
		// json数据id、updatetime、createtime都是字符串
		// 表结构里是int64
		// 建立一个兼容的结构体，插入时再使用原结构体
		type Temp struct {
			IdStr string `json:"id"`
			UpdateTimeStr string `json:"updateTime"`
			CreateTimeStr string `json:"createTime"`
			model.MusicData
		}
		list := make([]*Temp, 0, 20)
		if err := json.Unmarshal([]byte(r), &list); err != nil {
			fmt.Printf("json解析失败: %s\n", err.Error())
		} else {
			// 写入mysql，使用事务
			if err := db.Begin(); err != nil {
				fmt.Printf("%s\n", err.Error())
				return
			}

			for _, v := range list {
				// 转utf8
				v.Title = u16ToCh(v.Title)
				v.Artist = u16ToCh(v.Artist)
				v.Tags = u16ToCh(v.Tags)
				// 字符串转数字
				v.Id = to.Int64(v.IdStr)
				v.UpdateTime = to.Int64(v.UpdateTimeStr)
				v.CreateTime = to.Int64(v.CreateTimeStr)
				// 插入 music data
				if _, err := db.Insert(&v.MusicData); err != nil {
					db.Rollback()
					fmt.Printf("%s\n", err.Error())
					return
				}
			}

			if err := db.Commit(); err != nil {
				fmt.Printf("%s\n", err.Error())
				return
			}
		}
	}
}