users and groups as names , like in a server everyone have different responsibility and access to different roles and permissions 
to isolate each other in terms of their roles and responsibilities , we create groups and assign set of permissions for that group , and ultimately add users to those groups thats al!

so lets begin 

/etc/passwd -> this file stores user account information , everyone can read it but only root user can change the file 

/etc/group -> this file stores the group information

format of information -> username : password : UId :gid :gecos : homedir : shell

username -> user k name
password-> x placeholder 
uID -> unique id for every user ( 0 resrved for root )
gid - group id


--change password for a user

sudo passwd <username>
 

-- to add user

from root - adduser <username>

from user - sudo adduser <username>

-- to create group 

groupadd <groupname> | addgroup <groupname> \ user is adding to a primary group

--to add a user in a group 

sudo usermod -g <groupnam> <usernmae>

--to delete a group

sudo delgroup <groupname>

but as one user can be addded to multiple groups 
one primary other secondary 

--to add a user in multiple secondary group

sudo usermod -G <list of groups e.g admin , clerk> <username>

--to appned newe group to the existing 

sudo usermod -aG <newgroup> <usrname>


--to create user with group

sudo useradd -G <GROUPNAME> <username>
-G -> creates user with multiple secondary groups 

--remove user from a group 

sudo gpasswd -d <username> <groupname>


#GROUPS

Primary group (-g)

A user can have only one primary group.

Example:

sudo usermod -g developers tanishk

Now:

Primary group = developers
Secondary groups (-G)

A user can belong to many secondary groups.

Example:

sudo usermod -G sudo,docker tanishk

This sets the supplementary groups to exactly:

sudo
docker

⚠️ Important: -G replaces the user's current supplementary groups.

If tanishk was previously in:

users
audio
video

they will be removed unless you include them too.

Append secondary groups (-aG)

This is the command you'll use most often.

sudo usermod -aG docker tanishk

This means:

-a = append
-G = supplementary groups

Now docker is added while keeping existing groups.

For example:

Before:

sudo
users

After:

sudo
users
docker