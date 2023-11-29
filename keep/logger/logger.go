/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package logger

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var s *zap.SugaredLogger

func init() {
	logger := createLogger()
	defer logger.Sync()
	s = logger.Sugar()
}

func createLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: nil,
	}

	return zap.Must(config.Build())
}

func Info(arg ...interface{}) {
	s.Info(arg)
}

func Error(arg ...interface{}) {
	s.Error(arg)
}

func Debug(arg ...interface{}) {
	s.Debug(arg)
}

func Pretty(arg ...interface{}) {
	b, _ := json.MarshalIndent(arg, "", "  ")
	fmt.Println(string(b))
}
