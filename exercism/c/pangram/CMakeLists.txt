cmake_minimum_required(VERSION 3.17)
project(pangram C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(pangram
        src/pangram.c
        src/pangram.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_pangram.c)
