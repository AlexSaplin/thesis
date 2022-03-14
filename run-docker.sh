set -xeuo pipefail

 docker-compose -f docker-compose-services.yml kill
 docker-compose -f docker-compose-services.yml down

./run-db.sh

docker-compose -f docker-compose-services.yml build --parallel
docker-compose -f docker-compose-services.yml up -d

docker-compose -f docker-compose-services.yml logs -f || true

docker-compose -f docker-compose-services.yml kill
docker-compose -f docker-compose-services.yml down
