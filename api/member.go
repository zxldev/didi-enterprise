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
func (d DidiEs) MemberEditByEmail(email string, memberData Member) (err error) {

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
	return err

}

/**
修改接口
*/
func (d DidiEs) MemberEdit(id string, memberData Member) (err error) {

	data, err := json.Marshal(memberData)
	if err != nil {
		return
	}
	_, err = d.Post("/river/Member/edit", &MemberEditResuest{
		MemberId: id,
		Data:     string(data),
	})
	return err

}

/**
获取单个用户接口
// ！！ 用户可以在后台更改自己的邮箱，请以用户编码为准，只有首次才能以用户邮箱为准
*/
func (d DidiEs) MemberGetSingleByEmployCode(code string) (memberItem *MemberItem, err error) {
	ret, err := d.Get("/river/Member/get", &MemberGetRequest{
		Status:         MemberStatusNormal,
		EmployeeNumber: code,
		Offset:         0,
		Length:         100,
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
func (d DidiEs) MemberGet(member *MemberGetRequest) (memberList []MemberItem, err error) {
	ret, err := d.Get("/river/Member/get", member)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	} else {
		data := MemberList{}
		json.Unmarshal(ret, &data)
		log.Debug(string(ret))
		return data.Records, nil
	}
}

/**
批量获取所有信息
*/
func (d DidiEs) MemberGetAll(member *MemberGetRequest) (memberList []MemberItem, err error) {

	offset := 0
	memberList = []MemberItem{}
	for {
		clone := *member
		clone.Offset = offset
		clone.Length = 100
		ret, err := d.Get("/river/Member/get", &clone)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		} else {
			data := MemberList{}
			json.Unmarshal(ret, &data)
			if data.Total < 1 {
				break
			}
			memberList = append(memberList, data.Records...)
			if data.Total < 100 {
				break
			}
		}
		offset += 100
	}
	return memberList, nil
}

/**
删除用户 删除失败也会离职
*/
func (d DidiEs) MemberDelete(MemberId string) {
	ret, err := d.Post("/river/Member/del", &MemberDelRequest{
		MemberId: MemberId,
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
