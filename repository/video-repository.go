package repository

import (
	"PR_gin_g/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	Detail(video entity.Video) entity.Video
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVidepRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

func (db *database) Detail(video entity.Video) entity.Video {
	var videoDetail entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videoDetail, video.ID)
	return videoDetail
}

func (db *database) CloseDB() {
	db.connection.Close()
}
