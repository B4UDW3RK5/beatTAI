```
[Abstract]

    In 1998 Swatch introduced [Swatch Internet Time](https://en.wikipedia.org/wiki/Swatch_Internet_Time) (or .beat time), displayed as
 @xxx.xx (for example @198.26). It's a cute and fun way to display the time and
 I find it aesthetically very pleasing to look at.

    However, Swatch being based in Switzerland they chose to align .beats to
 Swiss time (UTC+1) which outside of countries in that timezone means very
 little and is impractical.

     I half-jokingly [posted on the Fediverse](https://hackers.town/@cat/107951155969547130) that I wanted a new .beat time
 aligned to UTC instead of UTC+1, and was met with some small discussion about
 aligning it to [International Atomic Time](https://en.wikipedia.org/wiki/International_Atomic_Time) (TAI) instead and, well, here we are:
 beatTAI or .tai for short.


[Format]

    In line with Swatch .beat time, .tai is a day divided into 1000 and
 represented as @xxx.xx - though perhaps a character other than @ should be
 used? @ is pretty played out. Answers on a postcard.


[Math]

   ((TAIminutes * 60) + (TAIhours * 3600)) / 86.4
```