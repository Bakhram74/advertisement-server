package handler

import (
	"errors"
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
	"github.com/Bakhram74/advertisement-server.git/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type userResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	PhoneNumber       string    `json:"phone_number"`
	Role              string    `json:"role"`
	IsBanned          bool      `json:"is_banned"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:                user.ID,
		Username:          user.Username,
		PhoneNumber:       user.PhoneNumber,
		Role:              user.Role,
		IsBanned:          user.IsBanned,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
func (h *Handler) signUp(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Invalid json provided", err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Server error", err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		PhoneNumber:    req.PhoneNumber,
	}

	user, err := h.services.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("Server error", err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Invalid json provided", err))
		return
	}

	user, err := h.services.GetUser(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse("Please provide valid login details", err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("Server error", err))
		return
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("Please provide valid login details", err))
		return
	}
	accessToken, _, err := h.tokenMaker.CreateToken(
		user.ID,
		h.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Server error", err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
