package password

import (
	"agnos/backend/pkgs/log"
	"agnos/backend/pkgs/response"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
)

type passwordHandler struct {
	mpaValidator MPA
	logger       *log.Queries
}

func NewPasswordHandler(mpaValidator MPA, logger *log.Queries) *passwordHandler {
	return &passwordHandler{
		mpaValidator: mpaValidator,
		logger:       logger,
	}
}

func (h *passwordHandler) RecommendMinimumPasswordAction(ctx *gin.Context) {
	payload := RecommendMinimumPasswordActionRequest{}
	if e := ctx.ShouldBindBodyWithJSON(&payload); e != nil {
		response.HandleResponse(ctx, http.StatusBadRequest, e)
		return
	}
	if e := validator.New().Struct(&payload); e != nil {
		response.HandleResponse(ctx, http.StatusBadRequest, e)
		return
	}

	response := RecommendMinimumPasswordActionResponse{
		NumOfSteps: h.mpaValidator.GetMinimumActionToValid(payload.InitPassword),
	}

	rawPayload, _ := json.Marshal(payload)
	rawResponse, _ := json.Marshal(RecommendMinimumPasswordActionResponse{
		NumOfSteps: h.mpaValidator.GetMinimumActionToValid(payload.InitPassword),
	})
	h.logger.CreateLog(context.Background(), log.CreateLogParams{
		Endpoint: "/api/strong_password_steps",
		Ip: pgtype.Text{
			String: ctx.ClientIP(),
			Valid:  true,
		},
		StatusCode: pgtype.Int4{
			Int32: int32(http.StatusOK),
			Valid: true,
		},
		Request: string(rawPayload),
		Response: pgtype.Text{
			String: string(rawResponse),
			Valid:  true,
		},
		CreateAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	})

	ctx.JSON(http.StatusOK, response)
}
