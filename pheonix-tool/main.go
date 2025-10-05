// This is a main code file for a Go application that provides functionalities like sending messages to webhooks and looking up IP information.
// It includes user input handling, HTTP requests, JSON processing, and file operations.
// This Tool is inspired by nfiveqox and is intended for educational purposes only.
// Do not use it for malicious activities.
// if you have any questions, reach out to me on discord: portexploit
// Enjoy :)

// pheonix-tool/main.go

// Это основной файл кода для приложения Go, которое предоставляет такие функции, как отправка сообщений на веб-перехватчики и поиск информации об IP-адресе.
// Он включает в себя обработку пользовательского ввода, HTTP-запросы, обработку JSON и операции с файлами.
// Этот инструмент создан пользователем nfiveqox и предназначен только для образовательных целей.
// Не используйте его для вредоносных целей.
// Если у вас есть вопросы, свяжитесь со мной в Discord: portexploit
// Приятного просмотра :)

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var logo = `
 ███████████  █████   █████ ██████████    ███████    ██████   █████ █████ █████ █████
░░███░░░░░███░░███   ░░███ ░░███░░░░░█  ███░░░░░███ ░░██████ ░░███ ░░███ ░░███ ░░███ 
 ░███    ░███ ░███    ░███  ░███  █ ░  ███     ░░███ ░███░███ ░███  ░███  ░░███ ███  
 ░██████████  ░███████████  ░██████   ░███      ░███ ░███░░███░███  ░███   ░░█████   
 ░███░░░░░░   ░███░░░░░███  ░███░░█   ░███      ░███ ░███ ░░██████  ░███    ███░███  
 ░███         ░███    ░███  ░███ ░   █░░███     ███  ░███  ░░█████  ░███   ███ ░░███ 
 █████        █████   █████ ██████████ ░░░███████░   █████  ░░█████ █████ █████ █████
░░░░░        ░░░░░   ░░░░░ ░░░░░░░░░░    ░░░░░░░    ░░░░░    ░░░░░ ░░░░░ ░░░░░ ░░░░░ 
                                                                                     
`

// createResultDir is implemented in result_files.go

// Structs

type IPInfo struct {
	Country  string `json:"country"`
	City     string `json:"city"`
	Region   string `json:"regionName"`
	Timezone string `json:"timezone"`
}

type WebhookData struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

// Utility function to get user input

func input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ipLookup() {
	fmt.Println("IP Lookup")
	ip := input("Enter IP: ")

	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data IPInfo
	json.Unmarshal(body, &data)

	fmt.Println("RESULTS")
	fmt.Printf("Country:  %s\n", data.Country)
	fmt.Printf("City:     %s\n", data.City)
	fmt.Printf("Region:   %s\n", data.Region)
	fmt.Printf("Timezone: %s\n", data.Timezone)

	// Save results to file
	dir := createResultDir("ip")
	if dir != "" {
		filePath := filepath.Join(dir, "result.txt")
		content := fmt.Sprintf("IP: %s\nCountry: %s\nCity: %s\nRegion: %s\nTimezone: %s\n",
			ip, data.Country, data.City, data.Region, data.Timezone)
		os.WriteFile(filePath, []byte(content), 0644)
		fmt.Println("Saved results to:", filePath)
	}

	input("\nPress enter to return...")
}

// createResultDir is implemented in result_files.go

func webhookSender() {
	fmt.Println("WEBHOOK SENDER")
	url := input("Webhook URL: ")
	message := input("Message: ")
	name := input("Webhook Name: ")

	spam := strings.ToLower(input("Send multiple messages? (y/n): "))

	data := WebhookData{Content: message, Username: name}
	payload, _ := json.Marshal(data)

	// Create results directory and log file
	dir := createResultDir("webhook")
	logFile := filepath.Join(dir, "log.txt")

	if spam == "y" {
		countStr := input("How many times to send: ")
		count, _ := strconv.Atoi(countStr)

		delayStr := input("Delay between messages (seconds): ")
		delay, _ := strconv.ParseFloat(delayStr, 64)

		for i := 1; i <= count; i++ {
			_, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
			entry := fmt.Sprintf("[%d/%d] Sent: %s\n", i, count, message)
			if err != nil {
				entry = fmt.Sprintf("[%d/%d] ERROR sending\n", i, count)
			}
			f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			f.WriteString(entry)
			f.Close()

			fmt.Print(entry)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
		fmt.Println("Spam completed! Logs saved to:", logFile)
	} else {
		_, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		entry := fmt.Sprintf("[Single] Message: %s\n", message)
		if err != nil {
			entry = "[Single] ERROR sending message\n"
		}
		os.WriteFile(logFile, []byte(entry), 0644)
		fmt.Println(entry)
		fmt.Println("Saved log to:", logFile)
		input("Press enter to return...")
	}
}

func telegramBotSender() {
	fmt.Println("TELEGRAM BOT SENDER")
	token := input("Bot Token: ")
	chatID := input("Chat ID: ")
	message := input("Message: ")

	spam := strings.ToLower(input("Send multiple messages? (y/n): "))

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	// Create results directory and log file
	dir := createResultDir("telegram")
	logFile := filepath.Join(dir, "log.txt")

	if spam == "y" {
		countStr := input("How many times to send: ")
		count, _ := strconv.Atoi(countStr)

		delayStr := input("Delay between messages (seconds): ")
		delay, _ := strconv.ParseFloat(delayStr, 64)

		for i := 1; i <= count; i++ {
			payload := map[string]string{
				"chat_id": chatID,
				"text":    message,
			}
			payloadBytes, _ := json.Marshal(payload)
			resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
			entry := fmt.Sprintf("[%d/%d] Sent: %s\n", i, count, message)
			if err != nil {
				entry = fmt.Sprintf("[%d/%d] ERROR sending\n", i, count)
			} else {
				resp.Body.Close()
			}
			f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			f.WriteString(entry)
			f.Close()

			fmt.Print(entry)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
		fmt.Println("Spam completed! Logs saved to:", logFile)
	} else {
		payload := map[string]string{
			"chat_id": chatID,
			"text":    message,
		}
		payloadBytes, _ := json.Marshal(payload)
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
		entry := fmt.Sprintf("[Single] Message: %s\n", message)
		if err != nil {
			entry = "[Single] ERROR sending message\n"
		} else {
			resp.Body.Close()
		}
		os.WriteFile(logFile, []byte(entry), 0644)
		fmt.Println(entry)
		fmt.Println("Saved log to:", logFile)
		input("Press enter to return...")
	}
}

func main() {
	for {
		fmt.Println(logo)
		fmt.Println("[1] Webhook Sender")
		fmt.Println("[2] IP Info")
		fmt.Println("[3] Telegram Bot Sender")
		fmt.Println("[4] ID Lookup [NOT WORKING]")
		fmt.Println("[5] Server Raid [NOT WORKING]")

		choice := input("Option: ")

		switch choice {
		case "1":
			webhookSender()
		case "2":
			ipLookup()
		case "3":
			telegramBotSender()
		default:
			fmt.Println("Invalid or not implemented.")
			time.Sleep(1 * time.Second)
		}
	}
}
