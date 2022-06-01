package zdpgo_uuid

import (
	"fmt"
	"testing"
)

func getUuid() *Uuid {
	return NewWithConfig(&Config{})
}

func TestUUID(t *testing.T) {
	fmt.Println(UUID())
}

func TestUuid_String(t *testing.T) {
	u := getUuid()
	s1 := u.String()

	for i := 0; i < 10000; i++ {
		s2 := u.String()
		if s1 == s2 {
			panic("UUID重复了")
		}
	}
}
