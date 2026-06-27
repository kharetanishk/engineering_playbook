NOTE - EVERYTHING IN LINUX IS A FILE

pwd - tell which directory are you in !
ls - list folder and files
cd dir_name - change directory name
mkdir - create directory 
touch filename - creates a file
rm filename - remove file 
rm -r dir_name - delete a non empty directory 
rmdir dir_name - delete an empty directory 
cd / - go to root folder 
cd [abosulte path] - will reach us to that exact folder 
ls [absoulte path] - list dir in that path 
mv current_filename new_filename - change the name of the file/folder
cp -r filename_You_want_to_copy filename_you_want_to_give - copy the folder (copying recursively) 
to copy a single file -> cp filenamewanttocopy filenameyougive
tanishk@Anubha:~$ ls -R t2 ( to know what contents inside that dirs)
t2:
t3  t4

t2/t3:
hello.txt  test.txt

t2/t4:
ta.txt
tanishk@Anubha:~$ 

history - tells the history of commands we executed

ctrl + r -> search the commands 
ls -a -> show hidden files in that dir 

cat filename - display file content

uname -a -> show system and kernal information

lscpu - show the configurations of the cpu
lsmem- memory information

adduser username - to add a user from root 
sudo adduser username - add a user as a userprofile 

to switch between user - [ su - username]

sudo su to get back to root user 