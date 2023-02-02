package seccore

import (
	"github.com/device-security-v1/seccore/internal/app/seccore/service"
)

var _reportService *service.ReportService

func initServiceWithConfig() {
	_reportService = &service.ReportService{
		Log: nil,
	}
}
