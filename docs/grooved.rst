.. _grooved(1):

grooved
=======

SYNOPSIS
--------

.. program:: grooved

**grooved [options]**

DESCRIPTION
-----------

**grooved** is a stupidly simple music player that runs as a daemon instead of
showing a fancy GUI. I'm kinda lazy, so instead or re-implementing some of the
features that make up a music player, I just used stuff someone else already
built. For example, grooved uses mpv to reproduce audio instead of implementing
its own half-working audio decoder and player, and piggybacks on tools like
beets instead of implementing yet another music database. Life gets so much
better once you let other people do the hard work.

OPTIONS
-------

.. option:: -c, --config=<file>

Specify the configuration file. (default: `~/.config/grooved/config.ini`)

.. optio