cmake_minimum_required(VERSION 3.17)
project(nucleotide_count C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(nucleotide_count
        src/nucleotide_count.c
        src/nucleotide_count.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_nucleotide_count.c)
