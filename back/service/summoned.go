package service

import (
	"back/common"
	"io/ioutil"
	"mime/multipart"
)

func GetSummonedById(id string) (*common.Summoned, error){
	var summoned common.Summoned
	err = common.DB.Where("id = ?", id).First(&summoned).Error
	return &summoned, err
}

func GetAllSummoned() []common.Summoned {
	var summoneds []common.Summoned
	common.DB.Find(&summoneds)
	return summoneds
}

func saveFile(file multipart.File, header *multipart.FileHeader) error {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile("./img/"+header.Filename, content, 0644); err != nil {
		return err
	}
	return nil
}

func NewSummoned(summoned common.Summoned, file multipart.File, header *multipart.FileHeader, userId interface{}) error{
	summoned.UserID = int(userId.(uint))
	summoned.Status = "Waiting"
	if err = saveFile(file, header); err != nil {
		return err
	}
	summoned.Img = header.Filename
	common.DB.Create(&summoned)
	return nil
}

func getSummonedOpUserId(op string, userId interface{}) []common.Summoned{
	var summoneds []common.Summoned
	common.DB.Where("user_id " + op + " ?", userId).Find(&summoneds)
	return summoneds
}

func GetSummonedByUserId(userId interface{}) []common.Summoned{
	return getSummonedOpUserId("=", userId)
}

func GetSummonedExceptUserId(userId interface{}) []common.Summoned{
	return getSummonedOpUserId("<>", userId)
}

func UpdateSummoned(summoned common.Summoned, file multipart.File, header *multipart.FileHeader, keepImg bool) error{
	var summonedInMysql common.Summoned
	common.DB.First(&summonedInMysql, summoned.ID)
	summonedInMysql.Type = summoned.Type
	summonedInMysql.Name = summoned.Name
	summonedInMysql.Desc = summoned.Desc
	summonedInMysql.People = summoned.People
	summonedInMysql.Ddl = summoned.Ddl
	if !keepImg {
		if err = saveFile(file, header); err != nil {
			return err
		}
		summonedInMysql.Img = header.Filename
	}
	common.DB.Save(&summonedInMysql)
	return nil
}

func DeleteSummoned(summoned common.Summoned) {
	var summonedInMysql common.Summoned
	common.DB.First(&summonedInMysql, summoned.ID)
	common.DB.Delete(&summonedInMysql)
}