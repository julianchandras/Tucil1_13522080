go build -o bin/main src/main.go src/functions.go src/txtprocessor.go src/type.go

if [ $? -eq 0 ]; then
    ./bin/main
else
    echo "Build failed!"
fi