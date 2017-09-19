IMAGE=localhost:5000/slotprovider
GIT_COMMIT?=latest

deps:
	godep save .

test:
	godep go test .

docker:
	docker build -t $(IMAGE):$(GIT_COMMIT) -f Dockerfile .

docker-bench:
	docker run --rm -t $(IMAGE):$(GIT_COMMIT) -- "go test -bench=."

bench-int:
	godep go test -bench=Benchmark_Atomic*

bench:
	godep go test -bench=.

clean:
	docker rmi $(IMAGE):$(GIT_COMMIT) $(IMAGE):latest