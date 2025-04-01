#!/bin/bash

if [ "$EUID" -eq 0 ]; then
  echo "Running script as root..."
else
  echo "You must be root to run this script."
  echo "Exiting..."
  exit
fi

coreRootDir="/usr/lib/eefenn-cli/"
config="${coreRootDir}/eefenn-cli.config.json"

function initConfig() {
  echo "initializing config"
  touch ${config}
  echo -e "{\n\t\n}" > ${config}
  echo "Config initialized."
}

sudo mkdir -p ${coreRootDir}

if ! [ -s ${config} ]; then
  echo "Config file is empty. Creating ${config} ..."
  initConfig
else
  echo "Config already exists."
fi
