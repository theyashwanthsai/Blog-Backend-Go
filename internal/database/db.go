// initialize database
package database

import (
	"crypto/internal/nistec"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB gorm.DB

func connectDB(){
	dsn := "username:pass@tcp(your_mysql_host:your_mysql_port)/new?parseTime=true"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if(err != nil){
		panic("Failed to connect to db")
	}

	DB = db

}
