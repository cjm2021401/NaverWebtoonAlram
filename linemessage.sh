curl -X POST -H "Authorization: Bearer $1 " -F "imageThumbnail=https://image.chosun.com/sitedata/image/202006/24/2020062400801_0.png" -F "imageFullsize=https://image.chosun.com/sitedata/image/202006/24/2020062400801_0.png"  -F "message=$2"  https://notify-api.line.me/api/notify