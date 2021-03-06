package service

import (
	"back/common"
	"github.com/dchest/uniuri"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"
)

func GetSummonedById(id string) (*common.Summoned, error){
	var summoned common.Summoned
	err = common.DB.Where("id = ?", id).First(&summoned).Error
	expired, _ := checkExpired(&summoned)
	if expired {
		common.DB.Save(&summoned)
	}
	return &summoned, err
}

func GetAllSummoned() []common.Summoned {
	var summoneds []common.Summoned
	common.DB.Find(&summoneds)
	for _, summoned := range summoneds {
		expired, _ := checkExpired(&summoned)
		if expired {
			common.DB.Save(&summoned)
		}
	}
	return summoneds
}

func saveFile(file multipart.File, fileName string) error {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile("./img/"+fileName, content, 0644); err != nil {
		return err
	}
	return nil
}

func checkExpired(summoned *common.Summoned) (bool, error){
	if summoned.Status != "Waiting" {
		return false, nil
	}
	now := time.Now()
	ddl, err := time.Parse("2006-01-02", summoned.Ddl)
	if err != nil {
		return false, err
	}
	if now.After(ddl) {
		summoned.Status = "Expired"
		return true, nil
	}
	summoned.Status = "Waiting"
	return false, nil
}

func NewSummoned(summoned common.Summoned, file multipart.File, userId interface{}) error{
	summoned.UserID = userId.(uint)
	summoned.Status = "Waiting"
	if _, err = checkExpired(&summoned); err != nil {
		return err
	}
	fileName := uniuri.New()
	if err = saveFile(file, fileName); err != nil {
		return err
	}
	summoned.Img = fileName
	common.DB.Create(&summoned)
	return nil
}

func getSummonedOpUserId(op string, userId interface{}) []common.Summoned{
	var summoneds []common.Summoned
	common.DB.Where("user_id " + op + " ?", userId).Find(&summoneds)
	for _, summoned := range summoneds {
		expired, _ := checkExpired(&summoned)
		if expired {
			common.DB.Save(&summoned)
		}
	}
	return summoneds
}

func GetSummonedByUserId(userId interface{}) []common.Summoned{
	return getSummonedOpUserId("=", userId)
}

func GetSummonedExceptUserId(userId interface{}) []common.Summoned{
	return getSummonedOpUserId("<>", userId)
}

func UpdateSummoned(summoned common.Summoned, file multipart.File, keepImg bool) error{
	var summonedInMysql common.Summoned
	common.DB.First(&summonedInMysql, summoned.ID)
	summonedInMysql.Type = summoned.Type
	summonedInMysql.Name = summoned.Name
	summonedInMysql.Desc = summoned.Desc
	summonedInMysql.People = summoned.People
	summonedInMysql.Ddl = summoned.Ddl
	if _, err = checkExpired(&summoned); err != nil {
		return err
	}
	if !keepImg {
		fileName := uniuri.New()
		if err = saveFile(file, fileName); err != nil {
			return err
		}
		summonedInMysql.Img = fileName
	}
	common.DB.Save(&summonedInMysql)
	return nil
}

func DeleteSummoned(summonedId int) error{
	var summonedInMysql common.Summoned
	common.DB.First(&summonedInMysql, summonedId)
	if err = os.Remove("./img/" + summonedInMysql.Img); err != nil {
		return err
	}
	summonedInMysql.Status = "Cancelled"
	common.DB.Save(&summonedInMysql)
	return nil
}