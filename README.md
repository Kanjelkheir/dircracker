# DirCracker

DirCracker is a command-line tool written in Go that helps you enumerate valid directories on a target web server using a wordlist. It traverses through a provided list of directories and checks their existence against a specified target URL, reporting which ones are found.

## Features

- **Fast Directory Enumeration**: Utilizes Go's concurrency features to check multiple directories simultaneously.
- **Wordlist Support**: Takes a wordlist file as input, allowing for flexible and extensive directory checks.
- **Clear Output**: Provides distinct messages for found and not-found directories.

## Installation

To install DirCracker, you need to have Go installed on your system (version 1.22 or higher).

1. **Clone the repository:**

   ```bash
   git clone https://github.com/kanjelkheir/dircracker.git
   cd dircracker
   ```

2. **Build the executable:**

   ```bash
   go build -o bin/dircracker ./cmd/dircracker
   ```

   This will create an executable named `dircracker` in the `bin/` directory.

## Usage

DirCracker requires two main arguments: a wordlist file (`-w` or `--wordlist`) and a target URL (`-t` or `--target`).

### Basic Usage

To run DirCracker, use the following command:

```bash
./bin/dircracker -w <path_to_wordlist.txt> -t <target_url>
```

- `-w`, `--wordlist`: Specifies the path to the wordlist file containing directories to check. Each directory should be on a new line and start with a `/`.
- `-t`, `--target`: Specifies the base URL of the target web server (e.g., `https://example.com`).

### Example

Let's assume you have a `wordlist.txt` file with the following content:

```
/admin
/test
/api
/files
/images
```

And you want to check these directories against `https://google.com`:

```bash
./bin/dircracker -w wordlist.txt -t https://google.com
```

### Expected Output

The tool will print messages indicating whether each directory was found or not:

```
https://google.com/admin not found!
https://google.com/test not found!
https://google.com/api not found!
https://google.com/files found!
https://google.com/images not found!
```

### Wordlist Format

Your wordlist file (`wordlist.txt` in the example) should contain one directory per line, and each directory **must** start with a forward slash (`/`).

**Correct format:**
```
/dir1
/path/to/dir2
/another_dir
```

**Incorrect format:**
```
dir1
path/to/dir2
```

## How it Works

DirCracker operates in the following steps:

1. **Argument Parsing**: It parses the command-line arguments for the wordlist file and the target URL.
2. **Wordlist Reading**: It reads the provided wordlist file, ensuring it contains valid directory formats (starting with `/`).
3. **Directory Enumeration**: For each directory in the wordlist, it constructs a full URL by appending the directory to the target URL.
4. **Concurrent Checks**: It concurrently sends HTTP GET requests to each constructed URL.
5. **Status Reporting**: Based on the HTTP response, it determines if the directory exists (any status code other than 404 Not Found) and prints the result. Specific error handling is in place for network issues or server-side errors.

## Contributing

If you'd like to contribute to DirCracker, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and write tests if applicable.
4. Ensure your code adheres to Go best practices and the existing coding style.
5. Submit a pull request.
