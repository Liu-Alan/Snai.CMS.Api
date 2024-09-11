FROM alpine

WORKDIR /app
COPY ./ .

RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 80

ENTRYPOINT ["./Snai.CMS.Api"]