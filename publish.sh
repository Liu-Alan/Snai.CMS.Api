cd /data/www/snai.cms.api;     
docker stop Snai.CMS.Api; 
docker rm Snai.CMS.Api;     
docker rmi Snai.CMS.Api-Service:v1.0;     
docker build -t Snai.CMS.Api-Service:v1.0 . ;    
docker run -d -p 8030:80 --restart always --name Snai.CMS.Api \
 -v /data/www/snai.cms.api/storage:/app/storage \
 -v /data/www/snai.cms.api/config.json:/app/config.json \
Snai.CMS.Api-Service:v1.0