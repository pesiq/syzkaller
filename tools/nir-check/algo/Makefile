TARGET = main
CC = g++
CFLAGS = -c -Wall

PREF_SRC = ./
PREF_BUILD = ./obj/

SRC = $(wildcard $(PREF_SRC)*.c)
OBJ = $(patsubst $(PREF_SRC)%.c, $(PREF_BUILD)%.o, $(SRC))

$(TARGET) : $(OBJ)
	$(CC) $(SRC) -o $(TARGET)

$(PREF_BUILD)%.o : $(PREF_SRC)%.c
	$(CC) $(CFLAGS) $< -o $@
