run:
	go run ./main -v interval "[1,10][20,23][6,19][2,7]"

go_build_version=`cat ./version.yaml |awk -F" " '{print $$2}'`
build:
	go build -o ./merge -ldflags "-X main.Version=$(go_build_version)"

test:
	go test ./... -cover
