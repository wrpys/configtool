package daos

import (
	"configtool/common"
	"configtool/models"
	"configtool/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/goinggo/mapstructure"
)

func GetDistinctProfile() []string {
	result := make([]string, 0)
	var list orm.ParamsList
	_, err := dao.QueryTable(new(models.ConfigProperties)).Distinct().ValuesFlat(&list, "Profile")
	if err != nil {
		fmt.Println(err)
	}
	for _, l := range list {
		result = append(result, utils.ToString(l))
	}
	return result
}

func GetDistinctApplication() []string {
	result := make([]string, 0)
	var list orm.ParamsList
	_, err := dao.QueryTable(new(models.ConfigProperties)).Distinct().ValuesFlat(&list, "Application")
	if err != nil {
		fmt.Println(err)
	}
	for _, l := range list {
		result = append(result, utils.ToString(l))
	}
	return result
}

func FindConfigProperties(pageInfo common.PageInfo, configProperties models.ConfigProperties) (int64, []models.ConfigProperties) {

	var total int64
	rows := make([]models.ConfigProperties, 0)

	cpQt := new(models.ConfigProperties)
	qs := dao.QueryTable(cpQt)

	if configProperties.Profile != "" && len(configProperties.Profile) > 0 {
		qs = qs.Filter("Profile", configProperties.Profile)
	}
	if configProperties.Application != "" && len(configProperties.Application) > 0 {
		qs = qs.Filter("Application", configProperties.Application)
	}
	if configProperties.ConfigKey != "" && len(configProperties.ConfigKey) > 0 {
		qs = qs.Filter("ConfigKey__icontains", configProperties.ConfigKey)
	}
	if configProperties.ConfigValue != "" && len(configProperties.ConfigValue) > 0 {
		qs = qs.Filter("ConfigValue__icontains", configProperties.ConfigValue)
	}

	if count, err := qs.Count(); err != nil {
		fmt.Println(err)
	} else if count == 0 {
		return count, rows
	} else {
		total = count
	}

	qs = qs.Limit(pageInfo.Limit, pageInfo.Offset)

	var maps []orm.Params
	num, err := qs.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			item := new(models.ConfigProperties)
			if err = mapstructure.Decode(m, item); err != nil {
				fmt.Println(err)
			}
			item.CreateDate = utils.ToDate(m["CreateDate"])
			item.ChangeDate = utils.ToDate(m["ChangeDate"])
			rows = append(rows, *item)
		}
	}
	return total, rows
}

func CreateConfigProperties(configProperties *models.ConfigProperties) {
	if _, err := dao.Insert(configProperties); err != nil {
		fmt.Println(err)
	}
}

func UpdateConfigProperties(configProperties *models.ConfigProperties) {
	if _, err := dao.Update(configProperties); err != nil {
		fmt.Println(err)
	}
}

func DeleteConfigProperties(configIds []int) {
	if len(configIds) == 0 {
		return
	}
	cpQt := new(models.ConfigProperties)
	qs := dao.QueryTable(cpQt)
	qs.Filter("ConfigId__in", configIds).Delete()
}

func BatchUpdateProfile(currentProfile string, deployProfile string) {
	cpQt := new(models.ConfigProperties)
	qs := dao.QueryTable(cpQt)
	qs.Filter("Profile", currentProfile).Update(orm.Params{
		"Profile": deployProfile,
	})
}
