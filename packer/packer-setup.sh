mkdir /etc/docker
echo '{"storage-driver": "vfs"}' > /etc/docker/daemon.json

apt update
apt install --no-install-recommends docker.io nginx vim -y

mkdir -m=500 /nginx
touch /nginx/server-cert.pem /nginx/server-key.pem /nginx/ca-cert.pem