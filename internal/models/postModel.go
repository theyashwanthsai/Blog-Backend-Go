// data models
package postModel

import(
	"gorm.io/gorm"
)
type Post struct{
	gorm.Model
	Title string
	Content string
}