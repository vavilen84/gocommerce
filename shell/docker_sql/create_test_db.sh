
docker-compose -f docker/dev/docker-compose.yaml --env-file=.env run db mysql -u root -e "CREATE DATABASE IF NOT EXISTS ${MYSQL_TEST_DATABASE}"