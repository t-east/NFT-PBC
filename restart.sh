#!/bin/bash
#
set -ev
docker stop gana
docker rm gana 
docker rmi backend_gana
# docker-compose -f docker-compose.yml down
sudo rm -rf ganache/ganache-data
docker system prune
docker-compose -f docker-compose.yml up -d user sp tpa gana
docker logs gana