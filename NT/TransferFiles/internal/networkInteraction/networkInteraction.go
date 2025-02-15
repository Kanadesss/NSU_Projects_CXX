package networkInteraction

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	SuccessMsg              = "success"
	ErrMsg                  = "error"
	EndSymbol               = '\a'
	TransferFileAcceptedMsg = "9876"
	TransferFileDeclineMsg  = "ABCD"
	bufferSize              = 1048576    // 1024 ^ 2
	bufferSizeBig           = 1073741824 // 1024 ^ 3
	// bufferSizeBig           = 1048576
)

func GetMessage(conn net.Conn) (msg string) {
	msg, err := bufio.NewReader(conn).ReadString(EndSymbol)
	if err != nil {
		fmt.Println("Error reading fileName: ", err.Error())
	}
	endSymbolIndex := strings.LastIndex(msg, string(EndSymbol))
	if endSymbolIndex == -1 {
		endSymbolIndex = len(msg) - 1
	}
	msg = msg[:endSymbolIndex]
	return
}

func SendMessage(msg string, conn net.Conn) {
	_, err := conn.Write([]byte(msg + string(EndSymbol)))
	if err != nil {
		println("Error sendResponse: ", err.Error())
	}
}

func GetBuffer(file *os.File) []byte {
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Error get file info: ", err.Error())
		return make([]byte, bufferSize)
	}
	return make([]byte, SetBufferSize(int64(fi.Size())))
}

func GetBufferBySize(fileSize int64) []byte {
	return make([]byte, SetBufferSize(fileSize))
}

func SetBufferSize(fileSize int64) int64 {
	if fileSize/bufferSize >= 50 {
		fmt.Println("choose big buffer")
		return bufferSizeBig
	}
	fmt.Println("choose std buffer")
	return bufferSize
}
