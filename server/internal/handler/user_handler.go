package handler

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"fmt"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

type UserDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromModelUser(u *model.User) *UserDto {
	return &UserDto{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name"  validate:"required"`
	Email string `json:"email" validate:"required,email"`
}


func (h *UserHandler) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	if err := validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
	); err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	user := model.NewUser(req.Name, req.Email)
	if err := h.userRepo.Create(c.Request().Context(), user); err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to create user: %v", err))
	}

	return c.JSON(http.StatusCreated, FromModelUser(user))
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.userRepo.GetByID(c.Request().Context(), id)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to get user: %v", err))
	}

	return c.JSON(http.StatusOK, FromModelUser(user))
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.userRepo.Delete(c.Request().Context(), id); err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to delete user: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}
