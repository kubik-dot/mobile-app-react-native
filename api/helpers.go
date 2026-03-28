package mobileappreactnative

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math"
	"net/smtp"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than zero")
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func SHA256(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func SendMail(subject string, body string, recipients []string, smtpServer string, smtpPort int, smtpUser string, smtpPassword string) error {
	from := "your-email@example.com"
	mensagem := "From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"Content-Transfer-Encoding: 8bit\r\n" +
		"MIME-Version: 1.0\r\n" +
		"\r\n" + body

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, strings.Split(smtpServer, ":")[0])
	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, from, recipients, []byte(mensagem))
	if err != nil {
		return err
	}
	return nil
}

func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

func GetFilesInDirectory(path string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}

func GetFileSize(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func ConvertStringsToBytes(strs []string) []byte {
	var b bytes.Buffer
	for _, s := range strs {
		b.WriteString(s)
		b.WriteByte('\n')
	}
	return b.Bytes()
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsBetween(value float64, min float64, max float64) bool {
	return math.Min(min, max) <= value && value <= math.Max(min, max)
}

func GetOneHourAgo() time.Time {
	return time.Now().Add(-time.Hour).Truncate(time.Second)
}