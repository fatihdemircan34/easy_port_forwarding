package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var (
	targetIP   string
	targetPort string
	listenPort string
)

func init() {
	flag.StringVar(&targetIP, "targetIP", "192.168.1.10", "Yönlendirme yapılacak hedef IP")
	flag.StringVar(&targetPort, "targetPort", "8080", "Yönlendirme yapılacak hedef port")
	flag.StringVar(&listenPort, "listenPort", ":9090", "DMZ'de dinlenecek port")
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("Port dinlenirken hata: %v", err)
	}
	defer listener.Close()

	log.Printf("Port %s dinleniyor...", listenPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Bağlantı kabul edilirken hata: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(sourceConn net.Conn) {
	defer sourceConn.Close()

	targetConn, err := net.Dial("tcp", targetIP+":"+targetPort)
	if err != nil {
		log.Printf("Hedefe bağlanırken hata: %v", err)
		return
	}
	defer targetConn.Close()

	go func() {
		_, err := io.Copy(targetConn, sourceConn)
		if err != nil {
			log.Printf("Veri kopyalanırken hata: %v", err)
		}
	}()

	_, err = io.Copy(sourceConn, targetConn)
	if err != nil {
		log.Printf("Veri kopyalanırken hata: %v", err)
	}
}
