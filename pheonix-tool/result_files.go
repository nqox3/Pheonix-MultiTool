//  pheonix-tool/result_files.go
//  This file contains functions for creating and managing result directories and files.
//  It is part of the pheonix-tool application, which provides features like webhook sending and IP lookup.
//  This code is intended for educational purposes only. Do not use it for malicious activities.
//  If you have any questions, reach out to me on discord: portexploit
//  Enjoy :)

// P.S. Этот файл содержит функции для создания и управления директориями и файлами результатов.
// Он является частью приложения pheonix-tool, которое предоставляет такие функции, как отправка веб-перехватчиков и поиск IP.
// Этот код предназначен только для образовательных целей. Не используйте его для вредоносных действий.
// Если у вас есть вопросы, свяжитесь со мной в Discord: portexploit
// Приятного просмотра :)

package main

import (
	"fmt"
	"os"
	"time"
)

// Создаём директорию по типу действия
func createResultDir(action string) string {
	now := time.Now()
	timestamp := now.Format("02.01.2006_15-04") // формат: 01.10.2025_22-37

	dirName := fmt.Sprintf("result_files_%s/%s", action, timestamp)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return ""
	}
	return dirName
}
