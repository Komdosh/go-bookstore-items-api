FROM golang

ENV REPO_URL=github.com/Komdosh/go-bookstore-items-api
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src

ENV ELASTIC_HOSTS=http://es01:9200

COPY src $WORKPATH
WORKDIR $WORKPATH
RUN go build -o items-api .

#Expose port
EXPOSE 8000

CMD ["./items-api"]