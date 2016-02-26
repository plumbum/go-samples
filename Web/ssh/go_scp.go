// Based on https://gist.github.com/jedy/3357393
// https://blogs.oracle.com/janp/entry/how_the_scp_protocol_works
package main

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
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
		content := time.Now().String() + "\n"
		fmt.Fprintln(w, "D0755", 0, "testdir") // mkdir
		fmt.Fprintln(w, "C0644", len(content), "testfile1")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00") // transfer end with \x00
		fmt.Fprintln(w, "C0644", len(content), "testfile2")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00")
	}()
	log.Println("Create files")
	if err := session.Run("/usr/bin/scp -tr ./"); err != nil {
		log.Fatalln("Failed to run: ", err)
	}
}
