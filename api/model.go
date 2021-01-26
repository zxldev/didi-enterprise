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
type MemberStatus string

const (
	BudgetTypeDepartment BudgetType = 1
	BudgetTypeProject    BudgetType = 2

	BudgetCycleNoLimit BudgetCycle = 0
	BudgetCycleMonthly BudgetCycle = 1

	IsLimitQuotaNoLimit IsLimitQuota = 0
	IsLimitQuotaLimit   IsLimitQuota = 1

	MemberStatusNormal  MemberStatus = "1"
	MemberStatusDelete  MemberStatus = "3"
	MemberStatusLeave   MemberStatus = "4"
	MemberStatusNoPhone MemberStatus = "6"
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

type BudgetCenterEditRequest struct {
	BaseParams  `json:",inline"`
	Id          string      `json:"id"`
	Name        string      `json:"name,omitempty"`
	BudgetCycle BudgetCycle `json:"budget_cycle,omitempty"`
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

type MemberSingle struct {
	BaseParams `json:",inline"`
	Data       string `json:"data"`
}

type MemberEditResuest struct {
	BaseParams `json:",inline"`
	MemberId   string `json:"member_id"`
	Data       string `json:"data"`
}

/**
   员工信息
名称	类型	必选	描述
phone	string	no	员工手机号
realname	string	no	员工姓名
employee_number	string	no	员工ID（员工在公司的员工号）
email	string	no	邮箱
department	string	no	部门名称（老），后续此参数会去掉
branch_name	string	no	所在分公司名称（老），后续此参数会去掉
system_role	int	no	系统角色(0-车辆预定人员，1-普通管理员，2-超级管理员)
role_ids	string	no	角色（默认为员工），通过角色获取API 。可以填多个，以_分隔。初始管理员、主管不支持新增
immediate_superior_phone	string	no	员工直属上级的手机号码（需为本企业中已存在账号），直属上级可在审批流中担任审批人
residentsname	string	no	常驻地中文
use_company_money	int	no	是否企业支付余额（0-否，1-是）
total_quota	string	no	每月配额
is_remark	int	no	叫车时备注信息是否必填(0-选填，1-必填，2-按制度填写)
budget_center_id	bigint	no	所在部门ID（新），默认为1（企业）。通过成本中心查询api获取id（类型为1）
regulation_id	string	no	用车制度ID（从 9.1用车制度查询 中选取，可以填多个，以_分隔，如 123_456_789）注:如不传用车制度ID,无法使用企业支付
project_ids	string	no	所在项目ID（新）。可以填多个，以_分隔。通过成本中心查询api获取id（类型为2）
set_dismiss_time	string	no	设置员工离职日期，到期后自动加入已离职名单，不传或为空时认为不设置离职时间，格式为 "2018-07-01“
*/
type Member struct {
	Phone                  string `json:"phone,omitempty"`
	Realname               string `json:"realname,omitempty"`
	EmployeeNumber         string `json:"employee_number,omitempty"`          //	no	员工ID（员工在公司的员工号）
	Email                  string `json:"email,omitempty"`                    //	no	邮箱
	Department             string `json:"department,omitempty"`               //	no	部门名称（老），后续此参数会去掉
	BranchName             string `json:"branch_name,omitempty"`              //	no	所在分公司名称（老），后续此参数会去掉
	SystemRole             int    `json:"system_role,omitempty"`              //  no	系统角色(0-车辆预定人员，1-普通管理员，2-超级管理员)
	RoleIds                string `json:"role_ids,omitempty"`                 //	no	角色（默认为员工），通过角色获取API 。可以填多个，以_分隔。初始管理员、主管不支持新增
	ImmediateSuperiorPhone string `json:"immediate_superior_phone,omitempty"` // no	员工直属上级的手机号码（需为本企业中已存在账号），直属上级可在审批流中担任审批人
	Residentsname          string `json:"residentsname,omitempty"`            //	no	常驻地中文
	UseCompanyMoney        int    `json:"use_company_money,omitempty"`        // no	是否企业支付余额（0-否，1-是）
	TotalQuota             string `json:"total_quota,omitempty"`              // no	每月配额
	IsRemark               int    `json:"is_remark,omitempty"`                //     no	叫车时备注信息是否必填(0-选填，1-必填，2-按制度填写)
	BudgetCenterId         string `json:"budget_center_id,omitempty"`         // no	所在部门ID（新），默认为1（企业）。通过成本中心查询api获取id（类型为1）
	RegulationId           string `json:"regulation_id,omitempty"`            // no	用车制度ID（从 9.1用车制度查询 中选取，可以填多个，以_分隔，如 123_456_789）注:如不传用车制度ID,无法使用企业支付
	ProjectIds             string `json:"project_ids,omitempty"`              //	no	所在项目ID（新）。可以填多个，以_分隔。通过成本中心查询api获取id（类型为2）
	SetDismissTime         string `json:"set_dismiss_time,omitempty"`         //no	设置员工离职日期，到期后自动加入已离职名单，不传或为空时认为不设置离职时间，格式为 "2018-07-01“
}

type MemberDelRequest struct {
	BaseParams `json:",inline"`
	MemberId   string `json:"member_id"`
}

type MemberGetRequest struct {
	BaseParams     `json:",inline"`
	Phone          string       `json:"phone,omitempty"`
	Realname       string       `json:"realname,omitempty"`
	EmployeeNumber string       `json:"employee_number,omitempty"` //	no	员工ID（员工在公司的员工号）
	Email          string       `json:"email,omitempty"`           //	no	邮箱
	Status         MemberStatus `json:"status"`
	Offset         int          `json:"offset"`
	Length         int          `json:"length"`
}

type MemberList struct {
	Total   int          `json:"total"`
	Records []MemberItem `json:"records"`
}

type MemberItem struct {
	Id                     string   `json:"id"`
	Phone                  string   `json:"phone,omitempty"`
	Realname               string   `json:"realname,omitempty"`
	EmployeeNumber         string   `json:"employee_number,omitempty"`            //	no	员工ID（员工在公司的员工号）
	Email                  string   `json:"email,omitempty"`                      //	no	邮箱
	Department             string   `json:"department,omitempty"`                 //	no	部门名称（老），后续此参数会去掉
	BranchName             string   `json:"branch_name,omitempty"`                //	no	所在分公司名称（老），后续此参数会去掉
	SystemRole             int      `json:"system_role,omitempty"`                //  no	系统角色(0-车辆预定人员，1-普通管理员，2-超级管理员)
	RoleIds                string   `json:"role_ids,omitempty"`                   //	no	角色（默认为员工），通过角色获取API 。可以填多个，以_分隔。初始管理员、主管不支持新增
	ImmediateSuperiorPhone string   `json:"immediate_superior_phone,omitempty"`   // no	员工直属上级的手机号码（需为本企业中已存在账号），直属上级可在审批流中担任审批人
	Residentsname          string   `jsonregulation_id:"residentsname,omitempty"` //	no	常驻地中文
	UseCompanyMoney        int      `json:"use_company_money,omitempty"`          // no	是否企业支付余额（0-否，1-是）
	TotalQuota             string   `json:"total_quota,omitempty"`                // no	每月配额
	IsRemark               int      `json:"is_remark,omitempty"`                  //     no	叫车时备注信息是否必填(0-选填，1-必填，2-按制度填写)
	BudgetCenterId         string   `json:"budget_center_id,omitempty"`           // no	所在部门ID（新），默认为1（企业）。通过成本中心查询api获取id（类型为1）
	RegulationId           []string `json:"regulation_id,omitempty"`              // no	用车制度ID（从 9.1用车制度查询 中选取，可以填多个，以_分隔，如 123_456_789）注:如不传用车制度ID,无法使用企业支付
	SetDismissTime         string   `json:"set_dismiss_time,omitempty"`
	UseCarConfig           []string `json:"use_car_config"`
	DismissTime            string   `json:"dismiss_time"`
}
