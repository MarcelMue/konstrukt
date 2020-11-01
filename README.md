# konstrukt

A simple commandline tool to generate SVGs roughly fitting the area of [concrete art](https://en.wikipedia.org/wiki/Concrete_art).

## Usage

This sections shows example usage of the implemented commands.

### Shining

The pattern of the carpet in the shining (code initially implemented [here](https://github.com/ajstarks/svgo-workshop/blob/master/code/svgplay-samples/shining.go)) can be easily generated and recolored with konstrukt.

The original output can be generated with:
```
konstrukt gen shining
```

![Shining](samples/shining.svg)

Changing the colors can be easily done like this:
```
konstrukt gen shining --foreground "#2c2c54" --background "#d1ccc0" --accent "#ffb142" --filename shining-spanish.svg
```

![Shining Spanish color inspiration](samples/shining-spanish.svg)

### Interruptions

The pattern and default implementation is inspired by [Horst Bartnigs](https://de.wikipedia.org/wiki/Horst_Bartnig) collection of works titled `72 Unterbrechungen`.

The default can be generated with:
```
konstrukt gen interruptions
```

![Interruptions](samples/interruptions.svg)

Changing the colors can be easily done like this:
```
konstrukt gen interruptions --color1 "#00a8ff" --color2 "#2f3640" --color3 "#fbc531" --filename interruptions-british.svg
```

![Interruptions British color inspiration](samples/interruptions-british.svg)
