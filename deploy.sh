#!/bin/sh
set -e

echo "Pulling The App from git"
git pull

echo "Enabling swag"
cd /usr/local/go/bin/
PATH=$(go env GOPATH)/bin:$PATH
swag -v
cd /var/www/go-api-eshop

echo "Generate documentation from place2connect-api"

chmod +x swagg.sh
./swagg.sh
echo "Api documentation successfully generated"

echo "Building and deploying the entire application"
docker-compose up

echo "Our app is successfully deployed"

echo "Reloading our application For Caddy"
# sudo service caddy reload

# OTHER CADDY COMMAND
# sudo service caddy restart
# sudo service caddy stop
# sudo service caddy status
# systemctl status caddy.service
# journalctl -xeu caddy.service
