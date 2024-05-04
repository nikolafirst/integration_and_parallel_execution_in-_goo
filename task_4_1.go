// Задание 4. Пакет ioutil
// Что нужно сделать:
// -Перепишите задачи 1 и 2, используя пакет ioutil.

// Задание 1. Работа с файлами
// Что нужно сделать:
// -Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:
// 2020-02-10 15:00:00 продам гараж.
// -При вводе слова exit программа завершает работу.

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"time"
// )

// func main() {
// 	filePath := "log.txt"
// 	writer, err := os.Create(filePath)
// 	if err != nil {
// 		fmt.Println("Ошибка создания файла:", err)
// 		return
// 	}
// 	defer writer.Close()

// 	scanner := bufio.NewScanner(os.Stdin)
// 	lineNumber := 1
// 	for {
// 		fmt.Print("Введите сообщение (для выхода введите 'exit'): ")
// 		scanner.Scan()
// 		text := scanner.Text()
// 		if text == "exit" {
// 			fmt.Println("Программа завершена.")
// 			break
// 		}

// 		currentTime := time.Now().Format("2006-01-02 15:04:05")
// 		message := fmt.Sprintf("%v %s\n", currentTime, text)

// 		err := ioutil.WriteFile(filePath, []byte(fmt.Sprintf("%v %s\n", lineNumber, message)), 0644)
// 		if err != nil {
// 			fmt.Println("Ошибка записи в файл:", err)
// 			break
// 		}
// 		// fmt.Printf("Запись #%v в файл: %s", lineNumber, message)
// 		lineNumber++
// 	}
// }