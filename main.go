package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	var ip, password string

	// Taking input for IP and Password
	fmt.Println(`

██████ ███████ ██      ███████ ███████ ████████ ██  █████  
██      ██      ██      ██      ██         ██    ██ ██   ██ 
██      █████   ██      █████   ███████    ██    ██ ███████ 
██      ██      ██      ██           ██    ██    ██ ██   ██ 
 ██████ ███████ ███████ ███████ ███████    ██    ██ ██   ██ 
                                                            
                                                            

 ██████╗ ███╗   ██╗███████╗     ██████╗██╗     ██╗ ██████╗██╗  ██╗    ███╗   ██╗ ██████╗ ██████╗ ███████╗
██╔═══██╗████╗  ██║██╔════╝    ██╔════╝██║     ██║██╔════╝██║ ██╔╝    ████╗  ██║██╔═══██╗██╔══██╗██╔════╝
██║   ██║██╔██╗ ██║█████╗      ██║     ██║     ██║██║     █████╔╝     ██╔██╗ ██║██║   ██║██║  ██║█████╗  
██║   ██║██║╚██╗██║██╔══╝      ██║     ██║     ██║██║     ██╔═██╗     ██║╚██╗██║██║   ██║██║  ██║██╔══╝  
╚██████╔╝██║ ╚████║███████╗    ╚██████╗███████╗██║╚██████╗██║  ██╗    ██║ ╚████║╚██████╔╝██████╔╝███████╗
 ╚═════╝ ╚═╝  ╚═══╝╚══════╝     ╚═════╝╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝    ╚═╝  ╚═══╝ ╚═════╝ ╚═════╝ ╚══════╝
         
                        
   `)

	fmt.Println("Enter VPS IP address: ")
	fmt.Scanln(&ip)
	fmt.Println("Enter VPS password: ")
	fmt.Scanln(&password)

	// Printing the input values
	fmt.Println("VPS IP address:", ip)
	fmt.Println("VPS password:", password)

	// Configuring SSH client
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Establishing SSH connection
	conn, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}

	// Create separate sessions for each command
	session1, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session1.Close()

	session3, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session3.Close()

	session4, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session4.Close()

	session5, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session5.Close()

	session6, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session6.Close()

	session7, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session7.Close()

	session8, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session8.Close()

	session9, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session9.Close()

	// Start the first command
	if err := session1.Start("sudo apt update"); err != nil {
		log.Fatalf("Failed to start command: %v", err)
	}

	// Wait for the first command to complete
	if err := session1.Wait(); err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	// Start the 3 command
	if err := session3.Start("sudo apt install curl tar wget clang pkg-config libssl-dev jq build-essential git make ncdu -y"); err != nil {
		log.Fatalf("Failed to start command: %v", err)
	}

	// Wait for the 3 command to complete
	if err := session3.Wait(); err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	// Start the 4 command
	if err := session4.Run("ver=\"1.20.2\" && cd $HOME && wget \"https://golang.org/dl/go$ver.linux-amd64.tar.gz\""); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	// Start the 5 command
	if err := session5.Start("sudo rm -rf /usr/local/go"); err != nil {
		log.Fatalf("Failed to start command: %v", err)
	}

	// Wait for the 5 command to complete
	if err := session5.Wait(); err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	// Start the 6 command
	if err := session6.Run("sudo tar -C /usr/local -xzf \"go1.20.2.linux-amd64.tar.gz\" && rm \"go1.20.2.linux-amd64.tar.gz\""); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}
	// Start the 7 command
	if err := session7.Run(`echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile; source $HOME/.bash_profile`); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	// Start the command 8
	if err := session8.Run(`
             cd $HOME;
             rm -rf celestia-node;
             git clone https://github.com/celestiaorg/celestia-node.git;
             cd celestia-node/;
			 git checkout tags/v0.9.3;
    `); err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	fmt.Println("Installation complete!")
}
