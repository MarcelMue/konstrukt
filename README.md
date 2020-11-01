# konstrukt

A simple commandline tool to generate SVGs roughly fitting the area of [concrete art](https://en.wikipedia.org/wiki/Concrete_art).

## Example usecase

The pattern of the carpet in the shining (code initially implemented [here](https://github.com/ajstarks/svgo-workshop/blob/master/code/svgplay-samples/shining.go)) can be easily generated and recolored with konstrukt.

The original output can be generated with:
```
konstrukt gen shining
```
Resulting in this image:
![Shining](samples/shining.svg)

Changing the colors can be easily done like this:
```
konstrukt gen shining --foreground "#2c2c54" --background "#d1ccc0" --accent "#ffb142" --filename shining-spanish.svg
```
Resulting in this image:
![Shining Spanish color inspiration](samples/shining-spanish.svg)
