go build
cp test-sg ./docker
rm test-sg
cd docker
docker-compose build
docker-compose up