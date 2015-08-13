package main

import (
	"log"
	"path/filepath"
	"io/ioutil"
	"golang.org/x/crypto/ssh"
	"fmt"
	"bufio"
)

func main() {

	keyFileName := "/home/ivan/.ssh/demo_id_rsa"
	log.Println("Load private key")
	keyFilePath, err := filepath.Abs(keyFileName)
	if err != nil {
		log.Fatalln(err)
	}
	privateKey, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		log.Fatalln("Can't open key file", err)
	}
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalln("Error key", err)
	}
	clientConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	log.Println("Dial to SSH server")
	client, err := ssh.Dial("tcp", "172.17.0.11:22", clientConfig)
	if err != nil {
		log.Fatalln("Failed to dial: ", err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatalln("Failed to create session: ", err)
	}
	defer session.Close()
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		r, _ := session.StdoutPipe()

		fmt.Fprintln(w, "/bin/ls -al")
		fmt.Fprintln(w, "/bin/ps aux")
		fmt.Fprintln(w, "/bin/pwd")
		fmt.Fprintln(w, "exit")

		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

	}()
	log.Println("Create files")
	if err := session.Run("/bin/bash"); err != nil {
		log.Fatalln("Failed to run: ", err)
	}
}
