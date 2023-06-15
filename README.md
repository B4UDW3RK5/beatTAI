<pre>
▄                          ▄      ▄ ▄ ▄ ▄ ▄    ▄ ▄    ▄ ▄ ▄
▄ ▄ ▄     ▄ ▄     ▄ ▄ ▄    ▄ ▄ ▄      ▄      ▄     ▄    ▄
▄    ▄  ▄     ▄         ▄  ▄          ▄      ▄ ▄ ▄ ▄    ▄
▄    ▄  ▄ ▀ ▀     ▄ ▀ ▀ ▄  ▄          ▄      ▄     ▄    ▄
▄ ▄ ▄     ▄ ▄ ▄     ▄ ▄ ▄    ▄ ▄      ▄      ▄     ▄  ▄ ▄ ▄

 -:[ a new Internet time for turbonerds & superweirdos ]:-

[Abstract]

    In 1998 Swatch introduced <a href="https://en.wikipedia.org/wiki/Swatch_Internet_Time">Swatch Internet Time</a> (or .beat time), displayed as
 @xxx.xx (for example @198.26). It's a cute and fun way to display the time and
 I find it aesthetically very pleasing to look at.

    However, Swatch being based in Switzerland they chose to align .beats to
 Swiss time (UTC+1) which outside of countries in that timezone means very
 little and is impractical.

     I half-jokingly <a href="https://hackers.town/@cat/107951155969547130">posted on the Fediverse</a> that I wanted a new .beat time
 aligned to UTC instead of UTC+1, and was met with some small discussion about
 aligning it to <a href="https://en.wikipedia.org/wiki/International_Atomic_Time">International Atomic Time</a> (TAI) instead and, well, here we are:
 beatTAI or .tai for short.


[Format]

    In line with Swatch .beat time, .tai is a day divided into 1000 and
 represented as :xxx.xx


[Math]

   ((TAIhours * 3600) + (TAIminutes * 60) + TAIseconds) / 86.4
</pre>