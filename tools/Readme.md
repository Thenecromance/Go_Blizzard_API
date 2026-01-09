# Blizzard API Code Generator

This tool leverages powerful Go templates to automatically generate Blizzard API client code.

It is designed to generate the latest client code based on current Blizzard API definitions. Response structures are automatically generated using [go-jsonstruct](github.com/twpayne/go-jsonstruct).

> **Note:** You must provide a valid Blizzard API Client ID and Secret to fetch the live data required for model generation.

This tool automatically generates the following contents:
- API call functions
- Request parameter structures
- Return value structures
- Routers (currently supported: Gin)

You can customize which API categories to generate by modifying [main.go](./main.go).

## Extra Files
The tool will automatically generate `./params.json`. You can modify the parameters in this file to customize the API requests used to generate the response models.

## How to Support Other Languages?
1. Rewrite `api.tmpl` and `model.tmpl` to match the syntax of your target language.
2. Integrate a JSON-to-struct converter for your target language (e.g., quicktype).
3. Run the tool.

## Notes

- This is currently a prototype version. Code refactoring is planned for future updates.
- To leverage Go's fast compilation speed, generation controls are currently hardcoded in `main.go`.

