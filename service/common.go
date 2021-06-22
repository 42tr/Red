package service

import "red/model"

func GetDic() []model.Dic {
	var dics []model.Dic
	model.DB.Find(&dics)
	return dics
}
