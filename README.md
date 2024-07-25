# Cloudinary CLI

An unofficial command-line interface (CLI) for Cloudinary, written in Go, leveraging the official
Cloudinary SDK. This tool allows users to interact with Cloudinary services directly from the
command line, making it easier to manage and upload media assets.

## Features

- Upload images to Cloudinary.
- Specify upload presets and folders.
- Process entire directories of image files.
- Specify file extensions to filter files for upload.
- Utilise the power and simplicity of Go for fast and efficient CLI operations.

## Installation

For convenience, pre-built binaries are available for various platforms. Download the appropriate binary for your system from the releases page.

### Homebrew

To install the Year Progress Indicator using Homebrew on macOS or Linux, you can follow these steps:

```shell
brew tap martinsirbe/clinkclank
brew install martinsirbe/clinkclank/cld
```

This will add the custom tap and install the `cld` CLI, making it readily accessible from any terminal.

### Build from Source
Make sure you have Go installed. You can then install the Year Progress Indicator globally via the following command:

```shell
go install github.com/martinsirbe/go-cloudinary/cmd/cld@v0.0.0
```

This command compiles and installs the binary to your Go bin directory, making it accessible from any terminal provided the directory is in your system's PATH.

## Configuration

Before using the CLI, you need to set the `CLOUDINARY_URL` environment variable. This variable
configures the required `cloud_name`, `api_key`, and `api_secret`.

You can set this environment variable by copying the API environment variable format from the API
Keys page of the Cloudinary Console Settings. Replace `api_key` and `api_secret` with your actual
values, while your cloud name is already correctly included in the format.

Example:

```bash
export CLOUDINARY_URL=cloudinary://api_key:api_secret@my_cloud_name
```

Optionally, set additional parameters, for example, `upload_prefix` and `secure_distribution`, to the
environment variable:

Example:

```bash
export CLOUDINARY_URL=cloudinary://api_key:api_secret@my_cloud_name?secure_distribution=example.com&upload_prefix=example
```

## Usage

Basic usage of the Cloudinary CLI:

```bash
cld [ -p <preset> | --preset <preset> ] [ -f <folder> | --folder <folder> ] [ -e <extensions> | --extensions <extensions> ] INPUT
```

### Arguments

- `INPUT`: The input file or directory to process.

### Options

- `-p, --preset`: Cloudinary upload preset.
- `-f, --folder`: Cloudinary upload folder.
- `-e, --extensions`: Comma-separated list of file extensions to upload.

## Examples

To upload a single image:
```bash
cld test.webp
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721945485/benqcqpstfjv5noutjih.webp
```

To upload a single image using a preset:
```bash
cld -p test-preset test.webp
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721945634/test-folder/test.webp
```

To upload a single image using a preset and folder:
```bash
cld -p test-preset -f test-folder/hello test.webp
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721945718/test-folder/hello/test.webp
```

To upload all images with a specific file extension `webp`:
```bash
cld -p test-preset -e webp .
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721945983/test-folder/test_2xl.webp
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721945984/test-folder/test_lg.webp
```

To upload all images with specific file extensions `jpg` and `png`:
```bash
cld -p test-preset -e jpg,png .
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721946123/test-folder/test.jpg
image uploaded: https://res.cloudinary.com/dyodcxgg5/image/upload/v1721946124/test-folder/test.png
```

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

Please ensure your code adheres to the following guidelines:

- Follow Go conventions and idiomatic practices.
- Have fun.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

## Disclaimer

This is an unofficial tool and is not endorsed or supported by Cloudinary. Use it at your own risk.
