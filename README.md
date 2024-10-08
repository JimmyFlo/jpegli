# jpegli

A helper tool for my own use (and maybe yours?) that reduces the pixel and byte size of an image in JPG, PNG or WEBP format by transcoding with JPEGli.

(WebP images are likely to be a little larger; but JPEG images still have greater support across the web!)

```bash
$ go install github.com/jphastings/jpegli@latest

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
