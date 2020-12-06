cmake_minimum_required(VERSION 3.17)
project(matching_brackets C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(matching_brackets
        src/matching_brackets.c
        src/matching_brackets.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_matching_brackets.c)
