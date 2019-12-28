cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME resistor_color_duo)

# set the project name
project(resistor-color-duo)

set(C_STANDARD 11)

# add the executable
add_executable(resistor_color_duo src/${NAME}.h src/${NAME}.c src/main.c)
