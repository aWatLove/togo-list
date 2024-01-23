package repository

import (
	"fmt"
	"strings"

	todo "github.com/aWatLove/togo-list/pkg/model"
	"gorm.io/gorm"
)

type TodoListPostgres struct {
	db *gorm.DB
}

func NewTodoListPostgres(db *gorm.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&list).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	id := list.Id

	if err := tx.Create(&todo.UsersList{UserId: userId, ListId: id}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit().Error
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	result := r.db.Joins("JOIN users_lists on users_lists.list_id = todo_lists.id").
		Where("users_lists.user_id=?", userId).Find(&lists)

	return lists, result.Error
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	result := r.db.Joins("Join users_lists on users_lists.list_id = todo_lists.id").
		Where("users_lists.user_id=? and users_lists.list_id=?", userId, listId).First(&list)

	return list, result.Error
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	result := r.db.Exec("DELETE FROM todo_lists tl USING users_lists ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		userId, listId)
	return result.Error
}

func (r *TodoListPostgres) Update(userId, listId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE todo_lists tl SET %s FROM users_lists ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
	setQuery, argId, argId+1)
	
	args = append(args, listId, userId)

	err := r.db.Exec(query, args...).Error

	return err
}
