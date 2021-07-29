package mylog

import (
	"io"
	"log"
	"os"
)

var (
	logFlags    = log.Ldate | log.Lmicroseconds | log.Llongfile
	infoPrefix  = "[INFO]\t"
	errorPrefix = "ERROR\t"
	debugPrefix = "DEBUG\t"
	tracePrefix = "TRACE\t"
)

type MyLog struct {
	info  *log.Logger
	debug *log.Logger
	trace *log.Logger
	error *log.Logger
}

func (l MyLog) Info(v ...interface{}) {
	l.info.Println(v)
}
func (l MyLog) Error(v ...interface{}) {
	l.error.Println(v)
}
func (l MyLog) Debug(v ...interface{}) {
	l.debug.Println(v)
}
func (l MyLog) Trace(v ...interface{}) {
	l.trace.Println(v)
}

func New() *MyLog {
	return &MyLog{
		trace: log.New(os.Stdout, tracePrefix, logFlags),
		info:  log.New(os.Stdout, infoPrefix, logFlags),
		error: log.New(os.Stdout, errorPrefix, logFlags),
		debug: log.New(os.Stdout, debugPrefix, logFlags),
	}
}

func init() {
	log.SetPrefix("INFO---")
	log.SetFlags(logFlags) // 1 | 4 | 8 = 13
}

func logDemo() {
	l := log.New(os.Stdout, infoPrefix, logFlags)
	l.Println("println log")
	//l.Fatalln("Fatalln log")
	//l.Panicln("Panicln log") // unreachable code
}

func mylogDemo() {
	mylog := New()
	mylog.Info("dsfdsfmmmmd")
	mylog.Error("dsfdsfd")
	mylog.Debug("dsfdsfd")
	mylog.Trace("dsfdsfd")
}

// FileLog 文件log
func FileLog() *log.Logger {
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	//Trace := log.New(ioutil.Discard, "[TRACE]: ", logFlags) // ioutil.Discard 忽略该等级日志
	Info := log.New(io.MultiWriter(file, os.Stdout), "[INFO]: ", logFlags)
	//Info.Println("Info log...")
	//Trace.Println("Trace log...")
	return Info
}

func main() {
	//	const (
	//		a = 1 << iota
	//		b
	//	)
	//	fmt.Println(a, b)

	//mylogDemo()
	//logDemo()
	FileLog()
}
