// Задание 2. Интерфейс io.Reader
// Что нужно сделать:
// -Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil.
// Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
// Рекомендация:
// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл или файл не найден", err)
		return
	}
	defer file.Close()

	f, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Printf("File name: %v\n", f.Name())
	fmt.Printf("Size: %d\n", f.Size())
	fmt.Printf("Mode: %v\n", f.Mode())
}
