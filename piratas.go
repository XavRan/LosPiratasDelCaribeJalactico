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

		if err != nil {
			fmt.Println("Error feo", err)
		}

		conn1, err := net.Dial("udp", ":123")
		conn2, err := net.Dial("udp", ":1234")
		conn3, err := net.Dial("udp", ":12345")

		if err != nil {
			fmt.Println("Error al conectarse al servidor", err)
		}

		conn1.Write("Capitan C1 encontró botín en Planeta PA, enviando solicitud de asignación")
		fmt.Println("Conección exitosa")

		fmt.Println("Capitan 1", conn1)
		fmt.Println("Capitan 2", conn2)
		fmt.Println("Capitan 3", conn3)
		fmt.Println("Tesoro", tesoros)
		fmt.Println("Planetas", planetas)
	}

}
