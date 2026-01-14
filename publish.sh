cd /data/www/snai.cms.api;
chmod +x Snai.CMS.Api;

# 如果 imagev.txt 不存在，则创建并初始化为 0
if [ ! -f imagev.txt ]; then
    echo 0 > imagev.txt
fi

# 获取当前版本号
imagev0=$(cat imagev.txt) 
imagev=$(($imagev0 + 1)); 

# 构建新镜像
docker build -t Snai.CMS.Api_"$imagev" .; 

# 写入新版本号
destdir=imagev.txt 
if [ -f "$destdir" ] 
then 
    echo "$imagev" > "$destdir" 
fi 

# 停掉旧容器
docker stop Snai.CMS.Api_"$imagev0"; 

# 启动新容器
docker run -it -d --name=Snai.CMS.Api_"$imagev" \
    --restart=always \
    -p 8030:80 \
    -v /data/www/snai.cms.api/storage:/app/storage \
    -v /data/www/snai.cms.api/config.json:/app/config.json \
Snai.CMS.Api_"$imagev"; 

# 删除旧容器和镜像
docker rm Snai.CMS.Api_"$imagev0"; 
docker rmi Snai.CMS.Api_"$imagev0";