package model

import (
	"Todo/repository/cache"
	"context"
	"gorm.io/gorm"
	"strconv"
)

type Task struct {
	gorm.Model
	Uid       uint
	Title     string `gorm:"index;not null"`
	Content   string `gorm:"type:longtext"`
	Status    int    `gorm:"default:0"`
	StartTime string
	EndTime   string `gorm:"default:0"`
}

func (Task *Task) View() int {
	countStr, _ := cache.RedisClient.Get(context.Background(), cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.Atoi(countStr)
	return count
}

func (Task *Task) AddView() {
	ctx := context.Background()
	cache.RedisClient.Incr(ctx, cache.TaskViewKey(Task.ID))               // 增加视频点击数
	cache.RedisClient.ZIncrBy(ctx, "rank", 1, strconv.Itoa(int(Task.ID))) // 增加排行点击数
}
