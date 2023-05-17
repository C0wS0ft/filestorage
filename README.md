# This is FileStorage - sample distributed file storage

It consists of two parts - backend and volume servers

backend is single

volumes are multiple

## backend (cmd/back)

### POST /v1/upload - upload file
format: multipart/form-data

### POST /v1/register - add new volume
format: { "url" : "volume url" }

### GET /v1/download - download file
example: http://host/v1/download?id=syslog

## volume (cmd/vol)

## POST /v1/upload - upload file
format: multipart/form-data

### GET /v1/download - download file
example: http://volume_host/v1/download?id=syslog

# Running
backend is always on 8001

volumes use optional ports

build binaries using make

start back

$./back

start volumes

./vol :8002

./vol :8003

./vol :8004

Feel free to add volumes on the go.

Use provided bash scripts to

_01_register.sh

_02_upload.sh  

_03_download.sh
