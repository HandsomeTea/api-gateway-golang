from golang:1.23.6
WORKDIR /usr/src/app
COPY ./tmp/api-gateway /usr/src/app/

CMD ["/usr/src/app/api-gateway"]
