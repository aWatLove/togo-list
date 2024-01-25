package repository

import (
	"fmt"
	"strings"

	todo "github.com/aWatLove/togo-list/pkg/model"
	"gorm.io/gorm"
)

type TodoItemPostgres struct {
	db *gorm.DB
}

func NewTodoItemPostgres(db *gorm.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, input todo.TodoItem) (int, error) {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&input).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	id := input.Id

	if err := tx.Create(&todo.ListsItem{ListId: listId, ItemId: id}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit().Error
}

func (r *TodoItemPostgres) GetAll(listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	result := r.db.Joins("JOIN lists_items on lists_items.item_id = todo_items.id").
		Where("lists_items.list_id=?", listId).Find(&items)

	return items, result.Error
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	result := r.db.Joins("INNER JOIN lists_items on lists_items.item_id = todo_items.id").
		Joins("INNER JOIN users_lists on users_lists.list_id = lists_items.list_id").
		Where("todo_items.id=? AND users_lists.user_id=?", itemId, userId).First(&item)

	return item, result.Error
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	result := r.db.Exec(`DELETE FROM todo_items ti USING lists_items li, users_lists ul 
					WHERE ti.id = li.item_id AND li.list_id=ul.list_id AND ul.user_id=$1 AND ti.id=$2`,
		userId, itemId)
	return result.Error
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *&input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *&input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *&input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE todo_items ti SET %s FROM lists_items li, users_lists ul
							WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		setQuery, argId, argId+1)
	args = append(args, userId, itemId)

	err := r.db.Exec(query, args...).Error

	return err
}
