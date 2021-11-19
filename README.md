docker build -t camera-server .

docker run -dp 8000:8000 camera-server

docker ps

docker kill <container_id>