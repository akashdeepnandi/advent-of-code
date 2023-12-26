CC = g++
CFLAGS = -std=c++11 -Wall

# Target executable name
TARGET = main

# Source files
SRCS = main.cpp helper.cpp

# Object files
OBJS = $(SRCS:.cpp=.o)

# Makefile rules
all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) $(OBJS) -o $(TARGET) -lssl -lcrypto

%.o: %.cpp
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f $(OBJS) $(TARGET)
