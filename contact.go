package godingtalk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type User struct {
	OAPIResponse
	Userid     string
	Name       string
	Mobile     string
	Tel        string
	Remark     string
	Order      int
	IsAdmin    bool
	IsBoss     bool
	IsLeader   bool
	Active     bool
	Department []int
	Position   string
	Email      string
	OrgEmail   string
	Avatar     string
	Extattr    interface{}
}

type UserListResponse struct {
	OAPIResponse
	HasMore  bool
	Userlist []User
}

type Department struct {
	OAPIResponse
	Id                    int
	Name                  string
	ParentId              int
	Order                 int
	DeptPerimits          string
	UserPerimits          string
	OuterDept             bool
	OuterPermitDepts      string
	OuterPermitUsers      string
	OrgDeptOwner          string
	DeptManagerUseridList string
}

type DepartmentListResponse struct {
	OAPIResponse
	Departments []Department `json:"department"`
}

// DepartmentList is 获取部门列表
func DepartmentList() (*DepartmentListResponse, error) {
	var data DepartmentListResponse
	payload, err := httpRequest("department/list", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, err
}

// DepartmentDetail is 获取部门详情
func DepartmentDetail(id int) (*Department, error) {
	var data Department
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	payload, err := httpRequest("department/get", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, err
}

// UserList is 获取部门成员
func UserList(departmentID, offset, size int) (*UserListResponse, error) {
	var data UserListResponse
	if size > 100 {
		return nil, fmt.Errorf("size 最大100")
	}

	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	params.Add("offset", fmt.Sprintf("%d", offset))
	params.Add("size", fmt.Sprintf("%d", size))
	payload, err := httpRequest("user/list", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, err
}

// CreateChat is
func CreateChat(name string, owner string, useridlist []string) (*string, error) {
	var data struct {
		OAPIResponse
		Chatid string
	}
	request := map[string]interface{}{
		"name":       name,
		"owner":      owner,
		"useridlist": useridlist,
	}
	payload, err := httpRequest("chat/create", nil, request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data.Chatid, err
}

// UserInfoByCode 校验免登录码并换取用户身份
func UserInfoByCode(code string) (*User, error) {
	var data User
	params := url.Values{}
	params.Add("code", code)
	payload, err := httpRequest("user/getuserinfo", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, err
}

// UserInfoByUserId 获取用户详情
func UserInfoByUserId(userid string) (*User, error) {
	var data User
	params := url.Values{}
	params.Add("userid", userid)
	payload, err := httpRequest("user/get", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, err
}

// UseridByUnionId 通过UnionId获取玩家Userid
func UseridByUnionId(unionid string) (*string, error) {
	var data struct {
		OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("unionid", unionid)
	payload, err := httpRequest("user/getUseridByUnionid", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}

	return &data.UserID, err
}

// UseridByMobile 通过手机号获取Userid
func UseridByMobile(mobile string) (*string, error) {
	var data struct {
		OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("mobile", mobile)
	payload, err := httpRequest("user/get_by_mobile", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data.UserID, err
}
