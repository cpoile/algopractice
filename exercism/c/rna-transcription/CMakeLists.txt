cmake_minimum_required(VERSION 3.17)
project(rna_transcription C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(rna_transcription
        src/rna_transcription.c
        src/rna_transcription.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_rna_transcription.c)
