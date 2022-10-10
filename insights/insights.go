package insights

import (
	"fmt"

	"aviary/config"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type Logger struct {
	InsightsClient appinsights.TelemetryClient
}

func New(prov config.Provider) *Logger {
	client := appinsights.NewTelemetryClient(prov.Azure.InsightsKey)
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
