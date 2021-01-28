package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zxldev/didi-enterprise/util"
	"net/http"

	"sync"
	"time"
)

type DidiEs struct {
	ClientId      string    `json:"client_id"`
	ClientSecret  string    `json:"client_secret"`
	SignKey       string    `json:"sign_key"`
	AdminPhone    string    `json:"admin_phone"`
	CompanyId     string    `json:"company_id"`
	AccessToken   string    `json:"access_token"`
	TokenExpire   time.Time `json:"token_expire"`
	ApiUrl        string    `json:"api_url"`
	TokenLock     sync.WaitGroup
	DepartmentMap map[string]string `json:"department_map"`
}

var DidiEsClient DidiEs

func (d *DidiEs) GenDepartmentMap() error {
	ds := d.GetDepartment(&BudgetCenterGetRequest{
		Offset: 0,
		Length: 100,
	})
	d.DepartmentMap = map[string]string{}
	if len(ds) > 0 {
		for _, i := range ds {
			d.DepartmentMap[i.OutBudgetId] = i.Id
		}
	}
	return nil
}

func (d *DidiEs) Init(ClientId, ClientSecret, SignKey, AdminPhone, CompanyId, ApiUrl string) {
	if ApiUrl == "" {
		d.ApiUrl = ServerApi
	} else {
		d.ApiUrl = ApiUrl
	}
	d.ClientId = ClientId
	d.ClientSecret = ClientSecret
	d.SignKey = SignKey
	d.AdminPhone = AdminPhone
	d.CompanyId = CompanyId
	d.GenDepartmentMap()
}

/**
判断是否同一个部门
*/
func (d *DidiEs) IsSameDepartment(OutBudgetId, DepartmentId string) bool {
	if v, ok := d.DepartmentMap[OutBudgetId]; ok {
		return v == DepartmentId
	} else {
		return false
	}
}

func (d *DidiEs) GetToken() (string, error) {
	if d.AccessToken == "" || time.Now().Unix() > d.TokenExpire.Unix() {
		//获取新Token，并且重置过期时间
		r, err := d.PostAuth("/river/Auth/authorize", AuthorizeRequest{
			ClientId:     d.ClientId,
			ClientSecret: d.ClientSecret,
			GrantType:    "client_credentials",
			Phone:        d.AdminPhone,
		})
		if err != nil {
			return "", ErrorNetWork
		}
		token := AuthorizeResp{}

		err = json.NewDecoder(r.Body).Decode(&token)
		if err != nil {
			return "", ErrorDecode
		}
		if token.AccessToken == "" {
			return "", ErrorGetAccessToken
		}
		d.AccessToken = token.AccessToken
		d.TokenExpire = time.Now().Add(time.Second * time.Duration(token.ExpiresIn-30))
	}

	return d.AccessToken, nil

}

const (
	ServerApi = "http://api.es.xiaojukeji.com"
)

func (d *BaseParams) BuildBaseParams(ClientId, AccessToken, CompanyId string) {
	d.ClientId = ClientId
	d.CompanyId = CompanyId
	d.AccessToken = AccessToken
}

func (d *DidiEs) PostAuth(url string, data interface{}) (resp *http.Response, err error) {
	return http.Post(d.ApiUrl+url, "application/json", bytes.NewReader(util.SignRequest(data, d.SignKey)))
}

func (d *DidiEs) Post(url string, data BaseParamsBuilder) (ret []byte, err error) {
	token, err := d.GetToken()
	if err != nil {
		return nil, err
	}
	data.BuildBaseParams(d.ClientId, token, d.CompanyId)
	resp, err := http.Post(d.ApiUrl+url, "application/json", bytes.NewReader(util.SignRequest(data, d.SignKey)))

	if err != nil {
		return nil, ErrorNetWork
	}

	baseresp := BaseResponse{}
	json.NewDecoder(resp.Body).Decode(&baseresp)

	if baseresp.Errorno == 0 {
		return json.Marshal(baseresp.Data)
	} else {
		return nil, errors.New(baseresp.Errmsg)
	}
}

func (d *DidiEs) Get(url string, data BaseParamsBuilder) (ret []byte, err error) {
	token, err := d.GetToken()
	if err != nil {
		return nil, err
	}
	data.BuildBaseParams(d.ClientId, token, d.CompanyId)
	resp, err := http.Get(d.ApiUrl + url + "?" + util.SignGetRequest(data, d.SignKey))

	if err != nil {
		return nil, ErrorNetWork
	}

	baseresp := BaseResponse{}
	json.NewDecoder(resp.Body).Decode(&baseresp)

	if baseresp.Errorno == 0 {
		return json.Marshal(baseresp.Data)
	} else {
		return nil, errors.New(baseresp.Errmsg)
	}
}
