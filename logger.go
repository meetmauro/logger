
package logger


import (
   "fmt"
   "io"
   "log"
   "os"
   "strings"
   "runtime"
   "path/filepath"
)

const (
   LogError   = 0
   LogWarning = 1
   LogInfo    = 2
   LogDebug   = 3
   LogTrace   = 4
)

// mapping from string to log levels
func LogLevelMapper(ll string) int {
   ll = strings.ToUpper(ll)
   switch ll {
      case "ERROR": return LogError
      case "WARNING": return LogWarning
      case "INFO": return LogInfo
      case "DEBUG": return LogDebug
      case "TRACE": return LogTrace
   }
   return LogDebug
}

var (
   _logger *log.Logger
   _level int
)


//initialize the logger at the specified log level and file
func Initialize(level int, logFile string) {
   var out io.Writer
   var err error

   if logFile == "" {
      out = os.Stdout
   } else {
      out, err = os.OpenFile(logFile, os.O_APPEND, 0666)
      if err != nil {
         out = os.Stdout
         fmt.Println("Failed to open log file: " + err.Error())
      }
   }

   _level = level
   flags := log.Ldate | log.Lmicroseconds
   _logger = log.New(out, "", flags)
   log.SetOutput(out)
   log.SetFlags(flags)
}

func GetLog() *log.Logger {
   return _logger
}

func Debug(format string, a ...interface{}) {
   if _level >= LogDebug {
      _logger.Output(2, "[DEBUG] " + fmt.Sprintf(format, a...))
   }
}

func Info(format string, a ...interface{}) {
   if _level >= LogInfo {
      _logger.Output(2, "[INFO] " + fmt.Sprintf(format, a...))
   }
}

func Warning(format string, a ...interface{}) {
   if _level >= LogWarning {
      _logger.Output(2, "[WARNING] " + fmt.Sprintf(format, a...))
   }
}

func Error(format string, a ...interface{}) {
   if _level >= LogError {
      pc, file, line, ok := runtime.Caller(1)
      msg := ""
      if ok {
         fn := runtime.FuncForPC(pc).Name()
         _, f := filepath.Split(file)
         msg = fmt.Sprintf("%s %d %s ", f, line, fn)
      }
      _logger.Output(2, "[ERROR] " + msg + fmt.Sprintf(format, a...))
   }
}

func Fatal(format string, a ...interface{}) {
   if _level >= LogError {
      pc, file, line, ok := runtime.Caller(1)
      msg := ""
      if ok {
         fn := runtime.FuncForPC(pc).Name()
         _, f := filepath.Split(file)
         msg = fmt.Sprintf("%s %d %s ", f, line, fn)
      }      
      _logger.Fatalln("[ERROR] " + msg + fmt.Sprintf(format, a...))
   }
}

func Trace(format string, a ...interface{}) {
   if _level >= LogTrace {
      _logger.Output(2, "[TRACE] " + fmt.Sprintf(format, a...))
   }
}


