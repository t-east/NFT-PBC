#!/bin/bash
#
set -ev
 
docker-compose -f docker-compose.yml down
sudo rm -rf ganache/ganache-data