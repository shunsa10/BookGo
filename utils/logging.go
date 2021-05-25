package utils

//logの設定
import (
	"io"
	"log"
	"os"
)

func LoggingSettings(LogFile string)  { //読み書き　　　作成　　　　　追記　　　　　パーミッション
	//logファイルの設定
	logfile, err := os. OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	//logの書き込み先を標準出力先とlogfileにしている
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //logのフォーマットの指定
	log.SetOutput(multiLogFile)//logの出力先
	
}