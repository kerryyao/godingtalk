package contact

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kerryyao/godingtalk"
)

// 创建用户
func UserCreate() {

}

// 更新用户信息
func UserUpdate() {

}

// 删除用户
func UserDelete() {

}

// 获取部门用户基础信息
func DepartmentUserSimpleList() {

}

// 获取部门用户userid列表
func DepartmentUserIDList() {

}

// 获取部门用户详情(未完成)
func DepartmentUserList(deptid, cursor, size int, orderfield, language string, accesslimit bool) (*UserListResponse, error) {
	var data UserListResponse
	if size > 100 {
		return nil, fmt.Errorf("size 最大100")
	}
	if language == "" {
		language = "zh_CN"
	}

	params := url.Values{}

	reqData := &struct {
		Cursor             int    `json:"cursor"`
		ContainAccessLimit bool   `json:"contain_access_limit"`
		Size               int    `json:"size"`
		OrderField         string `json:"order_field"`
		Language           string `json:"language"`
		DeptID             int    `json:"dept_id"`
	}{
		Cursor:             cursor,
		ContainAccessLimit: accesslimit,
		Size:               size,
		OrderField:         orderfield,
		Language:           language,
		DeptID:             deptid,
	}

	payload, err := godingtalk.HttpRequest("topapi/v2/user/list", params, reqData)
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

// 查询用户详情
func UserInfoByUserId(userid string) (*UserResponse, error) {
	var data UserResponse
	params := url.Values{}
	reqData := &struct {
		UserID string `json:"userid"`
	}{
		UserID: userid,
	}
	payload, err := godingtalk.HttpRequest("topapi/v2/user/get", params, reqData)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 根据unionid获取用户userid
func UseridByUnionId(unionid string) (*UserIDResponse, error) {
	var data UserIDResponse

	params := url.Values{}
	reqData := &struct {
		UnionID string `json:"unionid"`
	}{
		UnionID: unionid,
	}

	payload, err := godingtalk.HttpRequest("topapi/user/getbyunionid", params, reqData)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// 根据手机号查询用户
func UseridByMobile(mobile string) (*UserIDResponse, error) {
	var data UserIDResponse

	params := url.Values{}
	reqData := &struct {
		Mobile string `json:"mobile"`
	}{
		Mobile: mobile,
	}
	payload, err := godingtalk.HttpRequest("topapi/v2/user/getbymobile", params, reqData)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取员工人数
func UseridCount(active bool) (*UserCountResponse, error) {
	var data UserCountResponse

	params := url.Values{}
	reqData := &struct {
		OnlyActive bool `json:"only_active"`
	}{
		OnlyActive: active,
	}
	payload, err := godingtalk.HttpRequest("topapi/user/count", params, reqData)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取管理员列表
func AdminList() {

}

// 获取管理员通讯录权限范围
func AdminPrivilege() {

}
