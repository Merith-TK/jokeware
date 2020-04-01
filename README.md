[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/Merith-TK/jokeware) 

This is my joke malware project

how it is inteded to function
    1) run init.exe
    2) init.exe downloads/extracts runtime.exe to `shell:startup` on 
    the host PC,
    3) init.exe runs runtime.exe
    4) runtime.exe downloads popup.exe
    5) runtime.exe starts timer, upon timer end, run popup.exe
        if popup.exe notExist, download
        restart timer

### Known problems
`Makefile` is not functional, at all, will get it functional later
`runner` does not work on windows due to how it makes a syscall or some bullcrap i dont know