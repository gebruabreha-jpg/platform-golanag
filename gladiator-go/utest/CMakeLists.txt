
set(SRC_LIST
    main.cc
    test_flatbuffer.cc
    test_ipv4_pool.cc
    )

find_package(GTest REQUIRED)

list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/workers")
list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/comm")
list(APPEND EXTRA_INCLUDES "${PROJECT_SOURCE_DIR}/src/masque_client")
list(APPEND EXTRA_LIBS ${EV_LIB} ssl crypto ngtcp2_crypto_quictls ngtcp2 nghttp3 workers)
set(APP utest)
add_executable(${APP} ${SRC_LIST} $<TARGET_OBJECTS:comm>  ${PROJECT_SOURCE_DIR}/src/masque_client/IPAllocator.cc)
target_link_libraries(${APP} PRIVATE ${GTEST_LIBRARIES} ${EXTRA_LIBS})
target_include_directories(${APP} PUBLIC
                         "${PROJECT_BINARY_DIR}"
                         ${EXTRA_INCLUDES}
                         )
