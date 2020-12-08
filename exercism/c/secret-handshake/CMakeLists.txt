cmake_minimum_required(VERSION 3.17)
project(secret_handshake C)

set(CMAKE_C_STANDARD 11)

include_directories(src)
include_directories(test/vendor)

add_executable(secret_handshake
        src/secret_handshake.c
        src/secret_handshake.h
        test/vendor/unity.c
        test/vendor/unity.h
        test/vendor/unity_internals.h
        test/test_secret_handshake.c)
