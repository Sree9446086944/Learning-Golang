package models

import "gorm.io/gorm"

// https://gorm.io/docs/models.html

/*type User struct {
	gorm.Model    //gorm.Model will add 4 fields shown below automatically - ID,createdAt,updatedAt,DeletedAt
	Name string
  }
  // equals
  type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string
  }
*/

type Post struct {
	gorm.Model //adds 4 fields automatically
	Title      string
	Body       string
}
