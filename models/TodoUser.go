package models

import "time"

type User struct {
	UserID         int       `db:"user_id"`
	Name           string    `db:"name"`
	Password       string    `db:"pass"` // Avoid using "pass" as it can be misleading. Consider using "Password" instead.
	DisplayPicture *string   `db:"display_picture"`
	CreatedOn      time.Time `db:"created_on"`
	FirebaseToken  *string   `db:"firebase_token"`
	EmailID        *string   `db:"email_id"`
	MobileNo       string    `db:"mobile_no"`
	IsActive       bool      `db:"is_active"`
	IsPremium      bool      `db:"is_premium"`
}
