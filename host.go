package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

func main() {
	planets := []string{"PA", "PB", "PC", "PD", "PE", "PF"}
	register := make(map[string]int)
	for i := 0; i < len(planets); i++ {
		register[planets[i]] = rand.Intn(100)
	}

	addr := net.UDPAddr{
		Port: 8081,
		IP:   net.ParseIP("localhost"),
	}

	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		fmt.Println("Error al conectarse al servidor", err)
	}

	msg := make([]byte, 1024)
	for true {
		for key, value := range register {
			fmt.Println(key, "-", value)
		}

		n, err := conn.Read(msg)
		if err != nil {
			fmt.Println("Error al leer el mensaje", err)
		}

		dataString := string(msg[:n])
		data := strings.Split(dataString, ",")
		planet, captain, bountyStr := data[0], data[1], data[2]

		fmt.Println("Recepción de solicitud desde el Planeta %s, del capitán %s", planet, captain)

		bounty, _ := strconv.Atoi(bountyStr)
		destiny := MinPlanet(register)
		register[destiny] += bounty

		fmt.Println("Botín asignado al Planeta %s. cantidad actual: %d", destiny, register[destiny])

		conn.Close()
	}

}

func MinPlanet(register map[string]int) string {
	var minKey string
	var minValue int
	for p, b := range register {
		if b < minValue {
			minKey = p
			minValue = b
		}
	}
	return minKey
}
