cmake_minimum_required(VERSION 3.17)
project(gigasecond C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(gigasecond
        src/gigasecond.c
        src/gigasecond.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_gigasecond.c)
