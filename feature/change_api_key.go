package feature

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// ChangeKey cập nhật giá trị API_KEY trong file .env
func ChangeKey(apiKey string) {
    envFilePath := os.Getenv("WEA_ENV") // Đường dẫn đến file .env

    // Đọc file .env
    file, err := os.Open(envFilePath)
    if err != nil {
        fmt.Printf("Không thể mở file .env: %v\n", err)
        return
    }
    defer file.Close()

    // Đọc nội dung file
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Cập nhật giá trị API_KEY
        if strings.HasPrefix(line, "API_KEY=") {
            line = "API_KEY=" + apiKey
        }
        lines = append(lines, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Lỗi khi đọc file: %v\n", err)
        return
    }

    // Ghi lại nội dung vào file .env
    outFile, err := os.Create(envFilePath)
    if err != nil {
        fmt.Printf("Không thể ghi file .env: %v\n", err)
        return
    }
    defer outFile.Close()

    writer := bufio.NewWriter(outFile)
    for _, line := range lines {
        writer.WriteString(line + "\n")
    }
    writer.Flush()

}
