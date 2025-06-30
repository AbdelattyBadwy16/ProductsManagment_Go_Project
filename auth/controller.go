package auth

import (
	"project/Database/models"
	"project/utils"

	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c iris.Context) {
	var req RegisterRequest
	if err := c.ReadJSON(&req); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	err := ValidateCreateRequest(req)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	account, err := CreateAccountService(user)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := BrandTransformer(account)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)

}

func Login(c iris.Context) {
	var req LoginRequest
	if err := c.ReadJSON(&req); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	user, err := GetUser(req.Email)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{"error": "Invalid Password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{"error": "Servr Error"})
		return
	}
	response := AccountResponse{
		Name:  user.Name,
		ID:    user.ID,
		Token: token,
	}
	c.StatusCode(iris.StatusAccepted)
	c.JSON(response)
}
