package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the console! (Press x to finish)")
	reader := bufio.NewReader(os.Stdin)
	finish_app := false
	for !finish_app {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = get_text(input)
		if input != "x" {
			execute_console(input)
		} else {
			fmt.Println("Finishing the app...")
			finish_app = true
		}
	}
}

func execute_console(i string) {
	recognize_command(splitter(get_text(i)))
}

func get_text(txt string) string {
	if runtime.GOOS == "windows" {
		txt = strings.TrimRight(txt, "\r\n")
	} else {
		txt = strings.TrimRight(txt, "\n")
	}
	return txt
}

func recognize_command(commands []string) {
	switch strings.ToLower(commands[0]) {
	case "mkdisk":
		exec_mkdisk(commands)
	case "exec":
		sub_command := strings.Split(commands[1], "=")
		if strings.ToLower(sub_command[0]) == "-path" {
			readFile(sub_command[1])
		} else {
			fmt.Println("Not supported command! ")
			fmt.Println("You may say -path, press -help to see the list of commands avalibles")
		}
	case "rmdisk":
		exec_mrdisk(commands)
	case "fdisk":
		exec_fdisk(commands)
	case "pause":
		fmt.Print("Executing paused \nPress any key to continue... ")
		reader := bufio.NewReader(os.Stdin)
		x, _ := reader.ReadString('\n')
		x += ""
	default:
		fmt.Println("Not supported command! ")
	}
}

type binaryFile struct {
	size int
	path string
	name string
	unit string
}

type partition struct {
	size  string
	unit  string
	path  string
	_type string
	fit   string
	name  string
}

func exec_fdisk(com []string) {
	var new_disk binaryFile
	for _, element := range com {
		spplited_command := strings.Split(element, "=")
		switch strings.ToLower(spplited_command[0]) {
		case "-size":
			i, err := strconv.Atoi(spplited_command[1])
			if i > 0 {
				new_disk.size = i
				fmt.Println("Disk size:", new_disk.size)
			} else {
				fmt.Println("Disk size must be positive")
				return
			}
			log.Fatal(err)

		case "-path":
			new_disk.path = spplited_command[1]
			fmt.Println("Disk path", new_disk.path)
		case "-name":
			if strings.HasSuffix(spplited_command[1], ".dsk") {
				new_disk.name = spplited_command[1]
				fmt.Println("Disk name", new_disk.name)
			} else {
				fmt.Println("Error! Name must have .dsk extension")
			}
		case "-unit":
			new_disk.unit = spplited_command[1]
			fmt.Println("Disk unit", new_disk.unit)
		default:
			if spplited_command[0] != "mkdisk" {
				fmt.Println(spplited_command[0], "command unknow")
			}
		}
	}
	if new_disk.unit == "" {
		new_disk.unit = "m"
		fmt.Println("You dont especify an unit size")
	}

}

func exec_mrdisk(com []string) {
	splitted_command := strings.Split(com[1], "=")
	if splitted_command[0] == "-path" {
		file_name := splitted_command[1]
		file_name += ""
	} else {
		fmt.Println(splitted_command[0], "command unknow")
	}
}

func exec_mkdisk(com []string) {
	var new_disk binaryFile
	for _, element := range com {
		spplited_command := strings.Split(element, "=")
		switch strings.ToLower(spplited_command[0]) {
		case "-size":
			i, _ := strconv.Atoi(spplited_command[1])
			if i > 0 {
				new_disk.size = i
				fmt.Println("Disk size:", new_disk.size)
			} else {
				fmt.Println("Size must be positive! ")
				return
			}
		case "-path":
			new_disk.path = spplited_command[1]
			fmt.Println("Disk path", new_disk.path)
		case "-name":
			if strings.HasSuffix(spplited_command[1], ".dsk") {
				new_disk.name = spplited_command[1]
				fmt.Println("Disk name", new_disk.name)
			} else {
				fmt.Println("Error! Name must have .dsk extension")
			}
		case "-unit":
			new_disk.unit = spplited_command[1]
			fmt.Println("Disk unit", new_disk.unit)
		default:
			if spplited_command[0] != "mkdisk" {
				fmt.Println(spplited_command[0], "command unknow")
			}
		}
	}
	if new_disk.unit == "" {
		new_disk.unit = "m"
		fmt.Println("You dont especify an unit size")
	}
}

func splitter(txt string) []string {
	commands := strings.Split(txt, " ")
	return commands
}

func readFile(file_name string) {
	f, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("Executing ", scanner.Text(), "... ")
		execute_console(strings.TrimRight(scanner.Text(), " "))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
