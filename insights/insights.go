package insights

import (
	"fmt"
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

const (
	envName = "APPINSIGHTS_INSTRUMENTATIONKEY"
)

type Logger struct {
	InsightsClient appinsights.TelemetryClient
}

func New() *Logger {
	client := appinsights.NewTelemetryClient(os.Getenv(envName))
	return &Logger{
		InsightsClient: client,
	}
}

// Error logs an error to application insights.
func (l *Logger) Error(message string, err error) {
	wrappedErr := fmt.Sprintf(message+" %w", err)
	exception := appinsights.NewExceptionTelemetry(wrappedErr)
	exception.SeverityLevel = appinsights.Error
	l.InsightsClient.Track(exception)
}
