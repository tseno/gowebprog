// リスト6.2
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("ハローワールド！\n")

	// ファイルへの書き込みとファイルからの読み込み。WriteFileとReadFileを利用
	err := ioutil.WriteFile("ch06/02read_write_files/data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("ch06/02read_write_files/data1")
	fmt.Print(string(read1))

	// File構造体を利用したファイルの読み書き
	file1, _ := os.Create("ch06/02read_write_files/data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("ch06/02read_write_files/data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
