package greetingworker

import (
	"context"
)

// HelloWorldEventLogger logs hello world events.
type HelloWorldEventLogger struct {
	logger Logger
}

// NewHelloWorldEventLogger returns a new HelloWorldEventLogger.
func NewHelloWorldEventLogger(logger Logger) *HelloWorldEventLogger {
	return &HelloWorldEventLogger{
		logger: logger,
	}
}

// SaidHello logs a SaidHello event.
func (e *HelloWorldEventLogger) SaidHello(ctx context.Context, event SaidHello) error {
	e.logger.WithFields(LogFields{"message": event.Message}).Info("said hello")

	return nil
}
