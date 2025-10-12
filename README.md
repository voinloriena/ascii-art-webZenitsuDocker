# ASCII Art Web

## Description

ASCII Art Web is a web application that allows users to convert text into ASCII art using different banner styles. The project implements a Go-based HTTP server with a web GUI that supports generating ASCII art in three different banners: standard, shadow, and thinkertoy.

## Features

- Web-based GUI for generating ASCII art
- Support for three banner styles:
  - Standard
  - Shadow
  - Thinkertoy
- Input validation for ASCII characters
- Error handling for various scenarios


### Running the Application

1. Clone the repository:
   ```bash
   git clone https://01.tomorrow-school.ai/git/ynurmakh/ascii-art-web
   cd ascii-art-web
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Open a web browser and navigate to `http://localhost:8080`

4. Enter text in the input field
5. Select a banner style
6. Click submit to generate ASCII art

## Implementation Details

### Algorithm

The ASCII art generation follows these key steps:

1. Validate input text (ASCII characters only)
2. Verify banner file integrity using SHA-256 hashing
3. Process each character of the input:
   - Retrieve corresponding ASCII art lines from the selected banner file
   - Construct the full ASCII art representation
4. Handle special cases like newline characters and empty lines

### Project Structure

```
├── main.go          # HTTP server and route handlers
├── ascii.go         # ASCII art generation logic
├── banners/         # Banner text files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/       # HTML templates
│   ├── index.html
│   └── ascii-art.html
└── static/          # CSS and other static files
    └── style.css
```

### HTTP Endpoints

- `GET /`: Main page with input form
- `POST /ascii-art`: Generate and display ASCII art

### Error Handling

The application handles various error scenarios with appropriate HTTP status codes:
- 200 OK: Successful requests
- 400 Bad Request: Invalid input
- 404 Not Found: Missing resources
- 500 Internal Server Error: Server-side issues
