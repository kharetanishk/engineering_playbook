INPUT AND OUTPUT

In linux , output from one command can be the input for another command

 for e,g

 cat /var/log/syslog |-> this command produces an output 

 now the syntax for passing the output as an input to the next command is 

 "|" pipe-> pipes the output of the previous command as the input of another command
 
 e.g cat /var/log/systog | less

 here less command is used to show large contents in pages and user friendly manner

 grep -> global search for regular expression and print out , it searches for a particular pattern of characters and  displays all lines that contains that pattern 

 e.g   history | grep sudo

meaning -> all the output that history command produces pipe it to  the next command
and the command says from this output search for chars sudo 

for muliple words - grep "sudo chmod" -use " "

--REDIRECTION--

" > " -> Redirection symbol , it is used to  to redirect the output of the previous command to a file 

e.g history | grep sudo > sudo.commands.txt

this will grep all the sudo commands (used ) into a file named sudo.command.txt

to append the content we use ->. " >> "

running multiple commands 

echo ; sleep 3 ; echo "hey i am tanishk khare"
";" this is seperation of concern