# konstrukt

A simple commandline tool to generate SVGs roughly fitting the area of [concrete art](https://en.wikipedia.org/wiki/Concrete_art).

## Getting started

The `konstrukt` binary can be downloaded from [releases](https://github.com/MarcelMue/konstrukt/releases) section of this repository.
It is contained together with the [LICENSE](LICENSE) in a `.tar.gz` or `.zip` depending on the target operating system.

You can interact with `konstrukt` through your native [command line interface](https://en.wikipedia.org/wiki/Command-line_interface).

If you are unfamiliar with a command line then these examples should give you some guidance:
<details>
<summary>Using Windows</summary>

First download the latest `.zip` release from the [releases](https://github.com/MarcelMue/konstrukt/releases) page.

Unzip the archived files in a location of your choice (e.g. your `Downloads` folder).

Open a command line of your choice (e.g. `cmd.exe`).

Navigate tp the location of the unpacked `konstrukt` binary (e.g. `cd C:\Users\SomeUser\Downloads`).

Display the help text of `konstrukt` by calling it without arguments (`konstrukt.exe`).
The output should look like this:
```
Command line tool for generating konstruktive art.

Usage:
  konstrukt [flags]
  konstrukt [command]

Available Commands:
  gen         Generate files.
  help        Help about any command

Flags:
  -h, --help   help for konstrukt

Use "konstrukt [command] --help" for more information about a command.
```

Now you should be able to experiment with the usage examples below just remember to use `konstrukt.exe`!
</details>
<details>
<summary>Using Linux/MacOS</summary>

First download the correct `.tar.gz` release depending on your operating system from the [releases](https://github.com/MarcelMue/konstrukt/releases) page.

Unpack the `.tar.gz` in a location of your choice (e.g. your `Downloads` folder).

Open a terminal of your choice.

Navigate to the location of the unpacked `konstrukt` binary (e.g. `cd ~/Downloads`).

Display the help text of `konstrukt` by calling it without arguments (`./konstrukt`).
The output should look like this:
```
Command line tool for generating konstruktive art.

Usage:
  konstrukt [flags]
  konstrukt [command]

Available Commands:
  gen         Generate files.
  help        Help about any command

Flags:
  -h, --help   help for konstrukt

Use "konstrukt [command] --help" for more information about a command.
```

Now you should be able to experiment with the usage examples below!
</details>

## Usage

All commands allow for the following flags:
```
      --filename string   Name of the output file.
      --height int        Height of the output file in pixels. (default 500)
  -h, --help              help
      --randomize         Randomize all colors in the pattern, ignore other color flags.
      --width int         Width of the output file in pixels. (default 500)
```

Other flags for coloring patterns can be explored in the individual usage examples below.

### Shining

The pattern and default implementation is inspired by a carpet design by [David Nightingale Hicks](https://en.wikipedia.org/wiki/David_Nightingale_Hicks) commonly known as "Hicks' Hexagon".
It became known in popular culture as the pattern of the carpet in [The Shining](https://en.wikipedia.org/wiki/The_Shining_(film)).
The code for the pattern was initially implemented [here.](https://github.com/ajstarks/svgo-workshop/blob/master/code/svgplay-samples/shining.go)

<details>
<summary>Usage examples</summary>

The original output can be generated with:
```
konstrukt gen shining
```

![Shining](samples/shining.svg)

Changing the colors can be easily done like this:
```
konstrukt gen shining --color3 "#2c2c54" --color2 "#d1ccc0" --color1 "#ffb142" --filename shining-es.svg
```

![Shining ES color inspiration](samples/shining-es.svg)

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
konstrukt gen interruptions --color1 "#00a8ff" --color2 "#2f3640" --color3 "#fbc531" --filename interruptions-br.svg
```

![Interruptions BR color inspiration](samples/interruptions-br.svg)

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
konstrukt gen quadrat --color1 "#f6b93b" --color2 "#38ada9"  --filename quadrat-fr.svg
```

![Quadrat FR color inspiration](samples/quadrat-fr.svg)

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

### Whitegold

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `whitegold`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen whitegold
```

![Whitegold](samples/whitegold.svg)

Changing the colors can be easily done like this:
```
konstrukt gen whitegold --color1 "#dff9fb" --color2 "#6ab04c"  --filename whitegold-au.svg
```

![Whitegold AU color inspiration](samples/whitegold-au.svg)

</details>

### Fallingdaggers

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `fallingdaggers`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen fallingdaggers
```

![Fallingdaggers](samples/fallingdaggers.svg)

Changing the colors can be easily done like this:
```
konstrukt gen fallingdaggers --color1 "#e55039" --color2 "#f39c12"  --filename fallingdaggers-in.svg
```

![Fallingdaggers randomized color inspiration](samples/fallingdaggers-in.svg)

</details>

### Whitegold2

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `whitegold2`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen whitegold2
```

![Whitegold2](samples/whitegold2.svg)

Changing the colors can be easily done like this:
```
konstrukt gen whitegold2 --color1 "#1e272e" --color2 "#f53b57"  --filename whitegold2-se.svg
```

![Whitegold2 SE color inspiration](samples/whitegold2-se.svg)

</details>

### Blockplay

The pattern is inspired by [Sophie Adams-Foster](https://www.instagram.com/sharp.line.hunter/) work titled `blockplay`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen blockplay
```

![blockplay](samples/blockplay.svg)

Changing the colors can be easily done like this:
```
konstrukt gen blockplay --color1 "#ced6e0" --color2 "#ff6b81"  --filename blockplay-cn.svg
```

![blockplay CN color inspiration](samples/blockplay-cn.svg)

</details>

### Blockplay2

The pattern is inspired by [Sophie Adams-Foster](https://www.instagram.com/sharp.line.hunter/) work titled `blockplay`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen blockplay2
```

![blockplay2](samples/blockplay2.svg)

Changing the colors can be easily done like this:
```
konstrukt gen blockplay2 --color1 "#ced6e0" --color2 "#ff6b81"  --filename blockplay2-cn.svg
```

![blockplay2 CN color inspiration](samples/blockplay2-cn.svg)

</details>

### Octolines

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `octolines`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen octolines
```

![octolines](samples/octolines.svg)

Changing the colors can be easily done like this:
```
konstrukt gen octolines --color1 "#2c3e50" --color2 "#bdc3c7" --color3 "#b8e994"  --filename octolines-cn.svg
```

![octolines CN color inspiration](samples/octolines-cn.svg)

</details>

### Qbert

The pattern is inspired by the game [Q-bert](https://en.wikipedia.org/wiki/Q*bert) which is a variattion of [Rhombille Tiling](https://en.wikipedia.org/wiki/Rhombille_tiling).

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen qbert
```

![qbert](samples/qbert.svg)

Changing the colors can be easily done like this:
```
konstrukt gen qbert --color1 "#1abc9c" --color2 "#f39c12" --color3 "#079992" --filename qbert-ca.svg
```

![qbert CA color inspiration](samples/qbert-ca.svg)

</details>

### Hourglass

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `hourglass`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen hourglass
```

![hourglass](samples/hourglass.svg)

Changing the colors can be easily done like this:
```
konstrukt gen hourglass --color1 "#e55039" --color2 "#fa983a" --color3 "#fad390"  --filename hourglass-cn.svg
```

![hourglass CN color inspiration](samples/hourglass-cn.svg)

</details>

### Waves

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `waves`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen waves
```

![waves](samples/waves.svg)

Changing the colors can be easily done like this:
```
konstrukt gen waves --color1 "#e55039" --color2 "#ecf0f1" --color3 "#fad390"  --filename waves-ca.svg
```

![waves CA color inspiration](samples/waves-ca.svg)

</details>

### Riviera

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `riviera`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen riviera
```

![riviera](samples/riviera.svg)

Changing the colors can be easily done like this:
```
konstrukt gen riviera --color1 "#27ae60" --color2 "#2980b9" --color3 "#f1c40f"  --filename riviera-ca.svg
```

![riviera CA color inspiration](samples/riviera-ca.svg)

</details>

### Nolock

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `nolock`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen nolock
```

![nolock](samples/nolock.svg)

Changing the colors can be easily done like this:
```
konstrukt gen nolock --color1 "#6a89cc" --color2 "#f6b93b" --color3 "#9b59b6"  --filename nolock-ca.svg
```

![nolock CA color inspiration](samples/nolock-ca.svg)

</details>

### Pantheon

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `pantheon`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen pantheon
```

![pantheon](samples/pantheon.svg)

Changing the colors can be easily done like this:
```
konstrukt gen pantheon --color1 "#bdc3c7" --color2 "#f6b93b" --color3 "#6a89cc"  --filename pantheon-au.svg
```

![pantheon CA color inspiration](samples/pantheon-au.svg)

</details>

### Hex22

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `hex22`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen hex22
```

![hex22](samples/hex22.svg)

Changing the colors can be easily done like this:
```
konstrukt gen hex22 --color1 "#e74c3c" --color2 "#1e3799" --color3 "#82ccdd" --color4 "#ecf0f1"  --filename hex22-cv.svg
```

![hex22 CV color inspiration](samples/hex22-cv.svg)

</details>

### Euphonic

The pattern is inspired by `johal_geometrics` work title `Euphonic Colour No. 8` and implemented by [MarcelMues](https://github.com/MarcelMue).

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen euphonic
```

![euphonic](samples/euphonic.svg)

Changing the colors can be easily done like this:
```
konstrukt gen euphonic --color1 "#2980b9" --color2 "#2c3e50" --color3 "#e58e26" --color4 "#d35400" --color5 "#f6b93b" --color6 "#0c2461" --filename euphonic-cv.svg
```

![euphonic CV color inspiration](samples/euphonic-cv.svg)

</details>

### Woozoo

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `woozoo`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen woozoo
```

![woozoo](samples/woozoo.svg)

Changing the colors can be easily done like this:
```
konstrukt gen woozoo --color1 "#079992" --color2 "#ecf0f1" --color3 "#0c2461" --filename woozoo-cv.svg
```

![woozoo CV color inspiration](samples/woozoo-cv.svg)

</details>

### Triangles

The pattern and default implementation is [MarcelMues](https://github.com/MarcelMue) work titled `triangles`.

<details>
<summary>Usage examples</summary>

The default can be generated with:
```
konstrukt gen triangles
```

![triangles](samples/triangles.svg)

Changing the colors can be easily done like this:
```
konstrukt gen triangles --color1 "#079992" --color2 "#ecf0f1" --color3 "#0c2461" --filename triangles-ca.svg
```

![triangles CA color inspiration](samples/triangles-ca.svg)

</details>

## Guidelines

Additions to this project should follow these guidelines:
- Patterns should be repeatable and scaleable.
- Patterns should be constant across the generated SVG (e.g. no convex pattern).
- All patterns should be made up of low count polygons (no curves or similar).
- Accurate citation should be given if a pattern was inspired by an artwork or artist.
