mkdir /etc/docker
echo '{"storage-driver": "vfs"}' > /etc/docker/daemon.json

apt update
apt install --no-install-recommends docker.io nginx -y

mkdir -m=500 /nginx
chown www-data:www-data /nginx