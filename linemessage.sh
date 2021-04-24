curl -X POST -H "Authorization: Bearer $1 " -F "imageThumbnail=$3" -F "imageFullsize=$3"  -F "message=$2"  https://notify-api.line.me/api/notify
