go build .
.\api-client-gen.exe -inputFile="./swagger.json" -outputDirectory="./output" -regex="(?:.*\/v1\/)+({action})" -language=typescript