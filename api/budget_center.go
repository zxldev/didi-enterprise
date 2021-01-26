package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (d DidiEs) AddDepartment(req *BudgetCenterAddRequest) {
	ret, err := d.Post("/river/BudgetCenter/add", req)
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print("添加成功:", string(ret))
	}
}

func (d DidiEs) GetDepartment(req *BudgetCenterGetRequest) (department []BudgetCenter) {
	ret, err := d.Get("/river/BudgetCenter/get", req)

	if err != nil {
		log.Print(err.Error())
		return nil
	} else {
		data := BudgetCenterList{}
		json.Unmarshal(ret, &data)
		log.Print(string(ret))
		if data.Total != "0" {
			return data.Records
		}
	}
	return nil
}

func (d DidiEs) DelDepartment(id string) {
	_, err := d.Post("/river/BudgetCenter/del", &BudgetCenterDelRequest{
		Id: id,
	})
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print("删除成功:", id)
	}
}
