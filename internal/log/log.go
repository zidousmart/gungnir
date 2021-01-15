package log

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/zidousmart/gungnir/internal/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type GLogger struct {
	logger *zap.Logger

	AppName string
}

// 日志按天分割
func New(appName, path string) (*GLogger, error) {
	fileDir, fileName := filepath.Split(path)
	if fileDir == "" || fileName == "" {
		return nil, errors.New("file path error")
	}

	err := file.CreateDirByPath(fileDir)
	if err != nil {
		return nil, err
	}

	fileExt := filepath.Ext(fileName)
	fileNameOnly := strings.TrimSuffix(fileName, fileExt)

	filePath := fileDir + "/" + fileNameOnly + "_%Y%m%d.log"
	hook, err := rotatelogs.New(
		filePath,
		// rotatelogs.WithLinkName("/path/to/access_log"),
		// rotatelogs.WithMaxAge(24*time.Hour),
		// rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		return nil, err
	}
	w := zapcore.AddSync(hook)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.MessageKey = logKeyMsg
	encoderCfg.LevelKey = logKeyLevel
	encoderCfg.TimeKey = logKeyTimestamp
	encoderCfg.NameKey = ""
	encoderCfg.CallerKey = ""
	encoderCfg.StacktraceKey = ""
	encoderCfg.LineEnding = ""
	encoderCfg.EncodeLevel = levelEncoder
	encoderCfg.EncodeTime = timeEncoder
	encoderCfg.EncodeDuration = nil
	encoderCfg.EncodeCaller = nil
	encoderCfg.EncodeName = nil

	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			w,
			zapcore.DebugLevel,
		),
		// zap.AddCaller(),
		// zap.AddCallerSkip(1),
	)
	defer logger.Sync()

	return &GLogger{
		logger:  logger,
		AppName: appName,
	}, nil
}

func (log *GLogger) Logger() *zap.Logger {
	return log.logger
}

func (log *GLogger) Debug(c *gin.Context, a ...interface{}) {
	msg := fmt.Sprint(a...)
	list := log.listAppField(c, zapcore.DebugLevel)
	log.logger.Debug(msg, list...)
}

func (log *GLogger) Info(c *gin.Context, a ...interface{}) {
	msg := fmt.Sprint(a)
	list := log.listAppField(c, zapcore.InfoLevel)
	log.logger.Info(msg, list...)
}

func (log *GLogger) Warn(c *gin.Context, a interface{}) {
	msg := fmt.Sprint(a)
	list := log.listAppField(c, zapcore.WarnLevel)
	log.logger.Warn(msg, list...)
}

func (log *GLogger) Error(c *gin.Context, a interface{}) {
	msg := fmt.Sprint(a)
	list := log.listAppField(c, zapcore.ErrorLevel)
	log.logger.Error(msg, list...)
}

func (log *GLogger) Critical(c *gin.Context, a interface{}) {
	msg := fmt.Sprint(a)
	list := log.listAppField(c, zapcore.FatalLevel)
	log.logger.Fatal(msg, list...)
}

func (log *GLogger) listAppField(c *gin.Context, level zapcore.Level) []zap.Field {
	var list []zap.Field
	list = append(list, zap.String(logKeyName, log.AppName))
	list = append(list, zap.String(logKeyVersion, logValueVersion))

	xFile, xLine := getFieldFileLine(callerSkip)
	list = append(list, zap.String(logKeyFile, xFile))
	list = append(list, zap.String(logKeyLine, xLine))

	return list
}

func (log *GLogger) NgxInfo(c *gin.Context, response string, duration int64) {
	list := log.listNgxField(c, response, duration)
	log.logger.Info("", list...)
}

func (log *GLogger) NgxError(c *gin.Context, response string, duration int64) {
	list := log.listNgxField(c, response, duration)
	log.logger.Error("", list...)
}

func (log *GLogger) listNgxField(c *gin.Context, response string, duration int64) []zap.Field {
	var list []zap.Field
	list = append(list, zap.String(logKeyName, log.AppName))
	list = append(list, zap.String(logKeyVersion, logValueVersion))

	xFile, xLine := getFieldFileLine(1)
	list = append(list, zap.String(logKeyFile, xFile))
	list = append(list, zap.String(logKeyLine, xLine))

	reqBody, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
	request := string(reqBody)
	list = append(list, zap.String(logKeyRequest, request))
	list = append(list, zap.String(logKeyResponse, response))

	list = append(list, zap.Int64(logKeyDuration, duration))

	return list
}

func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString("[" + l.CapitalString() + "]")
	// enc.AppendString(l.String())
	enc.AppendString(getLevelName(l))
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
	enc.AppendInt64(t.Unix())
}

func durationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + d.String() + "]")
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func nameEncoder(name string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(name)
}

func getLevelName(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return "debug"
	case zapcore.InfoLevel:
		return "info"
	case zapcore.WarnLevel:
		return "warning"
	case zapcore.ErrorLevel:
		return "error"
	case zapcore.DPanicLevel:
		return "panic"
	case zapcore.PanicLevel:
		return "panic"
	case zapcore.FatalLevel:
		return "fatal"
	}

	return "unknown"
}

func getFieldFileLine(skip int) (string, string) {
	caller := zapcore.NewEntryCaller(runtime.Caller(skip))
	fileLine := strings.Split(caller.TrimmedPath(), ":")
	if len(fileLine) == 2 {
		return fileLine[0], fileLine[1]
	} else if len(fileLine) > 0 {
		return fileLine[0], ""
	}

	return "", ""
}
