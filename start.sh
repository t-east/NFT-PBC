#!/bin/bash
set -ev
export MSYS_NO_PATHCONV=1
 
docker-compose -f docker-compose.yml down

docker-compose -f docker-compose.yml up -d bn main #user_a user_b user_c sp tpa
docker ps -a