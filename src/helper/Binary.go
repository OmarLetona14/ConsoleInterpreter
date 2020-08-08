package helper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type payload struct {
	One   float32
	Two   float64
	Three uint32
}

func ReadingFile(file_name string) {
	file, err := os.Open(file_name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	m := payload{}
	for i := 0; i < 10; i++ {
		data := ReadNextBytes(file, 16)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &m)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}
		fmt.Println(m)
	}
}

func ReadNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func WriteFile() {
	file, err := os.Create("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		s := &payload{
			r.Float32(),
			r.Float64(),
			r.Uint32(),
		}
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.BigEndian, s)
		//b :=bin_buf.Bytes()
		//l := len(b)
		//fmt.Println(l)
		WriteNextBytes(file, bin_buf.Bytes())
	}
}

func WriteNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
