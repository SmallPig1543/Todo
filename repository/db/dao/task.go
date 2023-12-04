package dao

import (
	"Todo/repository/db/model"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(task *model.Task) error {
	return dao.Model(&model.Task{}).Create(&task).Error
}

func (dao *TaskDao) Update(task *model.Task) error {
	err := dao.DB.Save(task).Error
	return err
}

func (dao *TaskDao) FindTaskById(id uint, uid uint) (task *model.Task, err error) {
	err = dao.DB.Model(&model.Task{}).Where("id=? AND uid=?", id, uid).First(&task).Error
	fmt.Println(err)
	return task, err
}

func (dao *TaskDao) SearchTasksByInfo(uid uint, info string, start int) (tasks []*model.Task, count int64, err error) {
	info = "%" + info + "%"
	err = dao.DB.Model(&model.Task{}).Where("uid= ? AND (content like ? OR title like ?)", uid, info, info).Count(&count).
		Limit(10).Offset((start - 1) * 10).Find(&tasks).Error
	return
}

func (dao *TaskDao) SearchTasksByStatus(uid uint, status int, start int) (tasks []*model.Task, count int64, err error) {
	err = dao.DB.Model(&model.Task{}).Where("uid=? AND status=?", uid, status).Count(&count).
		Limit(10).Offset((start - 1) * 10).Find(&tasks).Error
	return tasks, count, err
}

func (dao *TaskDao) ListTasks(uid uint, limit int, start int) (tasks []*model.Task, count int64, err error) {
	err = dao.DB.Model(&model.Task{}).Where("uid=?", uid).Count(&count).
		Limit(limit).Offset((start - 1) * limit).
		Find(&tasks).Error
	return
}

func (dao *TaskDao) Delete(task *model.Task) error {
	err := dao.DB.Delete(&model.Task{}, task).Error
	return err
}

func (dao *TaskDao) DeleteWithStatus(uid uint, status int) (err error) {
	if status == 0 || status == 1 {
		err = dao.DB.Model(&model.Task{}).Where("uid = ? AND status = ?", uid, status).Delete(&model.Task{}).Error
	} else {
		err = dao.DB.Model(&model.Task{}).Where("uid = ?", uid).Delete(&model.Task{}).Error
	}
	return
}
