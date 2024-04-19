# generateTableFromDir

This generates a simple list of the **first level** of a given directory.

The output is provided in either an XLSX or CSV file.

## Usage

```
$ generateTableFromDir

Usage: generateTableFromDir --format <csv|xlsx> --output <filename> <directory>

--format      csv | xlsx
--output      { provide a file name, possibly with a path }
<directory>   provide a dir to scan
```

## Generate a build

You'll need to have Go installed. Please follow instructions at https://go.dev/doc/install(https://go.dev/doc/install)

```
$ go build -o generateTableFromDir main.go
```

For macOS there is already a binary in the repo: `bin/generateTableFromDir`.
