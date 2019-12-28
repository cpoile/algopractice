cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME armstrong_numbers)

# set the project name
project(${NAME})

set(C_STANDARD 11)

# add the executable
add_executable(${NAME} src/${NAME}.h src/${NAME}.c src/main.c)

