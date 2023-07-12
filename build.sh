
export GOOS=linux
export GOARCH=amd64
rm -rf ./output
go build -o sentinel
mkdir output
cp -r ./release/* ./output
mv sentinel ./output

zip ./output/sentinel-${GOOS}-${GOARCH}.zip ./output/*