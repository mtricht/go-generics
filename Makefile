# go build tags are either comma or space separated
# when a value such as the default date format contains a space
# you'd get an error:
# 	go: -tags space-separated list contains comma
# so make sure there is no space
build:
	commit=$(git rev-parse HEAD)
	date=$(date +%Y%m%d)
	go build -tags "commit=$commit,date=$date" cmd/main.go