package main

import (
	"flag"
	"github.com/mozillazg/go-pinyin"
	"github.com/zxldev/didi-enterprise/api"
	"strings"
)

func main() {
	api.DidiEsClient.MemberGet()
	testMember()

}

func init() {
	client_id := flag.String("client_id", "", "client_id")
	client_secret := flag.String("client_secret", "", "client_secret")
	sign_key := flag.String("sign_key", "", "sign_key")
	admin_phone := flag.String("admin_phone", "", "admin_phone")
	company_id := flag.String("company_id", "", "company_id")

	flag.Parse()
	api.DidiEsClient.Init(*client_id, *client_secret, *sign_key, *admin_phone, *company_id)
}

func testMember() {
	hans := "部门"
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	pingyinDepartment := strings.ToUpper(strings.Join(pinyin.LazyPinyin(hans, a), ""))

	api.DidiEsClient.MemberEdit("admin@kingsoft.com", api.Member{
		Realname:       "zxldev",
		EmployeeNumber: "066666",
		SystemRole:     0,
		Residentsname:  "北京",
		IsRemark:       2,
		BudgetCenterId: api.DidiEsClient.DepartmentMap[pingyinDepartment], //部门
		RegulationId:   "1125909344217498_1125909344276989",               //规则
	})
}

func testDepartment() {
	hans := "部门"
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	pingyinDepartment := strings.ToUpper(strings.Join(pinyin.LazyPinyin(hans, a), ""))

	api.DidiEsClient.AddDepartment(&api.BudgetCenterAddRequest{
		Name:        pingyinDepartment,
		Type:        api.BudgetTypeDepartment,
		BudgetCycle: api.BudgetCycleNoLimit,
		TotalQuota:  0.0,
		OutBudgetId: pingyinDepartment,
	})

	departments := api.DidiEsClient.GetDepartment(&api.BudgetCenterGetRequest{
		Name:   "删除的部门",
		Offset: 0,
		Length: 100,
	})

	for _, department := range departments {
		api.DidiEsClient.DelDepartment(department.Id)
	}

	api.DidiEsClient.GetDepartment(&api.BudgetCenterGetRequest{
		Offset: 0,
		Length: 100,
	})
}
