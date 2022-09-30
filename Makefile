GOPATH=${HOME}/go

%:
	@true

.PHONY: proto
proto:
	./scripts/proto.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: run
run:
	./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))