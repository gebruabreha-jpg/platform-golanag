set(SRC_LIST
Config.cc 
HttpRequest.cc
debug.cc 
FlatBuffer.cc
http_parser.c
network.cc
tls_session_base_quictls.cc 
util.cc 
util_openssl.cc
PM.cc
)

set(LIB comm)
add_library(${LIB} ${SRC_LIST})
target_include_directories(${LIB} PUBLIC
                         ${EXTRA_INCLUDES}
                         )