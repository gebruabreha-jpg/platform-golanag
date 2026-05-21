set(SRC_LIST
    Worker.cc
    )

set (local_path $ENV{LOCAL_LIB_PATH})

find_library(EV_LIB ev paths ${local_path}/lib)
include_directories( ${local_path}/include/)

message(STATUS "Your local lib path is ${local_path}")

if(EV_LIB)
    message(STATUS "Found libev: ${EV_LIB}")
else()
    message(FATAL_ERROR "Could not find libev")
endif()

set(LIB workers)
list(APPEND EXTRA_LIBS pthread ${EV_LIB})
add_library(${LIB} ${SRC_LIST})
target_link_libraries(${LIB} ${EXTRA_LIBS})
target_include_directories(${LIB} PUBLIC  ${local_path}/include/)