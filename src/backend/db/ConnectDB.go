package db

import (
	"backend/models"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	sqlDb, err := sql.Open("mysql", "root:Fchrgrib2310*@tcp(172.19.160.1:3306)/chatGPT")

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Chat{}, &models.QuestAns{}); err != nil {
		return nil, err
	}

	return db, nil
}
