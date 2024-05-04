// Задание 4. Пакет ioutil
// Что нужно сделать:
// -Перепишите задачи 1 и 2, используя пакет ioutil.

// Что нужно сделать:
// -Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике.
// Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	// Для второго задания

	fileName := "log.txt"

	dat, err := ioutil.ReadFile(fileName)
	check(err)

	fmt.Print(string(dat))

	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	b := make([]byte, 2)

	check(err)

	fmt.Printf(string(b))

}
