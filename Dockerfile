# FROM node:20-alpine
# FROM docker:dind
FROM golang:1.20-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/web3 ./cmd/web

EXPOSE 8080
CMD ["web3"]


# RUN apk fix && apk --no-cache --update add git git-lfs gpg less openssh patch nodejs npm sudo docker aws-cli docker-credential-ecr-login

# # create user for nodejs / npm
# RUN addgroup build-user && adduser -D build-user -G build-user

# RUN mkdir -p /home/build-user/app
# COPY . /home/build-user/app
# RUN chown -R build-user:build-user /home/build-user/app

# # sudo
# RUN echo '%wheel ALL=(ALL:ALL) NOPASSWD: ALL' > /etc/sudoers.d/wheel
# RUN adduser build-user wheel

# # docker
# RUN adduser build-user docker

# USER build-user
# WORKDIR /home/build-user/app

# # docker-credential-ecr-login
# RUN mkdir -p /home/build-user/.docker && echo '{    "credsStore": "ecr-login"   }' >> /home/build-user/.docker/config.json 

# CMD ["/bin/sh"]