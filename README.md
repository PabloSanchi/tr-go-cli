# TR-GO-CLI
Go alternative for https://codingchallenges.substack.com/p/coding-challenge-43-tr

# Step 1
In this step your goal is to support translating from one character to another, when reading from the standard input. For this your program should start up and wait for a line of input from the user. It will then output the line having made the substitution specified, for example:

```bash
% cctr c C
coding challenges
Coding Challenges
```

Type CTRL-D to send the EOF and terminate cctr.

Try it with the following
```bash
go build && echo coding challenges | ./cctr c C
```

# Step 2
In this step your goal is to support translation of a range of characters. As is common ranges are specified by the start and end separated by a hyphen. For example the upper case letters are A-Z and the lower case letters a-z. When you’ve implemented this it should look something like this:

```bash
% head -n3 test.txt
The Project Gutenberg eBook of The Art of War
    
This ebook is for the use of anyone anywhere in the United States and
 
% head -n3 test.txt | cctr A-Z a-z
the project gutenberg ebook of the art of war
    
this ebook is for the use of anyone anywhere in the united states and
```

Where we can see the first three lines as published by Project Gutenberg, then the same translated to lowercase in the second example.

Try it with the following
```bash
go build && head -n3 test.txt | ./cctr A-Z a-z
```