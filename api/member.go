package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

/**
采用连接邀请方式进入，暂时不实现
*/
func (d DidiEs) MemberSingle() {

}

/**
修改接口
*/
func (d DidiEs) MemberEdit(email string, memberData Member) (err error) {

	member, err := d.MemberGetSingle(email)
	if err != nil {
		return
	}

	data, err := json.Marshal(memberData)
	if err != nil {
		return
	}
	_, err = d.Post("/river/Member/edit", &MemberEditResuest{
		MemberId: member.Id,
		Data:     string(data),
	})
	if err != nil {
		return
	} else {
		return nil
	}

}

/**
获取单个用户接口
*/
func (d DidiEs) MemberGetSingle(email string) (memberItem *MemberItem, err error) {
	ret, err := d.Get("/river/Member/get", &MemberGetRequest{
		Status: MemberStatusNormal,
		Email:  email,
		Offset: 0,
		Length: 100,
	})
	if err != nil {
		log.Print(err.Error())
		return
	} else {
		data := MemberList{}
		json.Unmarshal(ret, &data)
		if data.Total == 1 {
			memberItem = &data.Records[0]
			return
		} else if data.Total == 0 {
			return nil, ErrorMemberNotFount
		} else {
			return nil, ErrorMemberFound
		}
	}
}

/**
查询接口
*/
func (d DidiEs) MemberGet() {
	ret, err := d.Get("/river/Member/get", &MemberGetRequest{
		Status: MemberStatusNormal,
		Offset: 0,
		Length: 100,
	})
	if err != nil {
		log.Print(err.Error())
	} else {
		data := MemberList{}
		json.Unmarshal(ret, &data)
		log.Print(string(ret))
		if data.Total != 0 {
			for _, item := range data.Records {
				itemjson, _ := json.Marshal(item)
				log.Info(string(itemjson))
			}
		}
	}
}
