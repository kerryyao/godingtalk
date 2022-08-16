package contact

import "github.com/kerryyao/godingtalk"

type User struct {
	godingtalk.OAPIResponse
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
	godingtalk.OAPIResponse
	HasMore  bool
	Userlist []User
}

type Department struct {
	godingtalk.OAPIResponse
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
	godingtalk.OAPIResponse
	Departments []Department `json:"department"`
}
