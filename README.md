docker build -t camera-server .

docker run -dp 8000:8000 camera-server

docker ps

docker kill <container_id>

docker kill $(docker ps -q)

gcloud iam service-accounts keys create gcsfuse-sa-key.json --iam-account=gcsfuse@phucvin.iam.gserviceaccount.com

gcsfuse --key-file /home/phucvin/gcsfuse-sa-key.json  test-bucket-01-rderw /home/tom/ftp/upload
fusermount -u /home/tom/ftp/upload

sudo mount /dev/sdb /home/tom/ftp/upload
sudo umount /home/tom/ftp/upload

sudo ls /home/tom/ftp/upload

systemctl status hello.service
cat /etc/systemd/system/hello.service
sudo systemctl daemon-reload
sudo systemctl restart hello.service
sudo systemctl enable hello.service
sudo systemctl list-unit-files --type=service
sudo systemctl list-units --type=service

su tom