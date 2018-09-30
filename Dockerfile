FROM alpine
RUN apk update && apk add ca-certificates
ADD scraper /scraper
CMD ["./scraper"]