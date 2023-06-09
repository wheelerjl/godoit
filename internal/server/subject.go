package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/wheelerjl/godoit/internal/database"
)

func (s Server) AddSubject(ctx *gin.Context) {
	var subject database.Subject
	if err := ctx.ShouldBindJSON(&subject); err != nil {
		SetErrorResponse(ctx, http.StatusBadRequest, "failure binding body", err)
		return
	}
	newSubject, err := s.Config.Database.AddSubject(ctx.Request.Context(), subject)
	if err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to add subject", err)
		return
	}
	ctx.JSON(http.StatusCreated, newSubject)
}

func (s Server) GetSubjects(ctx *gin.Context) {
	subjects, err := s.Config.Database.GetSubjects(ctx.Request.Context())
	if err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to get subjects", err)
		return
	}
	ctx.JSON(http.StatusOK, subjects)
}

func (s Server) GetSubject(ctx *gin.Context) {
	subjectID := ctx.Param("id")
	if _, err := uuid.Parse(subjectID); err != nil {
		SetErrorResponse(ctx, http.StatusBadRequest, "invalid uuid", err)
		return
	}

	newSubject, err := s.Config.Database.GetSubject(ctx.Request.Context(), subjectID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			SetErrorResponse(ctx, http.StatusNotFound, "subject not found", err)
		} else {
			SetErrorResponse(ctx, http.StatusInternalServerError, "unable to get subject", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, newSubject)
}

func (s Server) RemoveSubject(ctx *gin.Context) {
	subjectID := ctx.Param("id")
	if _, err := uuid.Parse(subjectID); err != nil {
		SetErrorResponse(ctx, http.StatusBadRequest, "invalid uuid", err)
		return
	}
	if err := s.Config.Database.RemoveSubject(ctx.Request.Context(), subjectID); err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to remove subject", err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
