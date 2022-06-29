package config 

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



var (
	db * gorm.DB
)

func Connect(){
	// Below, root is my admin name
	// Thedudzrocks@03 is my password for mysql root
	// globalfair_project is the database in mysql where all the data is stored
	// Please change the above 3 parameters according to your system
	d, err := gorm.Open("mysql", "root:Thedudzrocks@03@/globalfair_project?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}