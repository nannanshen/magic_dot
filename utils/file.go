package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Generate_impersonated_path(target_path string) string{

	// Split the path
	targetFilePathSplit := strings.Split(target_path, "\\")

	// Get the top level directory
	topLevelDirectory := targetFilePathSplit[0]+"\\"+filepath.Join(targetFilePathSplit[1:len(targetFilePathSplit)-1]...)


	// Get the path under the top level directory
	underTopLevelDirectory := filepath.Join(targetFilePathSplit[len(targetFilePathSplit)-1])

	// Build the impersonated path
	impersonatedPath := filepath.Join(topLevelDirectory+".", underTopLevelDirectory)
	return Nt_path(impersonatedPath)
}

func Nt_path(path string) string {
	if strings.HasPrefix(path, "\\??") {
		return path
	}

	return "\\??\\" + path
}

func Dos_path(path string) string {
	return strings.Replace(path,"\\??\\", "",-1)
}

func Nt_makedirs(path string) {
	// Split the path into head and tail parts
	path_head, path_tail := filepath.Split(Dos_path(path))
	if "" == path_tail{
		return
	}else{
		Nt_makedirs(path_head)
		err := os.Mkdir(Nt_path(path_head), os.ModePerm)

		if err != nil && !os.IsExist(err) {
			panic(any(err))
		}
	}

}

func Create_magic_dot_file(filePath string, copyFrom string) {
	// Get the absolute path of the source file
	copyFromAbsPath, err := filepath.Abs(copyFrom)
	if err != nil {
		panic(any(err))
	}

	// Open the source file for reading
	srcFile, err := os.Open(copyFromAbsPath)
	if err != nil {
		panic(any(err))
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(filePath)
	if err != nil {
		panic(any(err))
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		panic(any(err))
	}
}