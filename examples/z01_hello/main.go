package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_uuid/uuid"
)

func main() {
	uid := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", uid)

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)

	// 转换为字符串
	fmt.Println("\n转换为字符串：", uid.String())
}
