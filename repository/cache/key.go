package cache

import (
	"fmt"
	"strconv"
)

// TaskViewKey 点击数的key
func TaskViewKey(id uint) string {
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}
