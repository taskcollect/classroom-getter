distroless:
	docker build . -f Dockerfile.distroless -t ghcr.io/taskcollect/classroom-getter:latest

alpine:
	docker build . -f Dockerfile.alpine -t ghcr.io/taskcollect/classroom-getter:alpine