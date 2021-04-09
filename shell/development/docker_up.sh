sudo chmod 777 -R docker/db/dbdata
docker-compose -f docker/docker-compose.yaml -f docker/docker-compose.development.yaml --env-file=docker/.env.development up