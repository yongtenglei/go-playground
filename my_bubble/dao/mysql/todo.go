package mysql

import (
	"my_bubble/models"

	"go.uber.org/zap"
)

func TodoList() (todoList []*models.Todo, err error) {
	if err = db.Find(&todoList).Error; err != nil {
		zap.L().Error("Get TodoList failed")
		return nil, err
	}
	return
}

func CreateTodo(todo *models.Todo) (err error) {
	if err = db.Create(todo).Error; err != nil {
		zap.L().Error("Create Todo err", zap.String("err", err.Error()))
		return
	}
	return
}

func GetATodo(id int64) (todo *models.Todo, err error) {

	todo = new(models.Todo)
	if err = db.Select("id").Where("id = ?", id).First(&todo).Error; err != nil {
		zap.L().Error("Get todo with id failed")
		return nil, err
	}
	return
}

func UpdateTodo(todo *models.Todo) (err error) {
	if err = db.Save(todo).Error; err != nil {
		zap.L().Error("Update todo with id failed")
		return
	}
	return
}

func DeleteTodo(todo *models.Todo) (err error) {
	if err = db.Delete(todo).Error; err != nil {
		zap.L().Error("delete todo failed")
		return

	}

	return
}
