package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Đường dẫn tới tệp log
	logFilePath := "app.log"

	// Mở tệp log
	logFile, err := os.Open(logFilePath)
	if err != nil {
		fmt.Println("Lỗi khi mở tệp log:", err)
		return
	}
	defer logFile.Close()

	// Lấy kích thước tệp log ban đầu
	initialSize, err := logFile.Seek(0, os.SEEK_END)
	if err != nil {
		fmt.Println("Lỗi khi lấy kích thước tệp log ban đầu:", err)
		return
	}

	for {
		// Lấy kích thước hiện tại của tệp log
		currentSize, err := logFile.Seek(0, os.SEEK_END)
		if err != nil {
			fmt.Println("Lỗi khi lấy kích thước hiện tại của tệp log:", err)
			return
		}

		// Nếu kích thước hiện tại lớn hơn kích thước ban đầu
		if currentSize > initialSize {
			// Đặt con trỏ đọc tại vị trí ban đầu của phần log mới
			_, err := logFile.Seek(initialSize, os.SEEK_SET)
			if err != nil {
				fmt.Println("Lỗi khi đặt con trỏ đọc tại vị trí ban đầu của phần log mới:", err)
				return
			}

			// Đọc và in ra các dòng log mới
			scanner := bufio.NewScanner(logFile)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			// Cập nhật kích thước ban đầu
			initialSize = currentSize
		}

		// Đợi một khoảng thời gian trước khi kiểm tra lại
		time.Sleep(1 * time.Second)
	}
}
