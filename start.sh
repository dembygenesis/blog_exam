docker container rm test_db -f
docker container rm test_api -f
docker-compose -f  docker-compose.yml down --remove-orphans --volumes
docker-compose up --force-recreate --build