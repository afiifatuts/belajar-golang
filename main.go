package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Id   int
	Name string
}

var tasks []Task
var currentId int

func addTask(name string) {
	currentId++
	task := Task{
		Id:   currentId,
		Name: name,
	}
	tasks = append(tasks, task)
	fmt.Printf("Task %s berhasil ditambahkan", name)
}

func listTask() {
	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas")
		return
	}
	fmt.Println("Daftar Tugas: ")
	for _, task := range tasks {
		fmt.Printf("%d %s \n", task.Id, task.Name)
	}
}

func main() {
	fmt.Println("Selamat datang di aplikasi pengelola tugas ! ")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Daftar Tugas")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu:")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Masukan nama tugas : ")
			taskName, _ := reader.ReadString('\n')
			taskName = strings.TrimSpace(taskName)
			addTask(taskName)

		case "2":
			listTask()

		case "3":
			fmt.Println("Keluar dari aplikasi")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak valid.")
		}

	}

}
