package dtoroles

type CreateRoleRequest struct {
	ParentId         string `json:"parentId"`
	Topic            string `json:"topic" validate:"required"`
	Rules            string `json:"rules" validate:"required"`
	Goals            string `json:"goals" validate:"required"`
	ChildDescription string `json:"childDescription" validate:"required"`
	RoleName         string `json:"roleName" validate:"required"`
	RoleDescription  string `json:"roleDescription" validate:"required"`
}
