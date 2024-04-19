package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadOldschool(file_name string) (in_bytes []byte, in_string string) {
	file, err := os.Open(file_name)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// file statistics
	stats, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// prepare byte slice for reading
	in_bytes = make([]byte, stats.Size())
	_, err = file.Read(in_bytes) // read file into in_bytes
	if err != nil {
		panic(err)
	}

	in_string = string(in_bytes)
	return
}

func ReadShortWay(file_name string) (in_bytes []byte, in_string string) {
	in_bytes, err := ioutil.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	in_string = string(in_bytes)

	return
}

func ListDir(dirname string) []os.FileInfo {
	dir, err := os.Open(dirname)

	if err != nil {
		panic(err)
	}

	defer dir.Close()
	list, err := dir.Readdir(-1) // return all the files in this directory

	if err != nil {
		panic(err)
	}
	return list
}

func WalkDir(dirname string) {
	filepath.Walk(dirname, func(file_path string, file_info os.FileInfo, err error) error {
		fmt.Println(file_path, ": \n\t", file_info.Name(), file_info.Size(), "B", file_info.Mode())
		return nil
	})
	//The function you pass to Walk is called for every file
	// and folder in the root folder.
}

func main() {
	defer func() {
		err_msg := recover()

		if err_msg != nil {
			fmt.Println("FUCKED: ", err_msg)
		}
	}()
	in_bytes, in_string := ReadOldschool("./main.go")
	in_bytes2, in_string2 := ReadShortWay("./files.go")

	fmt.Println("\n-----------------FILE 1 IN BYTES---------------\n", in_bytes, "\n------------FILE 2 IN BYTES-----------\n", in_bytes2)
	file, err := os.Create("final.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(in_string + "\n\n" + in_string2)

	file_list := ListDir(".")
	fmt.Println("\n------------ FILES IN THE CWD --------------------\n")
	for _, f := range file_list {
		fmt.Println(f.Name())
	}

	fmt.Println("\n------------ STARTING A WALK HERE --------------------\n")
	WalkDir("../.")
}
