package service

import (
	"be-fiber/model"
	"be-fiber/repository"
	"be-fiber/router/request"
	"be-fiber/security"
	"errors"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(req *request.LoginRequest) (*model.User, string, error) {
	user, err := s.repo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, "", err
	}
	
	if user == nil {
		return nil, "", errors.New("người dùng không tồn tại")
	}

	if err := security.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, "", errors.New("lỗi xác thực mật khẩu")
	}

	refreshToken, err := security.GenerateToken(user.ID.Hex())
	if err != nil {
		return nil, "", errors.New("lỗi tạo token")
	}
	
	return user, refreshToken, nil
}

func (s *AuthService) Register(req *request.RegisterRequest) (*model.User, string, error) {
	exists, err := s.repo.ExistsByUsernameOrEmail(req.Username, req.Email)
	if err != nil {
		return nil, "", err
	}
	if exists {
		return nil, "", errors.New("username hoặc Email đã tồn tại")
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, "", errors.New("lỗi hash mật khẩu")
	}

	newUser := &model.User{
		Username: req.Username,
		Email: req.Email,
		LastName: req.LastName,
		FirstName: req.FirstName,
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return nil, "", errors.New("lỗi tạo người dùng")
	}

	refreshToken, err := security.GenerateToken(newUser.ID.Hex())
	if err != nil {
		return nil, "", errors.New("lỗi tạo token")
	}

	return newUser, refreshToken, nil
}