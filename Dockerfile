FROM gocv/opencv:4.5.4

ENV GOPROXY "https://goproxy.io/"

COPY . /home
WORKDIR /home
RUN go build .

EXPOSE 8000
CMD ["/home/face-detection"]
