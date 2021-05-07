go build .
.\api-client-gen.exe -inputFile="./swagger.json" -outputDirectory="./output" -regex=".*\/{version}\/{action}" -language=typescript