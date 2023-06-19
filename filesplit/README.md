# filesplit

A cli to split and save a file into Memcached.

## Usage

### Start

```bash
make build
build/filesplit
```

### Set file

```bash
filesplit set file.txt
```

### Get file

```bash
filesplit get file.txt
```

### Delete file

```bash
filesplit delete file.txt
```

## Useful commands

Create a 15MB file:

```bash
dd if=/dev/zero of=file.txt bs=15MB count=1
```
