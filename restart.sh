#!/bin/bash
#
set -ev
 
docker-compose -f docker-compose.yml down
sudo rm -rf ganache/ganache-data
docker-compose -f docker-compose.yml up -d user sp tpa gana
docker logs gana