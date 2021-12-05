docker build -t camera-server .

docker run -dp 8000:8000 camera-server

docker ps

docker kill <container_id>

docker kill $(docker ps -q)

gcloud iam service-accounts keys create gcsfuse-sa-key.json --iam-account=gcsfuse@phucvin.iam.gserviceaccount.com

sudo gcsfuse --key-file /home/phucvin/gcsfuse-sa-key.json  test-bucket-01-rderw /home/tom/ftp/upload
sudo fusermount -u /home/tom/ftp/upload

sudo mount /dev/sdb /home/tom/ftp/upload
sudo umount /home/tom/ftp/upload

sudo ls /home/tom/ftp/upload