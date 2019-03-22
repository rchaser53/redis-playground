# redis-playground

- chown `id -u $USER` dump.rdb
- chmod 666 dump.rdb
- docker run -u `id -u $USER` -d --name mycontainer -p 6379:6379 -v $(pwd)dump.rdb:/data/dump.rdb:rw redis:latest
