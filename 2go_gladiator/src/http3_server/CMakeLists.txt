set(SRC_LIST
   http.cc  
   server.cc  
   server_base.cc  
   tls_server_context_quictls.cc  
   tls_server_session_quictls.cc
  )

  #ADD_DEFINITIONS(-DHAVE_LINUX_RTNETLINK_H=1)

  set (local_path $ENV{LOCAL_LIB_PATH})

  find_library(CRYPTO_QUICTLS_LIB ngtcp2_crypto_quictls paths ${local_path}/lib)
  include_directories( ${local_path}/include/)
  
  message(STATUS "Your local lib path is ${local_path}")
  
  if(CRYPTO_QUICTLS_LIB)
      message(STATUS "Found ngtcp2_crypto_quictls: ${CRYPTO_QUICTLS_LIB}")
  else()
      message(FATAL_ERROR "Could not find ngtcp2_crypto_quictls")
  endif()


  set (local_path $ENV{LOCAL_LIB_PATH})

  find_library(EV_LIB ev paths ${local_path}/lib)
  include_directories( ${local_path}/include/)
  
  message(STATUS "Your local lib path is ${local_path}")
  
  if(EV_LIB)
      message(STATUS "Found libev: ${EV_LIB}")
  else()
      message(FATAL_ERROR "Could not find libev")
  endif()


# set(APP qtlsserver)
# list(APPEND EXTRA_LIBS ${EV_LIB} ssl crypto ngtcp2_crypto_quictls ngtcp2 nghttp3 workers)
# list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/workers")
# list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/http-parser")
# list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/qtlscomm")

set(APP http3server)
list(APPEND EXTRA_LIBS ${EV_LIB} ssl crypto ngtcp2_crypto_quictls ngtcp2 nghttp3 workers)
list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/workers")
list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/comm")

add_executable(${APP} ${SRC_LIST} $<TARGET_OBJECTS:comm>)
target_link_libraries(${APP} ${EXTRA_LIBS})
target_include_directories(${APP} PUBLIC
                         "${PROJECT_BINARY_DIR}"
                         ${EXTRA_INCLUDES}
                         )

