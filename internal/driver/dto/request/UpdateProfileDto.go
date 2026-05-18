package request

import (
	"github.com/google/uuid"
	"koala.com/internal/driver/entity"
)

type UpdateProfileDto struct {
	Id     *int64               `db:"id"`
	UserId *uuid.UUID           `db:"user_id"`
	Status *entity.DriverStatus `db:"status"`
}
