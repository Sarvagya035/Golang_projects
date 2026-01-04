package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
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

	fmt.Println("....Golang-ssh-demo....")

	conf := sshDemoWithPrivateKey()

	//open ssh connection

	sshClient, err := ssh.Dial("tcp", sshHostName, conf)
	checkerr(err)
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	checkerr(err)

	defer session.Close()

	//exec command on remote server

	var b bytes.Buffer

	session.Stdout = &b

	err = session.Run(commandToExec)
	checkerr(err)

	log.Printf("%s: %s", commandToExec, b.String())

	//open sftp connection

	sftpClient, err := sftp.NewClient(sshClient)

	checkerr(err)
	defer sftpClient.Close()

	//create a file

	createFile, err := sftpClient.Create(fileToDownload)
	checkerr(err)
	text := "This file is created by Golang ssh,\nThis will be downloaded by Golang SSH\n"
	_, err = createFile.Write([]byte(text))
	checkerr(err)
	fmt.Println("Created File", fileToDownload)

	//upload a file

	srcfile, err := os.Open(fileToUpload)
	checkerr(err)
	defer srcfile.Close()

	dstfile, err := sftpClient.Create(fileUploadLocation)
	checkerr(err)

	defer dstfile.Close()

	_, err = io.Copy(dstfile, srcfile)
	checkerr(err)
	fmt.Println("File Uploaded Sucessfully", fileUploadLocation)

	//Download a file

	remotefile, err := sftpClient.Open(fileToDownload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open remote file: %v\n", err)
		return
	}

	defer remotefile.Close()

	localfile, err := os.Create("./download.txt")
	checkerr(err)
	defer localfile.Close()

	_, err = io.Copy(localfile, remotefile)
	checkerr(err)
	fmt.Println("File Downloaded Sucessfully")

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
