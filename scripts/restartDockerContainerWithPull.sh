#!/bin/bash

echo "running"
sleep 3s

# -q only prints the image id of that tag
GOLANG_API_IMAGE_ID=$(docker images uditnair90/api-padhai-golang -q)
GOLANG_API_IMAGE=uditnair90/api-padhai-golang:latest
GOLANG_API_CONTAINER_NAME=uditnair90_api-padhai-golang


echo "Restarting container with latest image..."
docker stop $GOLANG_API_CONTAINER_NAME || true
docker rm $GOLANG_API_CONTAINER_NAME || true
docker run -d --pull=always --quiet --name $GOLANG_API_CONTAINER_NAME --env PORT=10000 --publish 10000:10000 $GOLANG_API_IMAGE
# docker run --detach --name $GOLANG_API_CONTAINER_NAME --env PORT=10000 --publish 10000:10000 $GOLANG_API_IMAGE
# docker run --detach --pull=always --quiet --name uditnair90_api-padhai-golang --env PORT=10000 --publish 10000:10000 uditnair90/api-padhai-golang:latest