package contact

import (
	"encoding/json"
	"net/url"

	"github.com/kerryyao/godingtalk"
)

// 创建部门
func DepartmentCreate() {

}

// 获取部门列表
func DepartmentList(parentid int) (*DepartmentListResponse, error) {
	var data DepartmentListResponse
	params := url.Values{}
	reqData := &struct {
		DeptID   int    `json:"dept_id"`
		Language string `json:"language"`
	}{
		DeptID:   parentid,
		Language: "zh_CN",
	}
	payload, err := godingtalk.HttpRequest("v2/department/listsub", params, reqData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 获取部门详情
func DepartmentDetail(id int) (*Department, error) {
	var data Department
	params := url.Values{}
	reqData := &struct {
		DeptID   int    `json:"dept_id"`
		Language string `json:"language"`
	}{
		DeptID:   id,
		Language: "zh_CN",
	}
	payload, err := godingtalk.HttpRequest("topapi/v2/department/get", params, reqData)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
