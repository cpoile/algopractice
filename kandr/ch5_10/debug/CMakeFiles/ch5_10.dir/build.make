# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.15

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/local/Cellar/cmake/3.15.5/bin/cmake

# The command to remove a file.
RM = /usr/local/Cellar/cmake/3.15.5/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug

# Include any dependencies generated for this target.
include CMakeFiles/ch5_10.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/ch5_10.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/ch5_10.dir/flags.make

CMakeFiles/ch5_10.dir/main.c.o: CMakeFiles/ch5_10.dir/flags.make
CMakeFiles/ch5_10.dir/main.c.o: ../main.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/ch5_10.dir/main.c.o"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/ch5_10.dir/main.c.o   -c /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/main.c

CMakeFiles/ch5_10.dir/main.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/ch5_10.dir/main.c.i"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/main.c > CMakeFiles/ch5_10.dir/main.c.i

CMakeFiles/ch5_10.dir/main.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/ch5_10.dir/main.c.s"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/main.c -o CMakeFiles/ch5_10.dir/main.c.s

# Object files for target ch5_10
ch5_10_OBJECTS = \
"CMakeFiles/ch5_10.dir/main.c.o"

# External object files for target ch5_10
ch5_10_EXTERNAL_OBJECTS =

ch5_10: CMakeFiles/ch5_10.dir/main.c.o
ch5_10: CMakeFiles/ch5_10.dir/build.make
ch5_10: CMakeFiles/ch5_10.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable ch5_10"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/ch5_10.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/ch5_10.dir/build: ch5_10

.PHONY : CMakeFiles/ch5_10.dir/build

CMakeFiles/ch5_10.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/ch5_10.dir/cmake_clean.cmake
.PHONY : CMakeFiles/ch5_10.dir/clean

CMakeFiles/ch5_10.dir/depend:
	cd /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10 /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10 /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug /Users/chris/go/src/github.com/cpoile/practice/kandr/ch5_10/debug/CMakeFiles/ch5_10.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/ch5_10.dir/depend

