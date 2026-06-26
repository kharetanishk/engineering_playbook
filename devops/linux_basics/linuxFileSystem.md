bin - has basics commands like cat ,echo .. are all located in bin dir (system wide)
sbin -  has system relevant commands need super user priviledge 
lib - hold libraries for those bin files
usr - system wide has all 3 dir above 
/usr contains software and files that are part of the operating system or installed through the system's package manager.

Examples:

/usr/bin
/usr/lib
/usr/share
/usr/local
Common directories
/usr/bin      -> User commands (ls, cat, grep, vim)
/usr/lib      -> Shared libraries
/usr/share    -> Documentation, icons, man pages
/usr/local    -> Software installed manually by admin

/opt

/opt is for optional third-party software that is not part of the standard OS.

Example:

/opt/google/chrome
/opt/idea
/opt/postman

Many companies install their applications under /opt.

Example:

/opt/postman/Postman

/ (root directory):

Top-level directory of the entire Linux filesystem.
Every file and directory exists under it.

/home:

Contains home directories of regular users.

/root:

Home directory of the root (administrator) user.

Example:

/
/home/tanishk
/root

Why separate /root from /home?

Even if:

/home

becomes unavailable due to a disk issue, the root user can still log in using:

/root

and repair the system.

That's one of the reasons Linux keeps the root user's home separate.


/etc - has configurations files of differnt applications

/media is used for removable storage devices.

Think:

/media = USBs, CDs, External Drives

When you plug in:

USB drive
External HDD
CD/DVD

HIDDEN FILES -
is primarily used to prevent important data from being accidentally deleted 
are also called dot files 