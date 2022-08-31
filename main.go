package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ListDirs(inf Inf) {
	files, err := ioutil.ReadDir(inf.path) //считывание файлов из директории
	if err != nil {
		log.Fatal(err)
	}

	var tabs string //переменная для хранения нужного количества табуляций

	for _, file := range files {
		tabs = ""                        //сброс переменной в пустое состояние
		for l := 1; l <= inf.levl; l++ { //определение нужного количества табуляций
			if inf.levl > 1 {
				tabs += "\t"
			}

		}

		//Дописывание DIR для папок FILE для файлов для более легкой навигации и понимания
		if file.IsDir() {
			fmt.Println(tabs + "DIR: " + file.Name())
		} else {
			fmt.Println(tabs + "FILE: " + file.Name())
		}

		//Проверка и вхождение в найденую папку
		if file.IsDir() {
			new_path := inf.path + file.Name() + "\\"   //создание нового пути с найденой папкой
			ListDirs(Inf{new_path, inf.levl + 1, true}) //Рекурсивный вызов с передачей новых параметров
		}

		//Проверка и уменьшение уровня табуляции при виходе из папки
		if inf.levl > 1 && inf.IsInDir == false {
			inf.levl -= 1
		}
	}
	inf.IsInDir = false //После выхода из папки меняем
}

type Inf struct { //структура для хранения параметров нужных функции ListDirs
	path    string //переменная пути исходной папки
	levl    int    //уровень для проставления нужного количкства табуляций при выводе фалов
	IsInDir bool   //переменная для понимания находится ли найденый файд в папке или нет
}

func main() {
	var path string
	fmt.Print("Enter path:")
	fmt.Scan(&path)

	inf := Inf{path, 1, false}

	ListDirs(inf)

}
