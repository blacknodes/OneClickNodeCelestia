# OneClickNode Introduction
OneClickNode is a Celestia tool to Deploy light with one click.
![intro](https://github.com/subhamgurjar/OneClickNode/assets/85839823/30f0a403-e6ff-4f01-99d7-0691bdf99cb3)


# VPS Automation CLI Tool

This is a command-line tool developed in Go (Golang) for automating VPS (Virtual Private Server) setup and app installation. The tool allows you to easily automate common tasks such as logging into your VPS, executing commands, and performing app installations.

## Features

- Easy VPS login: Input your VPS IP address and password to log in seamlessly.
- Automatic Command execution: Executes commands on your VPS with ease.
- App installation automation: Installs Golang and Download latest Celestia light node.
- Simple and efficient: Developed in Go for fast and efficient execution.

## Requirements

- Go (Golang) installed on your local machine.
- SSH access to your VPS.

## Getting Started

1. Clone the repository to your local machine.
2. Install the necessary Go dependencies using `go get`.
3. Build the CLI tool using `go build`.
4. Run the CLI tool from your terminal and enjoy VPS automation!
5. Enter your vps details when prompted 
   a)Enter VPS IP address: b)Enter VPS password:
6.Once Golang is installed and celestia app is downloaded you will get a success message `Installation complete!`
![done](https://github.com/subhamgurjar/OneClickNode/assets/85839823/3f095703-7ba9-48d3-823e-7e5c532b7a3d)

## Usage

You can customized the tool to suit your needs, you can execute it from the command line to automate your VPS setup and app installations. Simply run the compiled binary and let the tool handle the rest.

## Important
This tool does not generate the `make cel-key` because of the security concerns please follow the official tutorial for any help from [here](https://docs.celestia.org/nodes/celestia-node/).
Final dir should look like this
![dir](https://github.com/subhamgurjar/OneClickNode/assets/85839823/83c61ab4-89eb-4045-b437-b6fc1c88b1db)


## Contribution

Contributions are welcome! If you have any ideas, improvements, or bug fixes, feel free to submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

