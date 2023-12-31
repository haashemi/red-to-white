# red-to-white

This project (less than 50 lines, lol) is made just to convert Fortnite's material files to white transparent PNGs.

## Usage:

- Add all icons you want to convert to the `./input` folder.
- Run the application using `go run .`
- Get the converted files from the `./export` folder.

## What have I learned from this?

**NEVER EVER** use _RGBA_ again if you're gonna encode it as _PNG_.

### Why?

It's clearly documented at [image/png.Encode](https://pkg.go.dev/image/png#Encode) method.

> Any Image may be encoded, but images that are not image.NRGBA might be encoded lossily.

I didn't dive into the technical reasons behind it, but I also didn't expect to see any issues at all!

But it happened, and just this single case means that it's better to not accept the risk again and just go with NRGBA from now on.

### Here's an example:

| RAW                                  | NRGBA                                        | RGBA                                         |
| ------------------------------------ | -------------------------------------------- | -------------------------------------------- |
| ![Raw Image](/readme-assets/raw.png) | ![Success Image](/readme-assets/success.png) | ![Failure Image](/readme-assets/failure.png) |
