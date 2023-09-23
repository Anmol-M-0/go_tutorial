# docker build -f build/Dockerfile.build -t go_tutorial:latest .
docker build -f build/Dockerfile.binaries -t go_tutorial:latest --output=bin --target=binaries .

docker run go_tutorial:latest
