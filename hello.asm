section .data
hello db "Hello, GeeFriday", 0

section  .text
global _start

_start:
mov eax, 4
mov ebx, 1
mov ecx, hello
mov edx, 17
int 0x80

mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

mov eax, 1
int 0x80

section .data
newline db 10

;assemble the code using nasm
;;$ nasm -f elf64 hello.asm -o hello.o

;;Link the object file to create an executable:
;$ ld hello -o hello ;;uses the gnu linker to create an executable "hello" from "hello.o"
;;Run the program $ ./hello

