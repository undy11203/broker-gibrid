cmake_minimum_required(VERSION 3.10)

project(broker_core LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find Boost libraries for networking and asynchronous operations
find_package(Boost 1.70 REQUIRED COMPONENTS system thread)

if(Boost_FOUND)
    include_directories(${Boost_INCLUDE_DIRS})
else()
    message(FATAL_ERROR "Boost not found")
endif()

# Add source files (placeholder)
set(SOURCES
    broker.cpp
    # Add other source files here
)

add_library(broker_core STATIC \${SOURCES})

target_link_libraries(broker_core PRIVATE \${Boost_LIBRARIES})
