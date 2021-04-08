sudo chmod 777 -R docker/dev/db/dbdata
docker-compose -f docker/dev/docker-compose.yaml --env-file=.env up --build
