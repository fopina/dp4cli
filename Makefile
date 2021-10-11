dp4cli.exe: dp4cli.go
	GOOS=windows GOARCH=386 go build dp4cli.go

.PHONY: builder
builder:
	docker build -t dp4cli-wine --target wine .

test-setup: dp4cli.exe
	docker run --rm -i \
	           -v ${PWD}:/app \
	           -v ${PWD}/data.test:/root/.wine/drive_c/users/root/Application\ Data/dp4cli \
			   -w /app dp4cli-wine \
			   wine ./dp4cli.exe -setup

test-pin: dp4cli.exe
	docker run --rm -i \
	           -v ${PWD}:/app \
	           -v ${PWD}/data.test:/root/.wine/drive_c/users/root/Application\ Data/dp4cli \
			   -w /app dp4cli-wine \
			   wine ./dp4cli.exe
