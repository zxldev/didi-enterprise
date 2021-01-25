package main

import (
	"flag"
	"github.com/zxldev/didi-enterprise/api"
)



func main() {
	client_id:= flag.String("client_id","","client_id")
	client_secret:=flag.String("client_secret","","client_secret")
	sign_key:=flag.String("sign_key","","sign_key")
	admin_phone:=flag.String("admin_phone","","admin_phone")
	company_id:=flag.String("company_id","","company_id")

	flag.Parse()
	api.DidiEsClient.Init(*client_id,*client_secret,*sign_key,*admin_phone,*company_id)
	api.DidiEsClient.AddDepartment(&api.BudgetCenterAddRequest{
		Name:"中文测试3",
		Type:api.BudgetTypeDepartment,
		BudgetCycle:api.BudgetCycleNoLimit,
		TotalQuota:0.0,
	})

	ids := api.DidiEsClient.GetDepartment(&api.BudgetCenterGetRequest{
		Name: " 中文测试",
		Offset:0,
		Length:100,
	})

	for _,id := range ids {
		api.DidiEsClient.DelDepartment(id)
	}

	api.DidiEsClient.GetDepartment(&api.BudgetCenterGetRequest{
		Offset:0,
		Length:100,
	})
}
