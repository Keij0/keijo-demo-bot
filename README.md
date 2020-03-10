A simple bot written using telebot v2 golang library. It uses docker golang:alpine image as its runtime

To build the image, execute:

docker build -t bot .

To execute the application, execute:

docker run -it --rm --name bot-run bot
