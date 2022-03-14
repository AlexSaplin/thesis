cmake_minimum_required(VERSION 3.14)

include(FetchContent)

set(CMAKE_CXX_STANDARD 17)

project(selachii)

#
# Non-FetchContent dependencies
#

#

include_directories("/usr/local/include")
link_directories(/usr/local/lib)
include_directories("/usr/include")
link_directories("/usr/lib")


# threads
set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)

# openssl
if(LINUX)
    include_directories("/usr/local/opt/openssl/include")
    link_directories("/usr/local/opt/openssl/lib")
    find_package(OpenSSL REQUIRED)
ELSEIF(APPLE)
    SET(OPENSSL_ROOT_DIR /usr/local/opt/openssl/)
endif()

# Boost 
set(Boost_USE_STATIC_LIBS        ON) # only find static libs
set(Boost_USE_MULTITHREADED      ON)
set(Boost_USE_STATIC_RUNTIME    OFF)
find_package(Boost COMPONENTS date_time filesystem system)
include_directories(${Boost_INCLUDE_DIRS})

#  cpprestsdk
set (CMAKE_PREFIX_PATH /usr/lib/x86_64-linux-gnu/cmake/)
find_package(cpprestsdk REQUIRED)

# ONNX
include_directories("/code/onnxruntime/include/onnxruntime/core/session")
include_directories("/code/onnxruntime/include/onnxruntime/core/providers/cpu")
include_directories("/code/onnxruntime/include/onnxruntime/core/providers/cuda")

link_directories("/code/onnxruntime/build/Linux/Release")


# grpc
set (CMAKE_PREFIX_PATH /usr/local/lib/cmake/protobuf/)
find_package(protobuf REQUIRED)

# grpc
set (CMAKE_PREFIX_PATH /usr/local/lib/cmake/grpc/)
find_package(gRPC REQUIRED)

# spdlog
if (APPLE)
    set (CMAKE_PREFIX_PATH /usr/local/Cellar/spdlog/1.6.0/lib/cmake/spdlog/)
    find_package(spdlog REQUIRED)
endif()

# cxxopts
set (CMAKE_PREFIX_PATH /usr/local/lib/cmake/cxxopts/)
if(LINUX)
    find_package(cxxopts REQUIRED)
endif()

include_directories(${PROJECT_SOURCE_DIR}/src)

file(GLOB SOURCES "src/*.cpp" "src/*/*.cpp" "src/*.cc" "src/*/*.cc" "src/*.h" "src/*/*.h")

set(EXECUTABLE_OUTPUT_PATH bin)

add_executable(selachii ${SOURCES})

target_link_libraries(selachii Threads::Threads)
target_link_libraries(selachii ${Boost_FILESYSTEM_LIBRARY})
target_link_libraries(selachii ${Boost_SYSTEM_LIBRARY})
target_link_libraries(selachii protobuf)
target_link_libraries(selachii gRPC::grpc++)
target_link_libraries(selachii onnxruntime)
target_link_libraries(selachii spdlog)
target_link_libraries(selachii cpprestsdk::cpprest)
if(LINUX)
    target_link_libraries(selachii cxxopts::cxxopts)
endif()
