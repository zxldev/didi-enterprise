package api

type BaseParams struct {
	ClientId    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	CompanyId   string `json:"company_id"`
}

type BaseParamsBuilder interface {
	BuildBaseParams(ClientId, AccessToken, CompanyId string)
}

type BaseReponse struct {
	Errorno int         `json:"errorno"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

type AuthorizeRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Phone        string `json:"phone"`
}

type AuthorizeResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

/**
请求参数
名称	类型	必选	描述

name	string	yes	名称
type	int	yes	类型（1部门 2项目）
budget_cycle	int	yes	预算周期（0不限期；1自然月）
total_quota	float	yes	总金额（0表示不限额度，单位元，小数点后保留两位精度）
out_budget_id	string	no	编号(部门为必填，项目为选填)
leader_id	string	no	主管，通过员工-详情API获得对应id
parent_id	int	no	上级部门或项目的id(部门为选填，默认为顶级部门的id；项目为选填)
member_used	int	no	是否仅项目成员可用（0不是；1是）默认为0；当类型为项目时，此参数有效
start_date	string	no	项目开始日期（格式举例：2018-01-02），可以为空；当类型为项目时，此参数有效
expiry_date	string	no	项目结束日期（格式举例：2018-01-02），可以为空；当类型为项目时，此参数有效
注意事项
部门数量上限是2000，如贵司超过2000个部门，请联系相应的对接同学添加白名单进行扩展
部门添加时，请不要并发操作，防止部门添加失败
*/

type BudgetType int
type BudgetCycle int
type IsLimitQuota int

const (
	BudgetTypeDepartment BudgetType = 1
	BudgetTypeProject    BudgetType = 2

	BudgetCycleNoLimit BudgetCycle = 0
	BudgetCycleMonthly BudgetCycle = 1

	IsLimitQuotaNoLimit IsLimitQuota = 0
	IsLimitQuotaLimit   IsLimitQuota = 1
)

type BudgetCenterAddRequest struct {
	BaseParams  `json:",inline"`
	Name        string      `json:"name"`
	Type        BudgetType  `json:"type"`
	BudgetCycle BudgetCycle `json:"budget_cycle"`
	TotalQuota  float64     `json:"total_quota"`
	OutBudgetId string      `json:"out_budget_id,omitempty"`
	LeaderId    string      `json:"leader_id,omitempty"`
	ParentId    int         `json:"parent_id,omitempty"`
	MemberUsed  int         `json:"member_used,omitempty"`
	StartDate   string      `json:"start_date,omitempty"`
	ExpiryDate  string      `json:"expiry_date,omitempty"`
}

type BudgetCenterGetRequest struct {
	BaseParams  `json:",inline"`
	Name        string     `json:"name,omitempty"`
	Id          string     `json:"id,omitempty"`
	Type        BudgetType `json:"type,omitempty"`
	IsExactName string     `json:"is_exact_name,omitempty"`
	OutBudgetId string     `json:"out_budget_id,omitempty"`
	Offset      int        `json:"offset"`
	Length      int        `json:"length"`
}

type BudgetCenterDelRequest struct {
	BaseParams `json:",inline"`
	Id         string `json:"id,omitempty"`
}

type BudgetCenterList struct {
	Total   string         `json:"total"`
	Records []BudgetCenter `json:"records"`
}

/**
@see http://open.es.xiaojukeji.com/doc/erpapi/budgetCenter/list.html
id	string	id
name	string	名称
type	int	类型（1部门 2项目）
budget_cycle	int	预算周期（0不限期；1自然月）
out_budget_id	string	外部成本中心id
total_quota	number	总金额（0表示不限额度）
is_limit_quota	int	是否限额（0不限制，1限制）
member_num	int	在使用人数
available_quota	number	可用金额
freeze_quota	number	冻结金额
leader_id	string	主管
parent_id	int	上级部门或项目的id
member_used	int	是否仅项目成员可用（0不是；1是）；当类型为项目时，此参数有效
start_date	string	项目开始日期；当类型为项目时，此参数有效
expiry_date	string	项目结束日期；当类型为项目时，此参数有效
*/

type BudgetCenter struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Type        BudgetType  `json:"type"`
	BudgetCycle BudgetCycle `json:"budget_cycle"`
	TotalQuota  float64     `json:"total_quota"`
	OutBudgetId string      `json:"out_budget_id,omitempty"`
	LeaderId    string      `json:"leader_id,omitempty"`
	ParentId    int         `json:"parent_id,omitempty"`
	MemberUsed  int         `json:"member_used,omitempty"`
	StartDate   string      `json:"start_date,omitempty"`
	ExpiryDate  string      `json:"expiry_date,omitempty"`

	IsLimitQuota IsLimitQuota `json:"is_limit_quota"`
	MemberNum    int          `json:"member_num"`
}
