package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pagesize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pagesize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTageByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Update(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}
