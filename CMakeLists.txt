cmake_minimum_required(VERSION 3.17.0)

set(APP_NAME          	"ingest")
set(APP_VERSION   		"0.1.0")

project(${APP_NAME} VERSION ${APP_VERSION} LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 	17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

add_executable(ingest ingest.cc)