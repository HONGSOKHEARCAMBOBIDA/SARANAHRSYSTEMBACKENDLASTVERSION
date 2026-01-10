package service

import (
	"errors"
	"time"

	"HRbackend/config"
	models "HRbackend/model"
	authReq "HRbackend/request/Auth"
	authRes "HRbackend/response/Auth"
	userpart "HRbackend/response/UserPart"
	"HRbackend/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(input authReq.LoginReq) (*authRes.LoginResponse, error)
}

type authservice struct {
	db *gorm.DB
}

func NewAuthService() AuthService {
	return &authservice{
		db: config.DB,
	}
}

func (s *authservice) Login(input authReq.LoginReq) (*authRes.LoginResponse, error) {

	key := "login_attempt:" + input.Contact
	attempts, _ := utils.Redis.Get(utils.Ctx, key).Int()
	if attempts >= 5 {
		return nil, errors.New("អ្នកព្យាយាមចូលច្រើនពេក សូមព្យាយាមម្តងទៀតក្រោយ 10 នាទី")
	}
	// 1. Find user
	var user models.User
	if err := s.db.
		Where("(contact = ? OR email = ? OR username = ?) AND is_active = ?",
			input.Contact, input.Contact, input.Contact, 1).
		First(&user).Error; err != nil {

		return nil, errors.New("ព័ត៌មានមិនត្រឹមត្រូវ ឬ អ្នកប្រើប្រាស់ត្រូវបានបិទគណនី")
	}

	// 2. Check password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {

		utils.Redis.Incr(utils.Ctx, key)
		utils.Redis.Expire(utils.Ctx, key, 10*time.Minute)

		return nil, errors.New("ព័ត៌មានមិនត្រឹមត្រូវ")
	}
	utils.Redis.Del(utils.Ctx, key)

	// 3. Get user parts
	var userParts []userpart.UserPartResponse
	if err := s.db.Table("user_parts up").
		Select("up.id AS id, p.id AS part_id, p.name AS part_name").
		Joins("JOIN parts p ON p.id = up.part_id").
		Where("up.user_id = ?", user.ID).
		Scan(&userParts).Error; err != nil {

		return nil, err
	}

	// 4. Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"phone":   user.Contact,
		"role_id": user.RoleID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(utils.Jwtkey)
	if err != nil {
		return nil, err
	}

	// 5. Build response
	resp := &authRes.LoginResponse{
		ID:     user.ID,
		Name:   user.UserName,
		Phone:  user.Contact,
		RoleID: uint(user.RoleID),
		Parts:  userParts,
		Token:  tokenStr,
	}

	return resp, nil
}
