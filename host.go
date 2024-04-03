package main

import (
	"fmt"
	"net"
)

func main() {
	IP := "localhost:"
	PORT := "8001"

	serverAddr := IP + PORT
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)

	if err != nil {
		fmt.Println("Error feo", err)
	}

	planetas := [40]string{"A", "B", "C", "D", "E", "F"}

	tesoros := [40]string{"7", "7", "7", "7", "7", "7"}

	C1, C2, C3 := "123", "1234", "12345"

	// Conexión con el servidor
	conn, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		fmt.Println("Error al conectarse al servidor", err)
	}
	conn.Close()

	fmt.Println("Conección exitosa")

	for true {
		fmt.Println("escuchando")

		Capitan, err := net.ResolveUDPAddr("udp", C1)
		Capitan2, err := net.ResolveUDPAddr("udp", C2)
		Capitan3, err := net.ResolveUDPAddr("udp", C3)

		if err != nil {
			fmt.Println("Error feo", err)
		}

		conn1, err := net.ListenUDP("udp", Capitan)
		conn2, err := net.ListenUDP("udp", Capitan2)
		conn3, err := net.ListenUDP("udp", Capitan3)

		if err != nil {
			fmt.Println("Error al conectarse al servidor", err)
		}

		fmt.Println("Conección exitosa")

		fmt.Println("Capitan 1", conn1)
		fmt.Println("Capitan 2", conn2)
		fmt.Println("Capitan 3", conn3)
		fmt.Println("Tesoro", tesoros)
		fmt.Println("Planetas", planetas)
	}

}
