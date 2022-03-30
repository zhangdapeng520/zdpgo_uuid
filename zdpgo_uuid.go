package zdpgo_uuid

import (
	"github.com/zhangdapeng520/zdpgo_uuid/uuid"
)

// UUID 返回UUID字符串
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}
