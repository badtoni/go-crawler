package dbmodels

import (
	"gorm.io/gorm"
)

// type User struct {
// 	ID           uint
// 	Name         string
// 	Email        *string
// 	Age          uint8
// 	Birthday     *time.Time
// 	MemberNumber sql.NullString
// 	ActivedAt    sql.NullTime
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// }

type User struct {
	gorm.Model
	Name    string
	Address string
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  []Author `gorm:"embedded"`
	Upvotes int32
}

// type User2 struct {
// 	Name string `gorm:"<-:create"`          // allow read and create
// 	Name string `gorm:"<-:update"`          // allow read and update
// 	Name string `gorm:"<-"`                 // allow read and write (create and update)
// 	Name string `gorm:"<-:false"`           // allow read, disable write permission
// 	Name string `gorm:"->"`                 // readonly (disable write permission unless it configured )
// 	Name string `gorm:"->;<-:create"`       // allow read and create
// 	Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
// 	Name string `gorm:"-"`                  // ignore this field when write and read
// }
