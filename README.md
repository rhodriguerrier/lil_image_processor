# lil_image_processor

![Lil' Image Processor](https://github.com/rhodriguerrier/lil_image_processor/blob/main/img_process_examples/sobel_colour_japanese_wave.jpg?raw=true)

A command line based image processor using Go.

This allows the user to select an input image and apply one of the following effects to it:
- Grayscale
- Sharpening
- Gaussian Blur
- Sobel Edge Detection
- Sobel Edge Detection with Orientation Colouring

The program takes a number of flags to run:
- ``file (input file, the cli supports .jpg and .png files)``
- ``editMode (the mode of editing, sc=Sobel with colouring, s=Sobel, g=grayscale, b=gaussian blur, sh=sharpening)``
- ``outFile (output file name, must include .jpg or .png file extensions)``

To run, simply clone down and cd into the repository. Next, you will need to build:

``go build main.go``

The next command is dependent on OS. If you are using mac/linux or a linux based terminal in windows you can do the following:

``./main -file=example.jpg -editMode=sc -outFile=sobel_colour_example.jpg``

If you are using windows, you will need to provide the full file path for -file and -outFile:

``./main -file=C:\Users\someone\go\src\github.com\someone\lil_image_processor\example.jpg -editMode=sc -outFile=...``

Future Work:
- Allow chaining on image processes (e.g. being able to blur and then apply sobel edge detection in one command)
- Add in some more basic image processes (e.g. rotation, resizing)
- Allow custom kernel values for image effects
- Add in concurrent calculations for convolutions. This is not necessary but for bigger images there is a half second-ish delay which would be nice to cut down
- Better error handling


