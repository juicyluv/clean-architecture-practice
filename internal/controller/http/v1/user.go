package v1

import (
	"clean-arch/internal/entity"
	"clean-arch/internal/usecase"
	"clean-arch/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userRoutes struct {
	u usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.User, l logger.Interface) {
	r := &userRoutes{u, l}

	h := handler.Group("/users")
	{
		h.GET("/:id", r.getUser)
		h.POST("/", r.createUser)
		h.DELETE("/:id", r.deleteUser)
	}
}

func (r *userRoutes) getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user, err := r.u.GetById(c.Request.Context(), int64(id))
	if err != nil {
		errorResponse(c, err)
		return
	}

	type getUserResponse struct {
		User *entity.User `json:"user"`
	}

	c.JSON(http.StatusOK, getUserResponse{User: user})
}

func (r *userRoutes) createUser(c *gin.Context) {
	type createUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, err)
		return
	}

	userReq := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := r.u.Create(c.Request.Context(), userReq)
	if err != nil {
		errorResponse(c, err)
		return
	}

	type createUserResponse struct {
		User *entity.User `json:"user"`
	}

	c.JSON(http.StatusOK, createUserResponse{User: user})
}

func (r *userRoutes) deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	err := r.u.Delete(c.Request.Context(), int64(id))
	if err != nil {
		errorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}
