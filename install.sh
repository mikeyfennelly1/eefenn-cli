#!/bin/zsh

# Check if the script is run as root
if [[ $(id -u) -ne 0 ]]; then
    echo "You must be root to run this script. Use 'sudo'.." >&2
    exit 1
fi

eefennCLIBinaryPath="/usr/bin/eefenn-cli"

go build -o ${eefennCLIBinaryPath}
sudo chown 0:0 ${eefennCLIBinaryPath}
sudo chmod +s ${eefennCLIBinaryPath}
alias ef='${eefennCLIBinaryPath}'
source ~/.zshrc
