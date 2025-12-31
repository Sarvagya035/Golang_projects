package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println("Error happened", err)
		os.Exit(1)
	}
}

const (
	main_folder_name    = "parent"
	child_folder_name   = "child"
	empty_folder_name   = "empty"
	child_file_name     = "child.txt"
	parent_file_name    = "parent.txt"
	parent_file_content = "Hello this is the parent folder of the program"
	child_file_content  = "Hello this is the content of child folder"
)

func createFolder(targetpath string, foldername string) {

	fileinfo, err := os.Stat(targetpath)
	checkerr(err)

	if !fileinfo.IsDir() {
		fmt.Println("Present path is not a valid directory")
		os.Exit(1)
	}

	folderpath := filepath.Join(targetpath, foldername)
	err = os.Mkdir(folderpath, 0755)

	if err != nil {
		fmt.Println("Error creating folder", err)
	}

}

func createFile(targetpath string, filename string, filedata string) {

	fileinfo, err := os.Stat(targetpath)
	checkerr(err)

	if !fileinfo.IsDir() {
		fmt.Println("Present path is not a valid directory")
		os.Exit(1)
	}

	file_path := filepath.Join(targetpath, filename)

	err = os.WriteFile(file_path, []byte(filedata), 0644)

	if err != nil {
		fmt.Println("Error creating the file", err)
	}
}

func main() {

	fmt.Println("Welcome to file-folder-zip program")

	pwd, err := os.Getwd()
	checkerr(err)

	createFolder(pwd, main_folder_name)

	parent_path := filepath.Join(pwd, main_folder_name)

	createFolder(parent_path, empty_folder_name)
	createFolder(parent_path, child_folder_name)

	createFile(parent_path, parent_file_name, parent_file_content)

	child_path := filepath.Join(parent_path, child_folder_name)

	createFile(child_path, child_file_name, child_file_content)
}
