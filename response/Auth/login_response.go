package auth

import userpart "HRbackend/response/UserPart"

type LoginResponse struct {
	ID     uint                        `json:"id"`
	Name   string                      `json:"name"`
	Phone  string                      `json:"phone"`
	RoleID uint                        `json:"role_id"`
	Parts  []userpart.UserPartResponse `json:"parts"`
	Token  string                      `json:"token"`
}
