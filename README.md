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


# file structure

    init/bindata.go
        contains a compiled version of runner (goos=windows, x86_64)
    runner is litterally just a hello world at this moment
    popup is just the joke popup, straight forward
    