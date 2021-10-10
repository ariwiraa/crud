package controller

import (
	"crud-echo/entity"
	"crud-echo/internal/config"
	"net/http"
)

func CreateManager(data *entity.Manager) (entity.Manager, entity.Error) {
	db := config.CreateConnection()
	res := db.Create(data)

	if res.Error != nil {
		return entity.Manager{}, entity.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}

	return *data, entity.Error{}
}
