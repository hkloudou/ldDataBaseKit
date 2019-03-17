
init:
	@git config --local user.name hkloudou
	@git config --local user.email hkloudou@gmail.com
	@git config --local user.signingkey 575A0CADC23D0A96
	@git config --local commit.gpgsign true
	@git config --local autotag.sign true
all:
	@echo "make"
git:
	git autotag -commit 'auto commit' -tag=true -push=true
test:
	@go test ./
