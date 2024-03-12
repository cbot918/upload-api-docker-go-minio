docker run  -p 9000:9000 --name minio \
 -d --restart=always \
 -e MINIO_ACCESS_KEY=minio \
 -e MINIO_SECRET_KEY=minio@123 \
 -v /usr/local/minio/data:/data \
 -v /usr/local/minio/config:/root/.minio \
  minio/minio server /data  --console-address ":9000" --address ":9090"