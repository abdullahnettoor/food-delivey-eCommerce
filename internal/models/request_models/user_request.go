package req

type UserSignUpReq struct {
	FirstName       string `json:"firstName" validate:"required,gte=3"`
	LastName        string `json:"lastName"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,e164"`
	Password        string `json:"password" validate:"gte=3"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
}

type UserVerifyOtpReq struct {
	Otp string `json:"otp" validate:"required,number"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=3"`
}

type NewAddressReq struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"houseName" validate:"required"`
	Street    string `json:"street" validate:"required"`
	District  string `json:"district" validate:"required"`
	State     string `json:"state" validate:"required"`
	PinCode   string `json:"pinCode" validate:"required,len=6"`
	Phone     string `json:"phone" validate:"required,e164"`
}

type NewOrderReq struct {
	PaymentMethod string `json:"paymentMethod"`
	AddressID     string `json:"addressId"`
}
