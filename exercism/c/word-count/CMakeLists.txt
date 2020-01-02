cmake_minimum_required(VERSION 3.15)

# put the project name here
set(NAME word_count)

# set the project name
project(word_count)

set(C_STANDARD 11)

# add the executable
add_executable(word_count src/${NAME}.h src/${NAME}.c src/main.c)
