package contact

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kerryyao/godingtalk"
)

// DepartmentList is 获取部门列表
func DepartmentList() (*DepartmentListResponse, error) {
	var data DepartmentListResponse
	payload, err := godingtalk.HttpRequest("department/list", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DepartmentDetail is 获取部门详情
func DepartmentDetail(id int) (*Department, error) {
	var data Department
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	payload, err := godingtalk.HttpRequest("department/get", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
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
	payload, err := godingtalk.HttpRequest("user/list", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// CreateChat is
func CreateChat(name string, owner string, useridlist []string) (*string, error) {
	var data struct {
		godingtalk.OAPIResponse
		Chatid string
	}
	request := map[string]interface{}{
		"name":       name,
		"owner":      owner,
		"useridlist": useridlist,
	}
	payload, err := godingtalk.HttpRequest("chat/create", nil, request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data.Chatid, nil
}

// UserInfoByUserId 获取用户详情
func UserInfoByUserId(userid string) (*User, error) {
	var data User
	params := url.Values{}
	params.Add("userid", userid)
	payload, err := godingtalk.HttpRequest("user/get", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// UseridByUnionId 通过UnionId获取玩家Userid
func UseridByUnionId(unionid string) (*string, error) {
	var data struct {
		godingtalk.OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("unionid", unionid)
	payload, err := godingtalk.HttpRequest("user/getUseridByUnionid", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}

	return &data.UserID, nil
}

// UseridByMobile 通过手机号获取Userid
func UseridByMobile(mobile string) (*string, error) {
	var data struct {
		godingtalk.OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("mobile", mobile)
	payload, err := godingtalk.HttpRequest("user/get_by_mobile", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data.UserID, nil
}
