# api-client-gen action

This action Generate API Clients by Swagger.

## Inputs

### `inputFile`

**Required.** Path to swagger.json file.

### `outputDirectory`

**Required.** Path to output files directory.

### `language`

**Required.** Programming language for which clients will be generated.

## Outputs

### `result`

The state of the action, if the API Clients generated successfully.

## Example usage

```
uses: actions/api-client-gen@v1
with:
    inputFile: ./swagger.json
    outputDirectory: ./out
    language: typescript
```
