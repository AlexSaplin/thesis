set -xueo pipefail

cmake -DOPENSSL_ROOT_DIR=/usr/local/opt/openssl .

make