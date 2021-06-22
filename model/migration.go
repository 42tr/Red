package model

// 数据迁移
func migration() {
	DB.AutoMigrate(&Project{}, &Approval{}, &Apply{}, &Remark{}, &Dic{}, &Income{}, &Party{})
}
