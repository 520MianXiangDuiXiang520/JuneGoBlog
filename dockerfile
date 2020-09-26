FROM golang:latest

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
WORKDIR /blog
COPY . /blog
RUN chmod 777 ./run.sh
RUN chmod +x ./bin/JuneGoBlog
RUN make build
EXPOSE 8080
ENTRYPOINT ["./run.sh"]
