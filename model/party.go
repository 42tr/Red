package model

type Party struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetPartyMap() map[uint]string {
	var list []Party
	DB.Find(&list)

	m := make(map[uint]string)
	for _, item := range list {
		m[item.ID] = item.Name
	}
	return m
}
