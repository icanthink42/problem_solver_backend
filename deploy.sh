#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <ec2-ip-address>"
    exit 1
fi

EC2_IP=$1
EC2_USER="ec2-user"

echo "Building Go binary for ARM64..."
GOOS=linux GOARCH=arm64 go build -o problem_solver

if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Creating directories and setting permissions on EC2..."
ssh ${EC2_USER}@${EC2_IP} "sudo rm -rf ~/problem_solver ~/questions ~/start.sh && mkdir -p ~/questions && sudo chown -R ec2-user:ec2-user ~/"

echo "Copying binary and questions to EC2..."
scp problem_solver ${EC2_USER}@${EC2_IP}:~/
scp -r questions/* ${EC2_USER}@${EC2_IP}:~/questions/

echo "Copying start script to EC2..."
scp start.sh ${EC2_USER}@${EC2_IP}:~/start.sh

echo "Setting permissions..."
ssh ${EC2_USER}@${EC2_IP} "chmod +x ~/problem_solver ~/start.sh"

echo "Cleaning up local binary..."
rm problem_solver

echo "Deployment complete!"
echo "To start the server, SSH into the EC2 instance and run: ./start.sh"