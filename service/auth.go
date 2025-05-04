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

func (s *AuthService) Login(req *request.LoginRequest) (*model.User, error) {
	user, err := s.repo.GetUserByUserName(req.Username)
	if err != nil {
		return nil, err
	}
	
	if user == nil {
		return nil, errors.New("tài khoản không tồn tại")
	}
	
	if user.Password != req.Password { 
		return nil, errors.New("mật khẩu không đúng")
	}
	
	return user, nil
}

func (s *AuthService) Register(req *request.RegisterRequest) (*model.User, error) {
	exists, err := s.repo.ExistsByUsernameOrEmail(req.Username, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username hoặc Email đã tồn tại")
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("lỗi hash mật khẩu")
	}

	newUser := &model.User{
		Username: req.Username,
		Email: req.Email,
		LastName: req.LastName,
		FirstName: req.FirstName,
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}