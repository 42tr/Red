package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"red/model"
	"red/serializer"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetApprovalMenu(c *gin.Context) serializer.Response {
	var menu []model.Approval
	model.DB.Find(&menu)

	return serializer.Success(menu)
}

type Apply struct {
	ID          uint   `json:"id"`
	ProjectName string `json:"projectName"`
	Name        string `json:"name"`
	ApplyUserId string `json:"applyUserId"`
	ApplyDate   string `json:"applyDate"`
	Status      int    `json:"status"`
}

func getSubList(c *gin.Context) (subList []float64) {
	cookie := c.GetHeader("Cookie")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://172.16.32.95:8080/user/subList", nil)
	// 	req, err := http.NewRequest("GET", "http://localhost:8080/user/subList", nil)
	if err == nil {
		req.Header.Set("Cookie", cookie)
		fmt.Println(cookie)
		rsp, err := client.Do(req)
		defer rsp.Body.Close()
		if err == nil && rsp.StatusCode == http.StatusOK {
			// uid = rsp.Body
			body, err := ioutil.ReadAll(rsp.Body)
			fmt.Println(string(body))
			if err == nil {
				var rsp serializer.Response
				_ = json.Unmarshal(body, &rsp)
				datas := rsp.Data.([]interface{})
				for _, data := range datas {
					subList = append(subList, data.(float64))
				}
			}
		}
	}
	return
}

func GetApplyList(c *gin.Context, uid uint) serializer.Response {
	status := c.Param("status")

	var insql string
	for _, sub := range getSubList(c) {
		if len(insql) > 0 {
			insql += ","
		}
		insql += strconv.FormatFloat(sub, 'E', -1, 64)
	}

	var applyList []Apply
	var sql string
	switch status {
	case "0": // 待处理
		sql = "select a.id, b.name as project_name, c.name, a.user_id as apply_user_id, " +
			"a.created_at as apply_date, a.status from applies a, projects b, approvals c " +
			"where a.project_id = b.id and a.approval_id = c.id " +
			"and a.status = 0 and a.user_id in (" + insql + ")"
	case "1": // 已处理
		sql = "select a.id, b.name as project_name, c.name, a.user_id as apply_user_id, " +
			"a.created_at as apply_date, a.status from applies a, projects b, approvals c " +
			"where a.project_id = b.id and a.approval_id = c.id " +
			"and a.status = 1 and a.user_id in (" + insql + ")"
	case "2": // 已发起
		sql = "select a.id, b.name as project_name, c.name, a.user_id as apply_user_id, " +
			"a.created_at as apply_date, a.status from applies a, projects b, approvals c " +
			"where a.project_id = b.id and a.approval_id = c.id " +
			"and a.user_id = " + strconv.Itoa(int(uid))
	default: // 我收到的
		sql = "select a.id, b.name as project_name, c.name, a.user_id as apply_user_id, " +
			"a.created_at as apply_date, a.status from applies a, projects b, approvals c " +
			"where a.project_id = b.id and a.approval_id = c.id"
	}
	model.DB.Raw(sql).Scan(&applyList)
	return serializer.Success(applyList)
}

type ApplyDetail struct {
	ID            uint           `json:"id"`
	ProjectName   string         `json:"projectName"`
	Name          string         `json:"name"`
	ApplyUserName string         `json:"applyUserName"`
	ApplyDate     string         `json:"applyDate"`
	HandleDate    string         `json:"handleDate"`
	Money         string         `json:"money"`
	Remarks       []RemarkDetail `json:"remarks"`
}

type RemarkDetail struct {
	ID        uint      `json:"id"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"createdAt"`
	User      string    `json:"user"`
}

func GetApplyDetail(c *gin.Context) serializer.Response {
	id := c.Param("id")

	sql := "select a.id, b.name as project_name, c.name, d.nickname as apply_user_name, " +
		"a.created_at as apply_date, a.updated_at as handle_date, a.money " +
		"from applies a, projects b, approvals c, users d " +
		"where a.project_id = b.id and a.approval_id = c.id and a.user_id = d.id and a.id = " + id

	var applyDetail ApplyDetail
	model.DB.Raw(sql).Scan(&applyDetail)

	var remarks []RemarkDetail
	// model.DB.Where("apply_id = ?", id).Find(&remarks)
	remarkSql := "select a.id, a.remark, a.created_at, b.nickname as user from remarks a, users b where a.user_id = b.id and a.apply_id = " + id
	model.DB.Raw(remarkSql).Scan(&remarks)

	applyDetail.Remarks = remarks
	return serializer.Success(applyDetail)
}

type RemarkService struct {
	Remark string `form:"remark" json:"remark" binding:"required,min=2,max=400"`
}

func (service *RemarkService) AddRemark(c *gin.Context, uid uint) serializer.Response {
	id := c.Param("id")
	idi, _ := strconv.ParseUint(id, 10, 32)

	remark := model.Remark{
		ApplyId: uint(idi),
		Remark:  service.Remark,
		UserId:  uid,
	}
	model.DB.Create(&remark)
	return serializer.Success()
}

func ChangeApplyStatus(c *gin.Context) serializer.Response {
	id, status := c.Param("id"), c.Param("status")

	var apply model.Apply

	model.DB.Where("id = ?", id).First(&apply)
	statusId, _ := strconv.Atoi(status)
	apply.Status = statusId
	model.DB.Save(apply)
	return serializer.Success()
}

type AddApplyService struct {
	ApprovalId uint    `json:"menuId"`
	ProjectId  uint    `json:"projectId"`
	Money      float32 `json:"money"`
	Remark     string  `json:"remark"`
}

func (service *AddApplyService) AddApply(uid uint) serializer.Response {
	apply := model.Apply{
		ApprovalId: service.ApprovalId,
		ProjectId:  service.ProjectId,
		Money:      service.Money,
		UserId:     uid,
	}

	model.DB.Create(&apply)

	remark := model.Remark{
		ApplyId: apply.ID,
		Remark:  service.Remark,
		UserId:  uid,
	}
	model.DB.Create(&remark)
	return serializer.Success()
}
