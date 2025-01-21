package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var lists []string
	var menu int

	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lists = append(lists, scanner.Text())
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("===== Todo List Kelompok 10 =====")
		if len(lists) < 1 {
			fmt.Println("List kamu masih kosong")
		} else {
			for i, list := range lists {
				fmt.Printf("%d. %s\n", i+1, list)
			}
		}
		fmt.Println("===== Todo List Kelompok 10 =====")
		fmt.Println("1. Tambah List")
		fmt.Println("2. Edit List")
		fmt.Println("3. Hapus List")
		fmt.Println("4. Exit & Save")
		fmt.Println("===== Pilih Menu Diatas Dengan Input Angka ===== ")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:

			fmt.Print("Masukan Todo: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			lists = append(lists, input)

			clearCmd()

		case 2:
			var listIndex int
			fmt.Print("Pilih nomor list yang ingin di edit: ")
			fmt.Scanln(&listIndex)

			if listIndex > len(lists) {
				clearCmd()
				fmt.Println("!!! List tidak ditemukan !!!")
				continue
			}

			var list = lists[listIndex-1]
			fmt.Printf("Masukan Todo untuk mengubah %s: ", list)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			lists[listIndex-1] = input

			clearCmd()

		case 3:
			var listIndex int
			fmt.Print("Pilih nomor list yang ingin di hapus: ")
			fmt.Scanln(&listIndex)

			if listIndex > len(lists) {
				clearCmd()
				fmt.Println("!!! List tidak ditemukan !!!")
				continue
			}

			lists = append(lists[:listIndex], lists[listIndex+1:]...)

			clearCmd()
		case 4:
			file.Close()

			file, _ := os.OpenFile("data.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			defer file.Close()

			writer := bufio.NewWriter(file)

			for _, list := range lists {
				_, err = writer.WriteString(list + "\n")
			}
			err = writer.Flush()
			fmt.Println("Terima kasih telah menggunakan program Todo List Kelompok 10")
			os.Exit(0)
		}
	}

}

func clearCmd() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
