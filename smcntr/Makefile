
.PHONY: build run

build-single:
	@docker build --no-cache -t funcimp/shorts/smcntr:singlestage -f Dockerfile.single .

build-multi:
	@docker build --no-cache -t funcimp/shorts/smcntr:multistage -f Dockerfile.multi .

compare:
	@clear
	docker images funcimp/shorts/smcntr

build: build-single build-multi

run:
	@docker run -d -p 1337:1337 -it funcimp/shorts/smcntr:multistage /bin/bash