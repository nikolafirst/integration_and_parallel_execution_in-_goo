// Задание 3. Уровни доступа
// Что нужно сделать:
// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
// Рекомендация:
// Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// если удалить файл notes.txt, то при запуске компиляции файл создастся, при повторном запуске программы получится ошибка

	file, err := os.Create("notes.txt")
	if err := os.Chmod("notes.txt", 0444); err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer.WriteString("Hello World \n")
	writer.WriteString("Hello Go \n")
	writer.WriteString("Hello Golang \n")
}
