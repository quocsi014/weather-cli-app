#!/bin/bash

# 1. Build Go binary
echo "Building Go binary..."
go build -o wea main.go

# 2. Move binary to /usr/local/bin
echo "Moving binary to /usr/local/bin..."
sudo mv wea /usr/local/bin

# 3. Create .wea directory in user's home directory
echo "Creating .wea directory in home directory..."
mkdir -p $HOME/.wea

# 4. Create .env file in .wea directory and write API_KEY
echo "Creating .env file..."
echo "API_KEY=" > $HOME/.wea/.env

# 5. Export WEA_ENV globally
export WEA_ENV=$HOME/.wea/.env
if ! grep -q "export WEA_ENV=$HOME/.wea/.env" $HOME/.bashrc; then
    echo "export WEA_ENV=$HOME/.wea/.env" >> $HOME/.bashrc
fi

# Reload shell configuration

source $HOME/.bashrc

echo "Setup completed. You can now use the wea_app command."
