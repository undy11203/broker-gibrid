cmake_minimum_required(VERSION 3.10)

project(core_broker LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find Boost libraries for networking and asynchronous operations
find_package(Boost 1.70 REQUIRED COMPONENTS system thread)

if(Boost_FOUND)
    include_directories(${Boost_INCLUDE_DIRS})
else()
    message(FATAL_ERROR "Boost not found")
endif()

add_executable(core_broker main.cpp)

target_link_libraries(core_broker PRIVATE ${Boost_LIBRARIES})
