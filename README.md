# konstrukt

A simple commandline tool to generate SVGs roughly fitting the area of [concrete art](https://en.wikipedia.org/wiki/Concrete_art).

## Usage

This sections shows example usage of the implemented commands.

### Shining

The pattern and default implementation is inspired by a carpet design by [David Nightingale Hicks](https://en.wikipedia.org/wiki/David_Nightingale_Hicks) commonly known as "Hicks' Hexagon".
It became known in popular culture as the pattern of the carpet in [The Shining](https://en.wikipedia.org/wiki/The_Shining_(film)).
The code for the pattern was initially implemented [here.](https://github.com/ajstarks/svgo-workshop/blob/master/code/svgplay-samples/shining.go))

<details>
<summary>Usage examples</summary>

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

</details>

### Interruptions

The pattern and default implementation is inspired by [Horst Bartnigs](https://de.wikipedia.org/wiki/Horst_Bartnig) collection of works titled `72 Unterbrechungen`.

<details>
<summary>Usage examples</summary>

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

</details>

### Quadrat

The pattern and default implementation is inspired by [Horst Bartnigs](https://de.wikipedia.org/wiki/Horst_Bartnig) work titled `8 blaue und 8 schwarze Quadrate`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen quadrat
```

![Quadrat](samples/quadrat.svg)

Changing the colors can be easily done like this:
```
konstrukt gen quadrat --color1 "#f6b93b" --color2 "#38ada9"  --filename quadrat-french.svg
```

![Quadrat French color inspiration](samples/quadrat-french.svg)

</details>

### Janein

The pattern and default implementation is inspired by [Wolfgang Bosses](https://kulturanalyse.de//wolfgang_bosse/index.html) work titled `JA-NEIN`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen janein
```

![Janein](samples/janein.svg)

Changing the colors can be easily done like this:
```
konstrukt gen janein --color1 "#ced6e0" --color2 "#ffa502"  --filename janein-cn.svg
```

![Janein CN color inspiration](samples/janein-cn.svg)

</details>

### Fiftyfive

The pattern and default implementation is inspired by [Julia Breunigs](https://juliaskonkretekunst.wordpress.com/) work titled `Bild Nr. 55`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen fiftyfive
```

![Fiftyfive](samples/fiftyfive.svg)

Changing the colors can be easily done like this:
```
konstrukt gen fiftyfive --color1 "#cd84f1" --color2 "#4b4b4b" --color3 "#ffaf40"  --filename fiftyfive-tr.svg
```

![Fiftyfive TR color inspiration](samples/fiftyfive-tr.svg)

</details>

### Ladysweat

The pattern and default implementation is inspired by [Joshua Blankenships](https://blankenship.xyz/) work titled `Lady Sweat Repeating Pattern`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen ladysweat
```

![Ladysweat](samples/ladysweat.svg)

Changing the colors can be easily done like this:
```
konstrukt gen ladysweat --color1 "#303952" --color2 "#f8a5c2" --color3 "#778beb"  --filename ladysweat-ru.svg
```

![Ladysweat RU color inspiration](samples/ladysweat-ru.svg)

</details>

### Modernhive

The pattern and default implementation is inspired by [Emma Methods](http://www.emmamethod.com/) work titled `modern hive`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen modernhive
```

![Modernhive](samples/modernhive.svg)

Changing the colors can be easily done like this:
```
konstrukt gen modernhive --color1 "#B33771" --color2 "#55E6C1"  --filename modernhive-in.svg
```

![Modernhive IN color inspiration](samples/modernhive-in.svg)

</details>

### Swiss16

The pattern and default implementation is inspired by [Neo Geometrics](https://dribbble.com/vladikkk09) work titled `swiss-16`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen swiss16
```

![Swiss16](samples/swiss16.svg)

Changing the colors can be easily done like this:
```
konstrukt gen swiss16 --color1 "#222f3e" --color2 "#5f27cd" --color3 "#1dd1a1"  --filename swiss16-ca.svg
```

![Swiss16 CA color inspiration](samples/swiss16-ca.svg)

</details>

### Interlocking

The pattern and default implementation is inspired by [Cami Dobrins](https://camidraws.com/) work titled `Interlocking Abstract Pattern Background`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen interlocking
```

![Interlocking](samples/interlocking.svg)

Changing the colors can be easily done like this:
```
konstrukt gen interlocking --color1 "#f7f1e3" --color2 "#40407a" --color3 "#33d9b2"  --filename interlocking-es.svg
```

![Interlocking ES color inspiration](samples/interlocking-es.svg)

</details>

## Guidelines

Additions to this project should follow these guidelines:
- Patterns should be repeatable and scaleable.
- Patterns should be constant across the generated SVG (e.g. no convex pattern).
- All patterns should be made up of low count polygons (no curves or similar).
- Accurate citation should be given if a pattern was inspired by an artwork or artist.
