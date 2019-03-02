
init:
	@git config --local user.name hkloudou
	@git config --local user.email hkloudou@gmail.com
	@git config --local user.signingkey 575A0CADC23D0A96
	@git config --local commit.gpgsign true
	@git config --local autotag.sign true
all:
	@echo "make"
git:
	- git add . && git commit -m 'auto commit' && git push origin master -f --tags
tag:
	- git add . && git commit -m 'auto tag'
	- git autotag && git push origin master -f --tags
	@echo `git describe` > version
	@echo "current version:`git describe`"
test:
	@go test ./
