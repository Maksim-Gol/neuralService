package neuralservice

import (
	"github.com/google/uuid"
)

type ServiceCall struct {
	//? Импортнул библиотеку для uuid - мб оверкилл ли и сойдёт обычная строка?
	UserID    uuid.UUID `json:"user_id"`
	ModelID   uuid.UUID `json:"model_id"`
	RequestID uuid.UUID `json:"request_id"`
	Cost int `json:"cost"`
	Status string `json:"status"`

	//? Нужен ли какой-то особый тип для метаданных? Или string сойдёт?
	Metadata string `json:"metadata"`
}
