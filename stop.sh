#!/bin/bash
#
set -ev
docker stop gana
docker rm gana 
docker rmi -f backend_gana
# docker-compose -f docker-compose.yml down
sudo rm -rf ganache/ganache-data