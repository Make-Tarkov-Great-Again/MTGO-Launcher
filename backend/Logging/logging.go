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

var logsFolder string
var logsFolderAppDir string

func init() {
	var err error
	logsFolder, err = storage.GetAppDataDir()
	logsFolderAppDir = path.Join(logsFolder, "/logs")
	logsFolder = logsFolderAppDir
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

const loggingRoutesFile = ""

var logFileMap = make(map[string]string)
var flog = make(map[string]func(message string, args ...interface{}))

func GetFlog() map[string]func(message string, args ...interface{}) {
	return flog
}

func logDating() string {
	now := time.Now()
	milliseconds := now.Nanosecond() / 1e6
	return fmt.Sprintf("%02d:%02d:%02d:%03d", now.Hour(), now.Minute(), now.Second(), milliseconds)
}

func updateLoggingRoutesFile() {
	logTypeFolders := make(map[string]string)

	logTypes := []string{"warn", "error", "info", "debug", "server", "aki", "mtga", "online"}
	for _, logType := range logTypes {
		logTypeFolder := path.Join(logsFolder, logType)
		logTypeFolders[logType] = logTypeFolder

		if _, err := os.Stat(logTypeFolder); os.IsNotExist(err) {
			os.Mkdir(logTypeFolder, 0755)
		}
	}

	for logType, logTypeFolder := range logTypeFolders {
		if _, exists := logFileMap[logType]; !exists {
			logFileName := fmt.Sprintf("log_%s_%d.log", logType, time.Now().UnixNano())
			logFilePath := path.Join(logTypeFolder, logFileName)
			logFileMap[logType] = logFilePath
		}
	}

	var routes []string
	for logType, logFileName := range logFileMap {
		routes = append(routes, fmt.Sprintf("%s:%s", logType, logFileName))
	}
	routesContent := strings.Join(routes, "\n")
	err := os.WriteFile(loggingRoutesFile, []byte(routesContent), 0644)
	if err != nil {
		log.Printf("Failed to update logging routes file: %v\n", err)
	}
}

func readLoggingRoutesFile() {
	content, err := os.ReadFile(loggingRoutesFile)
	if err != nil {
		//log.Printf("Failed to read logging routes file: %v\n", err) I dont wanna hear "OH OH OH WE CANT READ THAT" x30 and im too lazy to change it. fuck you. i dont even think i used logging routes since fucking verson 1.
		return
	}

	routes := strings.Split(string(content), "\n")
	for _, route := range routes {
		parts := strings.Split(route, ":")
		if len(parts) == 2 {
			logFileMap[parts[0]] = parts[1]
		}
	}
}

func createLogFunction(logType string) func(message string, args ...interface{}) {
	readLoggingRoutesFile()
	return func(message string, args ...interface{}) {
		flogg(logType, message, false, args...)
	}
}

func flogg(logType, format string, silent bool, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	callerFunc := runtime.FuncForPC(pc)
	callerFuncName := "unknown"
	callerPackageName := "unknown"
	if callerFunc != nil {
		callerFuncNameFull := callerFunc.Name()
		parts := strings.Split(callerFuncNameFull, ".")

		if len(parts) > 1 {
			//If its a method receiver. (If it has "(*" ) extract the method receivers name, and put it infront of func name so i dont get shit like "Launcher.StartServer" WHICH ONE?
			if strings.Contains(parts[1], "(*") {
				receiverPart := strings.Trim(parts[1], "(*")
				receiverPart = strings.Trim(receiverPart, ")")
				packageParts := strings.Split(parts[0], "/")
				callerPackageName = packageParts[len(packageParts)-1]
				callerFuncName = receiverPart + "." + parts[2]
			} else {
				// Non-method receiver
				packageParts := strings.Split(parts[0], "/")
				callerPackageName = packageParts[len(packageParts)-1]
				callerFuncName = parts[1] + "." + parts[2]
			}
		}
	}

	if _, exists := logFileMap[logType]; !exists {
		logTypeFolder := path.Join(logsFolder, logType)
		if _, err := os.Stat(logTypeFolder); os.IsNotExist(err) {
			os.Mkdir(logTypeFolder, 0755)
		}

		logFileName := fmt.Sprintf("log_%s_%d.log", logType, time.Now().UnixNano())
		logFilePath := path.Join(logTypeFolder, logFileName)
		logFileMap[logType] = logFilePath
	}

	logFile, err := os.OpenFile(logFileMap[logType], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v\n", err)
		return
	}
	defer logFile.Close()

	prefix := fmt.Sprintf("[%s] [%s.%s]:", logDating(), callerPackageName, callerFuncName)

	var argsSlice []interface{}
	for _, arg := range args {
		argsSlice = append(argsSlice, arg)
	}

	message := fmt.Sprintf(format, argsSlice...)
	logFile.WriteString(prefix + " " + message + "\n")

	if !silent {
		fmt.Println(prefix, message)
	}
}
func LogInit() {
	var err error
	logsFolder, err = storage.GetAppDataDir()
	logsFolderAppDir = path.Join(logsFolder, "/logs")
	logsFolder = logsFolderAppDir
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create the logs folder if it doesn't exist
	if _, err := os.Stat(logsFolder); os.IsNotExist(err) {
		os.Mkdir(logsFolder, 0755)
	}

	// Create logType folders if they don't exist
	logTypes := []string{"warn", "error", "info", "debug", "server"}
	for _, logType := range logTypes {
		logTypeFolder := path.Join(logsFolder, logType)
		if _, err := os.Stat(logTypeFolder); os.IsNotExist(err) {
			os.Mkdir(logTypeFolder, 0755)
		}
	}

	// Update the logging routes file
	updateLoggingRoutesFile()
	readLoggingRoutesFile()

	flog = map[string]func(message string, args ...interface{}){
		"warn":   createLogFunction("warn"),
		"error":  createLogFunction("error"),
		"info":   createLogFunction("info"),
		"debug":  createLogFunction("debug"),
		"aki":    createLogFunction("aki"),
		"mtga":   createLogFunction("mtga"),
		"online": createLogFunction("online"),
	}

	flogg("info", "Hello world", true)
}

func Info(message string, args ...interface{}) {
	silent := false
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok {
			silent = val
			args = args[1:]
		}
	}
	flogg("info", message, silent, args...)
}

func Error(message string, args ...interface{}) {
	silent := false
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok {
			silent = val
			args = args[1:]
		}
	}
	flogg("error", message, silent, args...)
}

func Debug(message string, args ...interface{}) {
	silent := true //Def true becaue debug
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok {
			silent = val
			args = args[1:]
		}
	}
	flogg("debug", message, silent, args...)
}

// Flog: File Logging -> AKI
func AKIServerOutput(message string, args ...interface{}) {
	silent := true //Def true because its server output
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok {
			silent = val
			args = args[1:]
		}
	}
	flogg("aki", message, silent, args...)
}

// Flog: File Logging -> mtga
func MTGAServerOutput(message string, args ...interface{}) {
	silent := true
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok {
			silent = val
			args = args[1:]
		}
	}
	flogg("mtga", message, silent, args...)
}
