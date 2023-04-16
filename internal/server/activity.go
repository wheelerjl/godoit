package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/wheelerjl/godoit/internal/database"
)

func (s Server) AddActivity(ctx *gin.Context) {
	var activity database.Activity
	if err := ctx.ShouldBindJSON(&activity); err != nil {
		SetErrorResponse(ctx, http.StatusBadRequest, "failure binding body", err)
		return
	}
	newActivity, err := s.Config.Database.AddActivity(ctx.Request.Context(), activity)
	if err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to add activity", err)
		return
	}
	ctx.JSON(http.StatusCreated, newActivity)
}

func (s Server) GetActivities(ctx *gin.Context) {
	activities, err := s.Config.Database.GetActivities(ctx.Request.Context(), "")
	if err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to get activities", err)
		return
	}
	ctx.JSON(http.StatusOK, activities)
}

func (s Server) GetActivity(ctx *gin.Context) {
	activityID := ctx.Param("id")
	if _, err := uuid.Parse(activityID); err != nil {
		SetErrorResponse(ctx, http.StatusBadRequest, "invalid uuid", err)
		return
	}

	newActivity, err := s.Config.Database.GetActivity(ctx.Request.Context(), activityID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			SetErrorResponse(ctx, http.StatusNotFound, "activity not found", err)
		} else {
			SetErrorResponse(ctx, http.StatusInternalServerError, "unable to get activity", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, newActivity)
}

func (s Server) RemoveActivity(ctx *gin.Context) {
	activityID := ctx.Param("id")
	if _, err := uuid.Parse(activityID); err != nil {
		SetErrorResponse(ctx, http.StatusBadRequest, "invalid uuid", err)
		return
	}
	if err := s.Config.Database.RemoveActivity(ctx.Request.Context(), activityID); err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to remove activity", err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
