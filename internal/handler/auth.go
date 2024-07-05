package handler

import (
	"edusync/internal/model"
	"edusync/pkg/validator"
	"errors"

	"github.com/gin-gonic/gin"
)

// Login
// @Description Login User
// @Summary Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input model.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.Login(input)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, gin.H{
		"accessToken":  accessToken.Token,
		"refreshToken": refreshToken.Token,
	}, nil)
}

// Refresh token
// @Description Refresh Token
// @Summary Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token body model.RefreshRequest true "Refresh Token"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var (
		err   error
		input model.RefreshRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	if err = validator.ValidatePayloads(input); err != nil {
		errorResponse(c, BadRequest, err)
		return
	}

	claims, err := h.services.Authorization.ParseToken(input.Token)
	if err != nil {
		abortResponse(c, err.Error())
		return
	}

	if claims.Type != "refresh" {
		errorResponse(c, BadRequest, errors.New("token type must be refresh"))
		return
	}

	user, err := h.services.UserReader.GetById(claims.UserId)
	if err != nil {
		fromError(c, err)
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.GenerateTokens(user)
	if err != nil {
		fromError(c, err)
		return
	}

	successResponse(c, OK, gin.H{
		"accessToken":  accessToken.Token,
		"refreshToken": refreshToken.Token,
	}, nil)
}
