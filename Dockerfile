FROM golang
WORKDIR /temp/
COPY . .
RUN go build cmd
CMD ["cmd"]
