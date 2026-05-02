set(SRC_LIST
    Capsule.cc
    CliOptionParser.cc
    FdWatcher.cc
    Http3Connection.cc
    MasqueClient.cc
    InnerLayer.cc
    IPAllocator.cc
    QuicConnection.cc
    TimerWatcher.cc
    TransportLayer.cc
    OuterLayer.cc
    tls_client_context_quictls.cc
    tls_client_session_quictls.cc
    UdpLayer.cc
    PvdClient.cc
    main.cc
  )

set (local_path $ENV{LOCAL_LIB_PATH})
include_directories( ${local_path}/include/)

find_library(CRYPTO_QUICTLS_LIB ngtcp2_crypto_quictls paths ${local_path}/lib)

if(CRYPTO_QUICTLS_LIB)
    message(STATUS "Found ngtcp2_crypto_quictls: ${CRYPTO_QUICTLS_LIB}")
else()
    message(FATAL_ERROR "Could not find ngtcp2_crypto_quictls")
endif()


find_library(EV_LIB ev paths ${local_path}/lib)

if(EV_LIB)
    message(STATUS "Found libev: ${EV_LIB}")
else()
    message(FATAL_ERROR "Could not find libev")
endif()

find_library(JANSSON_LIB jansson paths ${local_path}/lib)

if(JANSSON_LIB)
    message(STATUS "Found libjansson: ${JANSSON_LIB}")
else()
    message(FATAL_ERROR "Could not find libjansson")
endif()

set(APP http3client)
list(APPEND EXTRA_LIBS ${EV_LIB} ssl crypto ngtcp2_crypto_quictls ngtcp2 nghttp3 workers ${JANSSON_LIB})
list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/workers")
list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/comm")

add_executable(${APP} ${SRC_LIST}  $<TARGET_OBJECTS:comm>)


target_link_libraries(${APP} ${EXTRA_LIBS})
target_include_directories(${APP} PUBLIC
                         "${PROJECT_BINARY_DIR}"
                         ${EXTRA_INCLUDES}
                         )