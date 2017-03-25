# godamnit
Replace ascii characters with identical unicode ones - inspired by [reinderien/mimic](https://github.com/reinderien/mimic) as well as the old [Greek question mark](https://twitter.com/peterritchie/status/534011965132120064) joke mentioned in mimic's README

Unlike mimic, godamnit trys to stick to unicode characters that look as close as possible to what they're replacing, but of course the practical joke possiblities are still limited by most IDEs

```html
usage: godamnit [<flags>] [<input file>] [<output file>]

By default takes text from stdin, encodes it and outputs it back to stdout. Can
instead take file paths with the args below.

Flags:
      --help    Show context-sensitive help (also try --help-long and
                --help-man).
  -d, --decode  Decode instead of encode
  -t, --test    Dumps comma seperated unicode code points to stdout

Args:
  [<input file>]   Input file path
  [<output file>]  Output file path
```

You can use either stdin & stdout:
```
cat test.txt | godamnit > out.txt
```

Or provide file paths for the in and out files:
```
godamnit.exe test.txt out.txt
```