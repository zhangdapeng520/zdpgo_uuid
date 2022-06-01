package zdpgo_uuid

import (
	"github.com/zhangdapeng520/zdpgo_uuid/uuid"
)

type Uuid struct {
	Config *Config
}

func New() *Uuid {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Uuid {
	u := &Uuid{}

	// 配置
	u.Config = config

	// 返回
	return u
}

// String 获取UUID字符串
func (u *Uuid) String() string {
	uid := uuid.NewV4()
	return uid.String()
}

// UUID 返回UUID字符串
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}
