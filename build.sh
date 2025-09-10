echo building gohttp $(cat version) ...
go version
echo current directory $(basename $(pwd))
go build -o bin/gohttp cmd/*.go internal/parser/*.go 
