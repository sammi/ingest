cmake_minimum_required(VERSION 3.17.0)

set(APP_NAME          	"ingest")
set(APP_VERSION   		"0.1.0")

project(${APP_NAME} VERSION ${APP_VERSION} LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 	17)
set(CMAKE_CXX_STANDARD_REQUIRED True)


list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}/cmake")
list(APPEND CMAKE_PREFIX_PATH "/opt/grpc" "/opt/protobuf")
find_package(Protobuf REQUIRED)
find_package(gRPC REQUIRED)


set(PROTOS ${CMAKE_CURRENT_SOURCE_DIR}/../../../protos/contract.proto)
set(PROTO_SRC_DIR ${CMAKE_CURRENT_BINARY_DIR}/protos-src)

file(MAKE_DIRECTORY ${PROTO_SRC_DIR})
include_directories(${PROTO_SRC_DIR})

protobuf_generate_cpp(PROTO_SRCS PROTO_HDRS ${PROTO_SRC_DIR} ${PROTOS})

grpc_generate_cpp(GRPC_SRCS GRPC_HDRS ${PROTO_SRC_DIR} ${PROTOS})

add_executable(ingest
 	${CMAKE_CURRENT_SOURCE_DIR}/ingest.cc
 	${PROTO_SRCS}
 	${GRPC_SRCS}
)

target_link_libraries(ingest
    gRPC::grpc++_reflection
    protobuf::libprotobuf
)

install(TARGETS ingest DESTINATION bin)
