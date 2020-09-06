
grooved
=======

**grooved** is a stupidly simple music player that runs as a daemon instead of
showing a fancy GUI. I'm kinda lazy, so instead or re-implementing some of the
features that make up a music player, I just used stuff someone else already
built. For example, grooved uses mpv_ to reproduce audio instead of implementing
its own half-working audio decoder and player, and piggybacks on tools like
beets_ instead of implementing yet another music database. Life gets so much
better once you let other people do the hard work.

grooved's operation is very simple: there's a single tracklist (or track queue,
or call-it-however-you-want) which can be filled manually by the user by adding
any kind of track (local files, network streams, YouTube videos, ..., you can
throw pretty much anything at it), which will be played in order. Once the