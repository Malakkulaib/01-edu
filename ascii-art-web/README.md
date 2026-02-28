# ASCII Art Web

## Description

ASCII Art Web is a web application built using Go (Golang). It allows users to enter text into a form and generate ASCII Art using different banner styles such as **standard**, **shadow**, and **thinkertoy**.

The application handles HTTP requests, validates user input, loads the selected banner file, converts text into ASCII art format, and renders the result using HTML templates.

The project also includes proper error handling with custom pages for:
- `400 Bad Request`
- `404 Not Found`
- `500 Internal Server Error`
- `405 Method Not Allowed`

---

## Authors

- Malak Kulaib
- Ahmed Fahmy
- Anton Maher

---

## Usage: How to Run

**1. Clone the repository:**

```bash
git clone https://github.com/Malakkulaib/01-edu.git
```

**2. Enter the project directory:**

```bash
cd ascii-art-web
```

**3. (Optional) Reset to the required commit:**

```bash
git reset --hard c638c70889a22b2a2e2917236eb64f71d45036fd
```

**4. Run the server:**

```bash
go run .
```

**5. Open your browser and visit:**

```
http://localhost:3112
```

---

## Implementation Details: Algorithm

1. The HTTP server starts using `http.ListenAndServe`.
2. Routes are registered using `http.HandleFunc`.
3. The form sends a `POST` request to `/ascii-art`.
4. The handler verifies that the request method is `POST`.
5. The server reads the user input (`text`) and the selected banner style.
6. Input validation is performed — if the input is empty or contains invalid characters, a `400` error page is rendered.
7. The corresponding banner file is loaded from the file system.
8. The input text is split into lines.
9. Each character is looked up in the banner file and converted into its ASCII art representation.
10. The final ASCII result is constructed efficiently using `strings.Builder`.
11. The result is passed to an HTML template and rendered for the user.
12. If any error occurs during processing, a custom error page is rendered with the appropriate HTTP status code.
