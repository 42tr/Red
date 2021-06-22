package serializer

import "red/model"

// User 用户序列化器
type Project struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CreateBy    uint    `json:"creator"`
	CreateAt    string  `json:"createAt"`
	Type        string  `json:"type"`
	Price       float32 `json:"price"`
	PartyId     uint    `json:"partyId"`
	Party       string  `json:"party"`
	Area        string  `json:"area"`
	Price1      float32 `json:"price1"`
	Price2      float32 `json:"price2"`
	Price3      float32 `json:"price3"`
}

// BuildUser 序列化用户
func BuildProject(project model.Project, party string, priceA, priceB, priceC float32) Project {
	return Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreateBy:    project.CreateBy,
		CreateAt:    project.CreatedAt.Format("2006-01-02 15:04:05"),
		Type:        project.Type,
		Price:       project.Price,
		PartyId:     project.PartyId,
		Party:       party,
		Area:        project.Area,
		Price1:      priceA,
		Price2:      priceB,
		Price3:      priceC,
	}
}

// BuildUserResponse 序列化用户响应
// func BuildProjectResponse(project model.Project, partyMap map[uint]string) Response {
// 	return Response{
// 		Data: BuildProject(project, partyMap[project.PartyId]),
// 	}
// }

func BuildProjects(projects []model.Project, partyMap map[uint]string, incomes []model.Income) []Project {
	data := make([]Project, 0)
	for _, project := range projects {
		var priceA, priceB, priceC float32
		for _, income := range incomes {
			if income.ProjectId == project.ID {
				if income.Type == "代理费" || income.Type == "审计费" {
					priceA += income.Amount
				} else if income.Type == "清单编制费" {
					priceB += income.Amount
				} else {
					priceC += income.Amount
				}
			}
		}
		data = append(data, BuildProject(project, partyMap[project.PartyId], priceA, priceB, priceC))
	}
	return data
}
