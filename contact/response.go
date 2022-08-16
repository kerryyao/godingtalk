package contact

import "github.com/kerryyao/godingtalk"

type UserCountResponse struct {
	godingtalk.OAPIResponse
	Result struct {
		Count int `json:"count"`
	} `json:"result"`
}

type UserIDResponse struct {
	godingtalk.OAPIResponse
	Result struct {
		ContactType string `json:"contact_type,omitempty"`
		Userid      string `json:"userid"`
	} `json:"result"`
}

type UserResponse struct {
	godingtalk.OAPIResponse
	Result struct {
		Active        bool   `json:"active"`
		Admin         bool   `json:"admin"`
		Avatar        string `json:"avatar"`
		Boss          bool   `json:"boss"`
		DeptIDList    []int  `json:"dept_id_list"`
		DeptOrderList []struct {
			DeptID int   `json:"dept_id"`
			Order  int64 `json:"order"`
		} `json:"dept_order_list"`
		Email            string `json:"email"`
		ExclusiveAccount bool   `json:"exclusive_account"`
		HideMobile       bool   `json:"hide_mobile"`
		JobNumber        string `json:"job_number"`
		LeaderInDept     []struct {
			DeptID int  `json:"dept_id"`
			Leader bool `json:"leader"`
		} `json:"leader_in_dept"`
		Mobile     string `json:"mobile"`
		Name       string `json:"name"`
		RealAuthed bool   `json:"real_authed"`
		Remark     string `json:"remark"`
		RoleList   []struct {
			GroupName string `json:"group_name"`
			ID        int    `json:"id"`
			Name      string `json:"name"`
		} `json:"role_list"`
		Senior      bool   `json:"senior"`
		StateCode   string `json:"state_code"`
		Telephone   string `json:"telephone"`
		Title       string `json:"title"`
		UnionEmpExt struct {
		} `json:"union_emp_ext"`
		Unionid   string `json:"unionid"`
		Userid    string `json:"userid"`
		WorkPlace string `json:"work_place"`
	} `json:"result"`
}

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
