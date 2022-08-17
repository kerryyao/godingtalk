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
	Result struct {
		AutoAddUser         bool          `json:"auto_add_user"`
		AutoApproveApply    bool          `json:"auto_approve_apply"`
		Brief               string        `json:"brief"`
		CreateDeptGroup     bool          `json:"create_dept_group"`
		DeptGroupChatID     string        `json:"dept_group_chat_id"`
		DeptID              int           `json:"dept_id"`
		DeptPermits         []interface{} `json:"dept_permits"`
		GroupContainSubDept bool          `json:"group_contain_sub_dept"`
		HideDept            bool          `json:"hide_dept"`
		Name                string        `json:"name"`
		Order               int           `json:"order"`
		OrgDeptOwner        string        `json:"org_dept_owner"`
		OuterDept           bool          `json:"outer_dept"`
		OuterPermitDepts    []interface{} `json:"outer_permit_depts"`
		OuterPermitUsers    []interface{} `json:"outer_permit_users"`
		ParentID            int           `json:"parent_id"`
		UserPermits         []interface{} `json:"user_permits"`
	} `json:"result"`
}

type DepartmentListResponse struct {
	godingtalk.OAPIResponse
	Result []Department `json:"result"`
}
