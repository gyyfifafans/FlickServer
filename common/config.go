package common

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/seefan/to"
)

type Config struct {
	data map[string]string
}

func NewConfig() *Config {
	return &Config{
		data: make(map[string]string, 20),
	}
}

func (self *Config) Load(fileName string) error {
	if conf, err := ini.Load(fileName); err != nil {
		return err
	} else {
		sections := conf.Sections()
		for _, v := range sections {
			keys := v.Keys()
			for _, vv := range keys {
				self.data[vv.Name()] = vv.String()
			}
		}
	}
	return nil
}

func (self *Config) Get(key string) string {
	return self.data[key]
}

func (self *Config) Int64(key string) int64 {
	return to.Int64(self.Get(key))
}

func (self *Config) Int(key string) int {
	return int(self.Int64(key))
}

func (self *Config) Has(key string) bool {
	if _, exist := self.data[key]; exist {
		return true
	}
	return false
}

func (self *Config) PrintAll() {
	fmt.Printf("[Config] %+v\n", self.data)
}
