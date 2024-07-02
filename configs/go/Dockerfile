# ------------------------------------------------
# ---               GOLANG                     ---
# ------------------------------------------------
FROM golang:1.23
ENV GO111MODULE=on \
GOOS=linux \
GOARCH=amd64

# Git is required for fetching the dependencies.
EXPOSE 8080

RUN export GO111MODULE=${GO111MODULE}

ADD ./hot-reload.sh /root/

WORKDIR $GOPATH/src/app

COPY ./src .

RUN go mod download
RUN go mod tidy

# option hot reload :
# -- Gin --
RUN go get github.com/codegangsta/gin
RUN go install github.com/codegangsta/gin@latest
ENTRYPOINT ["/root/hot-reload.sh"]
RUN chmod 755 /root/hot-reload.sh
RUN chmod +x /root/hot-reload.sh
