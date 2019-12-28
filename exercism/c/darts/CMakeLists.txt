cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME darts)

# set the project name
project(darts)

set(C_STANDARD 11)

# add the executable
add_executable(darts src/${NAME}.h src/${NAME}.c src/main.c)

