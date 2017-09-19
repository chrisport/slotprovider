# This Dockerfile can be used to Profile application in case of performance problems
FROM golang

# config
ENV HOME="/go/src/github.com/chrisport/slotprovider"

# build
WORKDIR $HOME

RUN apt-get update
RUN apt-get install -y graphviz

ADD . .

RUN go get github.com/tools/godep && \
        godep go build github.com/chrisport/slotprovider

# run
ENTRYPOINT ["/bin/sh","-c"]
CMD ["go test -run=^$ -bench=BenchmarkCall -cpuprofile=cpu.out && go tool pprof --pdf slotprovider.test cpu.out"]
