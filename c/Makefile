# Makefile for C demo programs

LIB = -lm
CC = gcc
CFLAGS = -O2 -Wall

OBJECTS = toolbox.o

APP1 = randomseq
SRC1 = randomseq.c
OBJ1 = randomseq.o

# Targets

default:
	make $(APP1)


$(APP1): $(OBJ1) $(OBJECTS)
	$(CC) -o $(APP1) $(OBJ1) $(OBJECTS)

clean:
	rm -f *.o $(APP1)


depend: $(OBJECTS:.o=.c)
	gcc -MM $^ > $@

# Inference Rules

%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

# Dependancies

include depend
