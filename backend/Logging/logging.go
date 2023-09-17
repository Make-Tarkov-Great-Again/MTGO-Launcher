/*
Package logging provides a flexible and extensible logging framework for the MTGO-Launcher application. It supports various log types and allows you to control log output both to files and the console.
FLog v3
Usage:

	log.Init()

	// Logging examples
	log.Info("This is an informative message") // Logs to file and console
	log.Error("This is an error message", true)  // Logs to file only (silent)
	log.Debug("This is a debug message", false)    // Logs to file only by default, pass false to make it print to console
	log.AKIServerOutput("This is a AKI server output") // Logs to file only by default pass false to make it print to console

Supported Log Types:
- "warn"
- "error"
- "info"
- "debug"
- "aki"
- "mtga"
- "online"

Logging functions are provided for each log type and accept an optional "silent" parameter to control console output.

Functions:
- Info(message string, silent ...bool)
- Error(message string, silent ...bool)
- Debug(message string, silent ...bool)
- AKIServerOutput(message string, silent ...bool)
- MTGAServerOutput(message string, silent ...bool)

Example:

	log.Info("This is an informative message")         // Log to file and console
	log.Info("This message is silent", true)           // Log to file only
	log.AKIServerOutput("Server message", false)       // Log to file and console
	log.MTGAServerOutput("MTGA server message", false)   // Log to file and console
*/
package logging

import (
	"fmt"
	"log"
	storage "mtgolauncher/backend/Storage"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	logFolder     string
	logFileMap    map[string]*os.File
	logConsole    *log.Logger
	logFilePrefix string
}

var (
	logger *Logger
)
var (
	successColor  = "\033[38;2;0;175;0m"    // Green
	warningColor  = "\033[38;2;255;204;0m"  // Yellow
	logErrorColor = "\033[38;2;250;0;0m"    // Red
	infoColor     = "\033[38;2;84;175;190m" // Custom color
	resetColor    = "\033[0m"               // Reset color to default
)

// Fuck you think it does? launch you to mars? It fucking initilizes the logger
func Init() error {
	logFolder, err := storage.GetAppDataDir()
	if err != nil {
		return err
	}

	logger = &Logger{
		logFolder:  path.Join(logFolder, "logs"),
		logConsole: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	if err := logger.initFolders(); err != nil {
		return err
	}

	if err := logger.initLogFiles(); err != nil {
		return err
	}

	return nil
}

func (l *Logger) initFolders() error {
	if _, err := os.Stat(l.logFolder); os.IsNotExist(err) {
		return os.Mkdir(l.logFolder, 0755)
	}
	return nil
}

func (l *Logger) initLogFiles() error {
	l.logFileMap = make(map[string]*os.File)

	logTypes := []string{"warn", "error", "info", "debug", "server", "aki", "mtga", "online"}

	for _, logType := range logTypes {
		logTypeFolder := path.Join(l.logFolder, logType)
		if _, err := os.Stat(logTypeFolder); os.IsNotExist(err) {
			if err := os.Mkdir(logTypeFolder, 0755); err != nil {
				return err
			}
		}

		logFileName := fmt.Sprintf("log_%s_%d.log", logType, time.Now().UnixNano())
		logFilePath := path.Join(logTypeFolder, logFileName)
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		l.logFileMap[logType] = file
	}

	return nil
}

func (l *Logger) logDating() string {
	now := time.Now()
	milliseconds := now.Nanosecond() / 1e6
	return fmt.Sprintf("%02d:%02d:%02d:%03d", now.Hour(), now.Minute(), now.Second(), milliseconds)
}

func (l *Logger) log(logType string, message string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	callerFunc := runtime.FuncForPC(pc)
	callerFuncName := "unknown"
	callerPackageName := "unknown"
	silentLogging := false

	// Check if bitch is silent
	if len(args) > 0 {
		if val, ok := args[len(args)-1].(bool); ok {
			silentLogging = val
			args = args[:len(args)-1] // get that fucking "EXTRA-BOOL=TRUE(MISSING)" out of my nice log output fucker
		}
	}

	if callerFunc != nil {
		callerFuncNameFull := callerFunc.Name()
		parts := strings.Split(callerFuncNameFull, ".")

		if len(parts) > 1 {
			if strings.Contains(parts[1], "(*") {
				receiverPart := strings.Trim(parts[1], "(*")
				receiverPart = strings.Trim(receiverPart, ")")
				packageParts := strings.Split(parts[0], "/")
				callerPackageName = packageParts[len(packageParts)-1]
				callerFuncName = receiverPart + "." + parts[2]
			} else {
				packageParts := strings.Split(parts[0], "/")
				callerPackageName = packageParts[len(packageParts)-1]
				callerFuncName = parts[1] + "." + parts[2]
			}
		}
	}

	// define the colors, so like black = uhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh
	var logTypeColor string
	switch logType {
	case "error":
		logTypeColor = logErrorColor
	case "warn":
		logTypeColor = warningColor
	case "info":
		logTypeColor = infoColor
	default:
		logTypeColor = successColor
	}

	if file, exists := l.logFileMap[logType]; exists {
		prefix := fmt.Sprintf("[%s] [%s.%s]:", l.logDating(), callerPackageName, callerFuncName)
		message := fmt.Sprintf(message, args...)
		file.WriteString(prefix + " " + message + "\n")
	}

	if !silentLogging {
		logTypePrefix := fmt.Sprintf("[%s]", strings.ToUpper(logType))
		prefix := fmt.Sprintf("[%s] [%s.%s]:", l.logDating(), callerPackageName, callerFuncName)
		message := fmt.Sprintf(message, args...)
		fmt.Printf("%s%s %s%s %s\n", logTypeColor, logTypePrefix, resetColor, prefix, message)
	}
}

// Info logs an info message
func Info(message string, args ...interface{}) {
	if logger != nil {
		logger.log("info", message, args...)
	}
}

// Warn logs a warning message
func Warn(message string, args ...interface{}) {
	if logger != nil {
		logger.log("warn", message, args...)
	}
}

// Error logs an error message
func Error(message string, args ...interface{}) {
	if logger != nil {
		logger.log("error", message, args...)
	}
}

// Debug logs a debug message
func Debug(message string, args ...interface{}) {
	if logger != nil {
		logger.log("debug", message, args...)
	}
}

// AKIServerOutput logs an AKI server output message
func AKIServerOutput(message string, args ...interface{}) {
	if logger != nil {
		logger.log("aki", message, args...)
	}
}

// MTGAServerOutput logs an MTGA server output message
func MTGAServerOutput(message string, args ...interface{}) {
	if logger != nil {
		logger.log("mtga", message, args...)
	}
}

// OnlineLog logs an online log message
func OnlineLog(message string, args ...interface{}) {
	if logger != nil {
		logger.log("online", message, args...)
	}
}
