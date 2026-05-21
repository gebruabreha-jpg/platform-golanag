CMAKE_MINIMUM_REQUIRED(VERSION 3.1)

PROJECT(gladiator)

# specify the C++ standard
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)



set(DURING_DEV_PHASE False)

if(DURING_DEV_PHASE)
    message(STATUS "Build ${PROJECT_NAME} for debug.")
    ADD_DEFINITIONS(-DDEBUG -Wall -Werror)
    # Add the compiler flags
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fsanitize=address -g")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fsanitize=address -g -Wunused-function -Wno-deprecated-enum-enum-conversion")
    # Add the linker flags
    set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -fsanitize=address")
else()
    message(STATUS "Build ${PROJECT_NAME} for release.")
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS}")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wunused-function -Wno-deprecated-enum-enum-conversion")
    ADD_DEFINITIONS(-O3 -g)
endif()

ADD_SUBDIRECTORY(src/workers)
ADD_SUBDIRECTORY(src/comm)
ADD_SUBDIRECTORY(src/masque_client)
ADD_SUBDIRECTORY(src/http3_server)
ADD_SUBDIRECTORY(utest)

