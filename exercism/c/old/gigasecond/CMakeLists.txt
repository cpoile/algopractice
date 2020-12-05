cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME gigasecond)

# set the project name
project(gigasecond)

set(C_STANDARD 11)

# add the executable
add_executable(gigasecond src/${NAME}.h src/${NAME}.c src/main.c)
