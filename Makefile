default: build

race:
	go install github.com/murphybytes/go-presentation/race
	go build -race -o race2 github.com/murphybytes/go-presentation/race
	@cp race2 $(GOPATH)/bin/.

async:
	go install -race github.com/murphybytes/go-presentation/async

fan-in:
	go install -race github.com/murphybytes/go-presentation/fan-in

fan-in2:
	go install -race github.com/murphybytes/go-presentation/fan-in2

worker:
	go install -race github.com/murphybytes/go-presentation/worker

pipeline:
	go install -race github.com/murphybytes/go-presentation/pipeline

clean:
	-rm $(GOPATH)/bin/*



build: race async fan-in fan-in2 worker pipeline

.PHONY: race async build fan-in fan-in2 worker pipeline 
