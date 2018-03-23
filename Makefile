all:
	@echo "make"
	go test ./
run:
	@make push
push:
	-git add . && git commit -m 'build auto commit' && git push origin master
