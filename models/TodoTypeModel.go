package models

import "time"

type TodoTypeModel struct {
	TypeId    int       `db:"type_id"`
	TypeName  string    `db:"type_name"`
	CreateID  int       `db:"create_id"`
	CreatedOn time.Time `db:"created_on"`
	ColorID   int       `db:"color_id"`
}
