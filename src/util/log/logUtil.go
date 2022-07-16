package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	errorLogger   *log.Logger
	infoLogger    *log.Logger
)

func init() {
	warnFile, warn_err := os.OpenFile("../log/warn.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	errorFile, error_err := os.OpenFile("../log/error.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	infoFile, info_err := os.OpenFile("../log/info.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if warn_err != nil {
		fmt.Println("warn日志文件初始化 失败!")
	}
	if error_err != nil {
		fmt.Println("error日志文件初始化 失败!")
	}
	if info_err != nil {
		fmt.Println("日志文件初始化 失败!")
	}

	// 各级别的日志 不仅会输出到对应的 日志文件，同时也会打印在控制台
	warnWriter := io.MultiWriter(os.Stdout, warnFile)
	errorWriter := io.MultiWriter(os.Stdout, errorFile)
	infoWriter := io.MultiWriter(os.Stdout, infoFile)

	warningLogger = log.New(warnWriter, "WARNING: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	errorLogger = log.New(errorWriter, "ERROR: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	infoLogger = log.New(infoWriter, "INFO: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func PrintWarnLog(content string) {
	warningLogger.Println(" -- " + content + " -- ")
}

func PrintErrorLog(content string) {
	errorLogger.Println(" -- " + content + " -- ")
}

func PrintInfoLog(content string) {
	infoLogger.Println(" -- " + content + " -- ")
}
