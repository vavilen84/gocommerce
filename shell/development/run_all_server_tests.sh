docker-compose -f docker/docker-compose.yaml -f docker/docker-compose.development.yaml --env-file=docker/.env.development run server go test ./... -p 1 -count=1 -v
