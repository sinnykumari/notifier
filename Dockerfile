FROM registry.fedoraproject.org/fedora:36

RUN dnf -y install golang

copy *.go ./

RUN go build -o notifier
CMD notifier