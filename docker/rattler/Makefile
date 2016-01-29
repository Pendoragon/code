# Makefile
#
# Targets:
#   build
#   test
#   release
TAG=v0.0.1

release:
	docker build -t tupachydralisk/rattler:${TAG} .
	docker push tupachydralisk/rattler:${TAG}
.PHONY: release
