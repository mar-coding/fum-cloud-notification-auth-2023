# fum-cloud-notification-auth-2023
[Project description]

## Local Development


To run this project locally, you will need to install the following dependencies:

1. Postgresql:

    If you want to change connection's configs, you must change it on below files:

    - ```envs/dev.env```
    - ```docker-compose.yml```

2. Golang 1.18.1:

   ```
   sudo apt update
   sudo apt upgrade
   sudo apt install golang-go
   ```

   You can test the installation by running `go version` in your terminal.

## Usage
Run this project locally
  ```
   go build cmd/main.go
   ./main
  ```