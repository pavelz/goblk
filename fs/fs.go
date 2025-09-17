ackage fs

import (
	"io"
	"os"
	"fmt"
)

type File interface{
	io.Reader
	io.Writer
}

type FsOps interface{
	Open(path string) (File, error)
}

type Fs struct{}

func (Fs) Open(path string) FsFile{
	f,err := os.Create(path)
	if (err != nil){
		fmt.Println(err)
		os.Exit(-1)
	}
	return FsFile{File: f}
}

type FsFile struct{
	File *os.File
}

func Hey(){

	fmt.Println("HEY")
}

func (f FsFile) read(){
	buffer := make([]byte, func() int {s,_ := f.File.Stat(); return int(s.Size())}())

	f.File.Read(buffer)
}

func (f FsFile) Write(buffer []byte) (int, error){
	return f.File.Write(buffer)
}

func (f FsFile) Close(){
	f.File.Close()
}

