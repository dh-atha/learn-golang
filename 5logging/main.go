package main

import (
	"log"
	"os"
	"time"
)

// contoh implementasi log yang native library

func main() {
	// kode untuk open atau create file
	file, err := os.OpenFile("logging.info.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lmsgprefix | log.Lshortfile)
	log.SetPrefix("INFO : ")

	// for loop selamanya, untuk keluar bisa langsung
	for {
		time.Sleep(time.Second)
		log.Println("halo")
	}

}
