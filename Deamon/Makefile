CFLAGS=-Wall -pedantic -std=gnu99

all: Deamon
Deamon:
	gcc $(CFLAGS) 'pkg-config --cflags --libs libnotify' watch_file.c -o build/watch_filed
