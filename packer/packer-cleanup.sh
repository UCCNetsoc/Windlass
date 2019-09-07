apt-get clean

DIRS_TO_TRIM="/usr/share/man
/var/cache/apt
/var/lib/apt/lists
/usr/share/locale
/var/log
/usr/share/info
"

for DIR in $DIRS_TO_TRIM; do
  rm -r "$rootfsDir/$DIR"/*
done