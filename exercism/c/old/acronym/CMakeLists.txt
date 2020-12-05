cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME acronym)

# set the project name
project(acronym)

set(C_STANDARD 11)

# add the executable
add_executable(acronym src/${NAME}.h src/${NAME}.c src/main.c)
