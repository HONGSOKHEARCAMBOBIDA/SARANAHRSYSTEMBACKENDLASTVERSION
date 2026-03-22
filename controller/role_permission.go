package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRolePermissions(c *gin.Context) {
	var input models.CreateRolePermissionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.PermissionIDs) == 0 {
		share.RespondError(c, http.StatusBadRequest, "permission_ids cannot be empty")
		return
	}

	var rolePermissions []models.RolePermission
	for _, permissionID := range input.PermissionIDs {
		rolePermissions = append(rolePermissions, models.RolePermission{
			RoleID:       uint(input.RoleID),
			PermissionID: uint(permissionID),
		})
	}

	tx := config.DB.Begin()
	if err := tx.Create(&rolePermissions).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	share.ResponeSuccess(c, 200, "បង្កើតបានជោគជ័យ")
}

func DeleteRolePermission(c *gin.Context) {
	var input models.DeleteRolePermissionsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Where("role_id = ? AND permission_id IN ?", input.RoleID, input.PermissionIDs).
		Delete(&models.RolePermission{}).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "លុបបានជោគជ័យ")
}
func GetRolePermission(c *gin.Context) {
	id := c.Param("id")
	var permissions []models.PermissionWithAssignedRole

	err := config.DB.Table("permissions").
		Select(`
            permissions.id AS id,
            permissions.name AS name,
            permissions.display_name AS display_name,
            CASE 
                WHEN role_has_permissions.permission_id IS NULL THEN false 
                ELSE true 
            END AS assigned
        `).
		Joins(`
            LEFT JOIN role_has_permissions 
            ON permissions.id = role_has_permissions.permission_id 
            AND role_has_permissions.role_id = ?
        `, id).
		Scan(&permissions).Error

	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, permissions)
}
