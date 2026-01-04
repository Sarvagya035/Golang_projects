package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func checkerr(err error) {

	if err != nil {
		fmt.Println("Error Happened", err)
		os.Exit(1)
	}
}

var (
	sshUserName        = "vagrant"
	sshPassword        = "vagrant"
	sshKeyPath         = "C:/Users/DELL/Desktop/Vagrant_vms/.vagrant/machines/default/virtualbox/private_key"
	sshHostName        = "127.0.0.1:2222"
	commandToExec      = "ls -la /home/vagrant"
	fileToUpload       = "./upload.txt"
	fileUploadLocation = "/home/vagrant/upload.txt"
	fileToDownload     = "/home/vagrant/download.txt"
)

func main() {

	fmt.Println("....Golanf=g-ssh-demo....")

	conf := sshDemoWithPrivateKey()

	//open ssh connection

	sshClient, err := ssh.Dial("tcp", sshHostName, conf)
	checkerr(err)

	session, err := sshClient.NewSession()
	checkerr(err)

	defer session.Close()

	//exec command on remote server

	var b bytes.Buffer

	session.Stdout = &b

	err = session.Run(commandToExec)
	checkerr(err)

	log.Printf("%s: %s", commandToExec, b.String())

}

func sshDemoWithPrivateKey() *ssh.ClientConfig {

	keyByte, err := os.ReadFile(sshKeyPath)
	checkerr(err)

	key, err := ssh.ParsePrivateKey(keyByte)
	checkerr(err)

	conf := &ssh.ClientConfig{
		User: sshUserName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf
}
