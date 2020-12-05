cmake_minimum_required(VERSION 3.17)
project(word_count C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(word_count
        src/word_count.c
        src/word_count.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_word_count.c)
