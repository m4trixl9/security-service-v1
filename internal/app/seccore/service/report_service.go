package service

import (
	"fmt"
	"io"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ReportService struct {
	Log *zap.Logger
}

func (r *ReportService) HandleReportV1(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.WithMessage(err, "http body"))
		return
	}

	fmt.Println(string(body))
}
