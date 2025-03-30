#!/bin/zsh

go build
sudo mv ./eefenn-cli /usr/bin/eefenn-cli
echo "alias ef='/usr/bin/eefenn-cli'" >> ~/.zshrc
source ~/.zshrc
