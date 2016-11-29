package dao

import (
	"esvodsApi/forms"
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//FeedDao ...
type FeedDao struct{}

//Get ...
func (d FeedDao) Get(id uint) (feed models.Feed, err error) {
	err = GetDB().First(&feed, id).Error
	return feed, err
}

//Find ...
func (d FeedDao) Find(s forms.FeedSearch) (feeds []models.Feed, err error) {
	err = db.Where(getQuery(structs.New(s))).Find(&feeds).Error
	return feeds, err
}

//Save ...
func (d FeedDao) Save(feed *models.Feed) (err error) {
	if GetDB().NewRecord(feed) {
		err = GetDB().Create(&feed).Error
	} else {
		err = GetDB().Save(&feed).Error
	}

	return err
}

//Delete ...
func (d FeedDao) Delete(id uint) error {
	feed, err := d.Get(id)
	if err == nil {
		err = db.Delete(&feed).Error
	}
	return err
}