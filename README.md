# Goprompt

Goprompt is a VCS prompt for Git (and Mercurial).

This is my first Go application. I used it to learn about Interfaces.

Goprompt has no external dependencies other than the standard library

## How to build and install
```
go install
```


Update your bash/zsh shell prompt:
```bash
PROMPT='(goprompt -f t:bnm)'
```

By default it prints out a prompt like this:
```bash
[git:master?] # if there are untracked files
```

And:
```bash
[git:master+] # if there are changes not staged for commit
```
