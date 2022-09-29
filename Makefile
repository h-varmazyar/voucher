GOPATH=${HOME}/go

%:
	@true

.PHONY: proto
proto:
	./scripts/proto.sh $(filter-out $@,$(MAKECMDGOALS))