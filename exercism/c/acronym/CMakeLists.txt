cmake_minimum_required(VERSION 3.17)
project(acronym C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(acronym
        src/acronym.c
        src/acronym.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_acronym.c)
