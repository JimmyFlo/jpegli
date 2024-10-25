# jpegli

A simple tool to reduce the pixel- and byte-size of JPEG, PNG, or WebP images using the JPEGli compressor.

(WebP images are likely to be a little smaller; but JPEG images still have greater support across the web!)

Currently scales to a maximum of 2048×1920px and a(n excellent) quality of 75.

## Usage

```bash
$ go install github.com/jphastings/jpegli@latest
$ brew install jphastings/tools/jpegli

$ ls -lah *.{jpg,png,webp}
619K  example1.png
4.0M  example2.jpg
173K  example3.webp

$ jpegli *.{jpg,png,webp}
Complete. 3 images standardized

$ ls -lah *.{jpg,png,webp}
619K  example1.png
 59K  example1.jpg

4.0M  example2.original.jpg
566K  example2.jpg

173K  example3.webp
202K  example3.jpg
```
