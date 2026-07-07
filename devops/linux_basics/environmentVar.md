ENVIRONMENT VARIABLES 

Environment variables store configuration and context about the shell and operating system environment that programs can use while running.

its a key value pair

keys are in capital letter

KEY=value

to know the env variables we run commands like
->env

or

printenv

or

export


so these environment variables are made coz operating system uses it to know about the environment

-- we can create our own environment var---

to create env var that is avialable for the whole env -

we do -

". export KEY=VALUE. " [ this sets env for that session only .. not permanent]

TO REMOVE IT use -> unset

TO SET THE ENV VARIABLES permanently we edit shell speicific configuration file
 suppose i am using bash shell

 then have to edit 

 ./bashrc file (vi bashrc) in the file export KEY=VALUE
 and after editing the file do -> source .bashrc

 so lemme tell you -> .bashrc thing persist env varibles user specific
 not system wide

 --ENV VAR SYSTEM WIDE--

 system wide file -> /etc/environment this file contains a global var PATH

 PATH - list of directories to executable files , seperated by : 
 it tells the shell, which directories to search in for the executable , in response to our executable command

 --add a custom program--

 create a file - write any script in it
 give all user executable permission

 then in the user specific bashrc file - > refer to path var
 PATH=$PATH:$USER