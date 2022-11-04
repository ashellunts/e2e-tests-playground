docker compose -f sut/docker-compose.yml up -d

timeout 10 docker compose -f sut/docker-compose.yml logs -f | grep -q "answer  | Message from DataChannel"
timeout 10 docker compose -f sut/docker-compose.yml logs -f | grep -q "offer   | Message from DataChannel"

docker compose -f sut/docker-compose.yml stop
docker compose -f sut/docker-compose.yml rm -f
