cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME isogram)

# set the project name
project(isogram)

set(C_STANDARD 11)

# add the executable
add_executable(isogram src/${NAME}.h src/${NAME}.c src/main.c)
