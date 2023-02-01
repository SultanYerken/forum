FROM golang

WORKDIR /app

COPY . .

LABEL project="forum" \
      author="Sultanye" \
      link="https://01.alem.school/git/Sultanye/forum"

RUN go build -o forum  cmd/main.go


EXPOSE 8080
CMD ["./forum"]
