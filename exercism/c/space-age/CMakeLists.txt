cmake_minimum_required(VERSION 3.17)
project(space_age C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(space_age
        src/space_age.c
        src/space_age.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_space_age.c)
