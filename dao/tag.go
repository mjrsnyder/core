package dao

import "esvodsCore/models"

//TagDao ...
type TagDao struct{}

//Get ...
func (d TagDao) Get(id uint) (tag models.Tag, err error) {
	err = GetDB().First(&tag, id).Error
	return tag, err
}

//Find ...
func (d TagDao) Find(query map[string]interface{}) (tags []models.Tag, err error) {
	err = db.Where(query).Find(&tags).Error
	return tags, err
}

//Save ...
func (d TagDao) Save(tag *models.Tag) (err error) {
	if GetDB().NewRecord(tag) {
		err = GetDB().Create(&tag).Error
	} else {
		err = GetDB().Model(&tag).Updates(getUpdates(tag)).Error
	}

	return err
}

//Delete ...
func (d TagDao) Delete(id uint) (tag models.Tag, err error) {
	tag, err = d.Get(id)
	if err == nil {
		err = db.Delete(&tag).Error
	}
	return tag, err
}
