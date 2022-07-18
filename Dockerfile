FROM registry.fedoraproject.org/fedora:36

RUN dnf -y install golang

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o notifier

CMD ["/app/notifier"]