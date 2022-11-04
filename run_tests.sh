
docker compose -f sut/docker-compose.yml up -d

timeout 10 docker compose -f sut/docker-compose.yml logs -f | grep -q "answer  | Message from DataChannel"
xx=$?
timeout 10 docker compose -f sut/docker-compose.yml logs -f | grep -q "offer   | Message from DataChannel"
yy=$?

docker compose -f sut/docker-compose.yml stop
docker compose -f sut/docker-compose.yml rm -f

if (( $xx == 0 && $yy == 0))
then
    echo "PASS"
    exit 0
else
    echo "FAIl"
    exit 1
fi
