# api-client-gen action

This action Generate API Clients by Swagger.

## Inputs

### `inputFile`

**Required.** Path to swagger.json file.

### `outputFolder`

**Required.** Path to output files folder.

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
    outputFolder: ./out
    language: typescript
```
