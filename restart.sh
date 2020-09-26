#!bin/bash
cd $1
git pull
docker restart a4286d5817b0
