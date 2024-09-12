# Version Control System

## Git
Git adalah sebuah version control system yang populer digunakan untuk mengelola kode - kode yang dibuat oleh programmer. Git juga digunakan untuk mempermudah komunikasi dan kerja sama tim dalam melakukan project.

## Git Basic Commands & Branch Management
Git memiliki commands untuk melakukan pengelolaan dari repository itu sendiri salah satu contohnya ada
- git add, berguna untuk staging perubahan yang dilakukan
- git commit, berguna untuk commit perubahan dengan menambahkan message
- git branch, berguna untuk melihat branch apa saja yang ada di dalam repository kita
- git checkout, berguna untuk pindah ke branch yang kita inginkan, atau pindah ke commit yang sudah dilakukan sebelumnya 
- git merge, berguna untuk menggabungkan 2 branches
- git push, berguna untuk mengirim perubahan yang sudah di commit ke remote server
- git pull, berguna untuk mengambil perubahan yang sudah di push dari remote server

## Gitflow & Pull Request
Gitflow adalah alur kerja yang membantu dalam pengelolaan branch yang lebih terstruktur, contohnya jika ingin menambah fitur, programmer harus membuat branch baru dengan nama fitur yang mau dibuat dan setelah jadi dilakukan Pull Request dari origin/main untuk dilakukan code review oleh team leader atau quality assurance agar tidak ada bug atau error yang masuk ke dalam origin/main
