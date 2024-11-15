package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	domain "github.com/linn-phyo/go_gin_clean_architecture/src/domain"
	services "github.com/linn-phyo/go_gin_clean_architecture/src/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID          string    `copier:"must"`
	UserName    string    `copier:"must"`
	Email       string    `copier:"must"`
	Password    string    `copier:"must"`
	CreatedDate time.Time `copier:"must"`
	UpdatedDate time.Time `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// FindAll godoc
// @summary Get all users
// @description Get all users
// @tags users
// @security ApiKeyAuth
// @id FindAll
// @produce json
// @Router /api/users [get]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &users)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id := paramsId
	// id, err := strconv.Atoi(paramsId)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "cannot parse id",
	// 	})
	// 	return
	// }

	user, err := cr.userUseCase.FindByID(c.Request.Context(), id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) Save(c *gin.Context) {
	var user domain.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	log.Println("USER DATA >>", user)
	user, err := cr.userUseCase.Save(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id := paramsId

	// id, err := strconv.Atoi(paramsId)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Cannot parse id",
	// 	})
	// 	return
	// }

	ctx := c.Request.Context()
	user, err := cr.userUseCase.FindByID(ctx, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if user == (domain.Users{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found!",
		})
		return
	}

	cr.userUseCase.Delete(ctx, user)

	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}
