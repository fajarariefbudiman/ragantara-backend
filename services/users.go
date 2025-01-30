package services

import (
	"errors"
	"fmt"
	"net/http"
	"real-time-application/database"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id                int       `json:"id"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Age               int       `json:"age"`
	Address           string    `json:"address"`
	Email             string    `json:"email" gorm:"unique"`
	Password          string    `json:"password"`
	Role              string    `json:"role" gorm:"type:enum('customer','owner','admin','player');default:'customer'"`
	Phone             string    `json:"phone"`
	SkillLevel        string    `json:"skill_level" gorm:"type:enum('beginner','intermediate','advanced');default:'beginner'"`
	Position          string    `json:"position" gorm:"type:enum('goalkeeper','defender','midfielder','forward');default:'midfielder'"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		var user User
		err := database.DB.Where("email = ?", fl.Field().String()).First(&user).Error
		if err == nil {
			return false
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	})
}

func AuthUser(email, password string) (*User, error) {
	var user = &User{}

	validate := validator.New()
	if err := validate.Var(email, "required,email"); err != nil {
		return nil, fmt.Errorf("email wajib diisi & dengan format yang benar")
	}
	if err := validate.Var(password, "required"); err != nil {
		return nil, fmt.Errorf("password wajib diisi")
	}

	if err := database.DB.Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("email tidak diketahui")
		}
		return nil, fmt.Errorf("terjadi kesalahan saat mengambil data: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("password salah")
	}

	return user, nil
}

func Register(firstname, lastname, email, password string) (Response, error) {
	var res Response
	validate := validator.New()
	if err := validate.Var(firstname, "required"); err != nil {
		return res, fmt.Errorf("nama depan wajib diisi")
	}
	if err := validate.Var(email, "required,email"); err != nil {
		return res, fmt.Errorf("email wajib diisi & dengan format yang benar")
	}
	if err := validate.Var(password, "required,min=8"); err != nil {
		return res, fmt.Errorf("password wajib diisi")
	}
	var existingUser User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return res, fmt.Errorf("email sudah terdaftar")
	}
	var user = &User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
	}
	result := database.DB.Create(user)
	if result.Error != nil {
		return res, fmt.Errorf("Disini")
	}
	res.Status = http.StatusOK
	res.Message = "success register"
	res.Data = user
	// fmt.Println(user)
	return res, nil
}
