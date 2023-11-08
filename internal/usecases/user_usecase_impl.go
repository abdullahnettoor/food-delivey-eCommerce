package usecases

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	hashpassword "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/hash_password"
	otphelper "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/twilio"
)

type userUcase struct {
	userRepo interfaces.IUserRepository
	dishRepo interfaces.IDishRepository
}

func NewUserUsecase(userRepo interfaces.IUserRepository, dishRepo interfaces.IDishRepository) *userUcase {
	return &userUcase{userRepo, dishRepo}
}

func (uc *userUcase) SignUp(req *req.UserSignUpReq) (*entities.User, error) {

	_, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil && err != e.ErrNotFound {
		return nil, err
	}

	hashedPwd, err := hashpassword.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	var user = entities.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPwd,
	}

	newUser, err := uc.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	err = otphelper.SendOtp(user.Phone)
	if err != nil {
		uc.userRepo.DeleteByPhone(user.Phone)
		return nil, err
	}

	return newUser, nil
}

func (uc *userUcase) VerifyOtp(phone string, req *req.UserVerifyOtpReq) error {

	if ok, err := otphelper.VerifyOtp(phone, req.Otp); err != nil || !ok {
		fmt.Println("Inside otp helper")
		return errors.New("invalid otp")
	}
	if err := uc.userRepo.Verify(phone); err != nil {
		fmt.Println("Inside verify user")
		return err
	}

	return nil
}

func (uc *userUcase) SendOtp(phone string) error {
	if err := otphelper.SendOtp(phone); err != nil {
		return err
	}
	return nil
}

func (uc *userUcase) Login(req *req.UserLoginReq) (*entities.User, error) {

	user, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := hashpassword.CompareHashedPassword(user.Password, req.Password); err != nil {
		return nil, e.ErrInvalidPassword
	}

	return user, nil
}

func (uc *userUcase) SearchDish(search string) (*[]entities.Dish, error) {
	return uc.dishRepo.Search(search)
}

func (uc *userUcase) GetDishesPage(page, limit string) (*[]entities.Dish, error) {
	p, err := strconv.ParseUint(page, 10, 0)
	if err != nil {
		return nil, err
	}
	l, err := strconv.ParseUint(limit, 10, 0)
	if err != nil {
		return nil, err
	}

	return uc.dishRepo.FindPageWise(uint(p), uint(l))
}

func (uc *userUcase) GetDish(id string) (*entities.Dish, error) {
	return uc.dishRepo.FindByID(id)
}
