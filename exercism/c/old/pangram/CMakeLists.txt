cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME pangram)

# set the project name
project(pangram)

set(C_STANDARD 11)

# add the executable
add_executable(pangram src/${NAME}.h src/${NAME}.c src/main.c)
