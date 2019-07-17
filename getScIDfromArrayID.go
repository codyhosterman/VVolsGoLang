package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Please enter a FlashArray ID:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	h1 := md5.New()
	h1.Write([]byte(text))
	firstHash := h1.Sum([]byte{})
	namespaceUUID, _ := uuid.FromBytes(firstHash)
	result_uuid := uuid.NewMD5(namespaceUUID, []byte{})
	returnValue := "vvol:" + result_uuid.String()
	returnValue = strings.Replace(returnValue, "-", "", -1)
	finalValue := returnValue[:21] + "-" + returnValue[21:]
	fmt.Println(finalValue)
}
