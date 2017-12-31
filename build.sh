name=${1:-redis-server}
docker run -itp 6379:6379 --name $name -d redis