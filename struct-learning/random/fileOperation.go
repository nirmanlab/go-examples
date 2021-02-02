package random

import "fmt"

type FileReader struct {
	FileString string
}
type StreamReader struct {
	StreamString string
}
type NetwrorkReader struct {
	NetwrorkString string
}

func (file *FileReader) Read() {
	fmt.Println("File Reader struct", file.FileString)
}
func (stream *StreamReader) Read() {
	fmt.Println("Stream reader struct", stream.StreamString)
}
func (network *NetwrorkReader) Read() {
	fmt.Println("Netwrok Reader Struct", network.NetwrorkString)
}
