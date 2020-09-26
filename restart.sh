echo "$1, $2"
cd $1
git pull
docker restart $2
