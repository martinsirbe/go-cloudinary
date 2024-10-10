# Cloudinary CLI

An unofficial command-line interface (CLI) for Cloudinary, written in Go, leveraging the official
Cloudinary SDK. This tool allows users to interact with Cloudinary services directly from the
command line, making it easier to manage and upload media assets.

## Features

- Upload image, video and audio files to Cloudinary.
- Specify upload presets and folders.
- Process entire directories of files.
- Specify file extensions to filter files for upload.
- Utilise the power and simplicity of Go for fast and efficient CLI operations.
- Support multiple Cloudinary accounts.

## Supported File Formats

The following file extensions are supported for upload.

### Image
`ai`, `gif`, `png`, `webp`, `bmp`, `bw`, `djvu`, `dng`, `ps`, `ept`, `eps`, `eps3`, `fbx`, `flif`, `glb`, `gltf`, `heif`, `heic`, `ico`, `indd`, `jpg`, `jpeg`, `jp2`, `wdp`, `jxr`, `hdp`, `jxl`, `obj`, `pdf`, `ply`, `psd`, `arw`, `cr2`, `cr3`, `svg`, `tga`, `tif`, `tiff`, `u3ma`, `usdz`.
For more details, see the [supported image formats](https://cloudinary.com/documentation/image_transformations#supported_image_formats).

### Video
`3g2`, `3gp`, `avi`, `flv`, `m3u8`, `ts`, `m2ts`, `mts`, `mov`, `mkv`, `mp4`, `mpeg`, `mpd`, `mxf`, `ogv`, `webm`, `wmv`.
For more details, see the [supported video formats](https://cloudinary.com/documentation/video_manipulation_and_delivery#supported_video_formats).

### Audio
`aac`, `aiff`, `amr`, `flac`, `m4a`, `mp3`, `ogg`, `opus`, `wav`.
For more details, see the [supported audio formats](https://cloudinary.com/documentation/audio_transformations#supported_audio_formats).

## Installation

For convenience, pre-built binaries are available for various platforms. Download the appropriate
binary for your system from the releases page.

### Homebrew

To install the Year Progress Indicator using Homebrew on macOS or Linux, you can follow these steps:

```shell
brew tap martinsirbe/clinkclank
brew install martinsirbe/clinkclank/cld
```

This will add the custom tap and install the `cld` CLI, making it readily accessible from any
terminal.

### Build from Source

Make sure you have Go installed. You can then install the Year Progress Indicator globally via the
following command:

```shell
go install github.com/martinsirbe/go-cloudinary/cmd/cld@v0.0.1
```

This command compiles and installs the binary to your Go bin directory, making it accessible from
any terminal provided the directory is in your system's PATH.

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

Optionally, set additional parameters, for example, `upload_prefix` and `secure_distribution`, to
the environment variable:

Example:

```bash
export CLOUDINARY_URL=cloudinary://api_key:api_secret@my_cloud_name?secure_distribution=example.com&upload_prefix=example
```

### Multiple URL Support

You can now configure multiple Cloudinary URLs by separating them with commas in
the `CLOUDINARY_URL` environment variable. This is useful if you need to work with multiple
Cloudinary accounts or environments.

Example:

```bash
export CLOUDINARY_URL=cloudinary://api_key:api_secret@my_cloud_name,cloudinary://api_key:api_secret@my_other_cloud_name
```

If multiple URLs are configured but no specific URL is selected when invoking the CLI, the **first
URL** will be used by default. To select a specific account, you can provide the `-a` or `--api-key`
option when running the CLI:

Example:

```bash
cld -a api_key
```

This will instruct the CLI to use the account associated with the provided `api_key`. Ensure that
each `cloudinary://` URL has its own `api_key`, `api_secret`, and `cloud_name`.

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
