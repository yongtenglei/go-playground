package logic

import (
	"my_bubble/dao/mysql"
	"my_bubble/models"

	"go.uber.org/zap"
)

func TodoList() (todoList []*models.Todo, err error) {
	todoList, err = mysql.TodoList()
	if err != nil {
		zap.L().Error("controllers.TodoList err", zap.String("err", err.Error()))
		return nil, err
	}
	return
}

func GetATodo(id int64) (todo *models.Todo, err error) {
	todo = new(models.Todo)
	if todo, err = mysql.GetATodo(id); err != nil {
		zap.L().Error("controllers.GetATodo err", zap.String("err", err.Error()))
		return nil, err
	}
	return
}
func CreateTodo(todo *models.Todo) (err error) {
	if err = mysql.CreateTodo(todo); err != nil {
		zap.L().Error("controllers.CreateTodo err", zap.String("err", err.Error()))
		return
	}
	return
}

func UpdateTodo(todo *models.Todo) (err error) {
	if err = mysql.UpdateTodo(todo); err != nil {
		zap.L().Error("controllers.UpdateTodo err", zap.String("err", err.Error()))
		return
	}
	return
}

func DeleteTodo(todo *models.Todo) (err error) {
	if err = mysql.DeleteTodo(todo); err != nil {
		zap.L().Error("controllers.DeleteTodo err", zap.String("err", err.Error()))
		return
	}
	return

}
