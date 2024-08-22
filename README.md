# Structure-Walker

**Structure-Walker** is an interactive demonstration and walkthrough of data structures, written in C and Go. It utilizes C for data structure implementation and Go for interactivity and walkthrough features.

## Features
- Interactive demonstration of various data structures.
- Walkthroughs to understand how each data structure works.
- Written in C (for data structures) and Go (for interactivity).

## Usage

To use the project, run the compiled executable. You can either download it from the [releases page](#) (work in progress) or compile it manually.

### Running the Application

```sh
./data-structures
```

or on Windows:

```sh
./data-structures.exe
```

## Compiling

### Prerequisites

This project uses C for data structure implementation and requires the following:

- **C Library:** `murmurhash.c` for hashing in the hash table. You can install it from its [GitHub repository](https://github.com/jwerle/murmurhash.c). Make sure it's placed in `/usr/local/lib` or update the path as needed in the source code.

- **Go Modules:** Download Go modules with:

  ```sh
  go mod download
  ```

### Build Instructions

1. **Compile the Project**

   ```sh
   go build .
   ```

2. **Run the Application**

   ```sh
   ./data-structures
   ```

   or on Windows:

   ```sh
   ./data-structures.exe
   ```

## Contributing

If you’d like to contribute, please fork the repository and submit a pull request. See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

For questions or feedback, please contact [your email address or GitHub profile](#).
