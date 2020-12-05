cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME resistor_color)

# set the project name
project(resistor-color)

set(C_STANDARD 11)

# add the executable
add_executable(resistor_color src/${NAME}.h src/${NAME}.c src/main.c)
