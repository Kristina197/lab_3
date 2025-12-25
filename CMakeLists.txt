cmake_minimum_required(VERSION 3.28)
project(nstu_lab3_3)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} --coverage -O0 -g")

# Объявление путей до структур данных
set(STRUCTS_SOURCES
        src/Array/Array.cpp
        src/Stack/Stack.cpp
        src/Queue/Queue.cpp
        src/SLL/SinglyLinkedList.cpp
        src/DLL/DoublyLinkedList.cpp
        src/FBT/FullBinaryTree.cpp
        src/ChainingHash/ChainingHash.cpp
        src/OpenAddrHash/OpenAddrHash.cpp
        src/redirect/OutputRedirect.cpp
)

# Использование линтера
find_program(CLANG_TIDY_EXE NAMES clang-tidy)
if(CLANG_TIDY_EXE)
    set(CMAKE_CXX_CLANG_TIDY
            ${CLANG_TIDY_EXE}
            -header-filter=.*
            -extra-arg=-Wno-unknown-warning-option
    )

    add_custom_target(clang-tidy-all
            COMMAND ${CLANG_TIDY_EXE}
            ${CMAKE_CURRENT_SOURCE_DIR}/src/**/*
            ${CMAKE_CURRENT_SOURCE_DIR}/tests/**/*
            -header-filter=.*
            --config-file=
            WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    )
endif()


# Boost Test
find_package(Boost REQUIRED COMPONENTS unit_test_framework)
add_executable(BoostTest
        tests/BoostTest.cpp
        ${STRUCTS_SOURCES}
)
target_link_libraries(BoostTest Boost::unit_test_framework)
target_compile_options(BoostTest PRIVATE --coverage -O0 -g)
target_link_options(BoostTest PRIVATE --coverage)

add_custom_command(TARGET BoostTest POST_BUILD
        COMMAND $<TARGET_FILE:BoostTest>
        COMMAND bash ${CMAKE_SOURCE_DIR}/run_coverage.sh boost
        WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
)


# Google Test
add_executable(GoogleTest
        tests/GoogleTest.cpp
        ${STRUCTS_SOURCES}
)
target_link_libraries(GoogleTest gtest gmock pthread)
target_compile_options(GoogleTest PRIVATE --coverage -O0 -g)
target_link_options(GoogleTest PRIVATE --coverage)

add_custom_command(TARGET GoogleTest PRE_BUILD
        COMMAND cppcheck --enable=all ${CMAKE_CURRENT_SOURCE_DIR}
)
add_custom_command(TARGET GoogleTest POST_BUILD
        COMMAND bash ${CMAKE_SOURCE_DIR}/run_coverage.sh gtest
        WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
)


# Catch2 Test
find_package(Catch2 REQUIRED)
add_executable(Catch2Test
        tests/Catch2Test.cpp
        ${STRUCTS_SOURCES}
)
target_link_libraries(Catch2Test Catch2::Catch2WithMain)
target_compile_options(Catch2Test PRIVATE --coverage -O0 -g)
target_link_options(Catch2Test PRIVATE --coverage)

add_custom_command(TARGET Catch2Test PRE_BUILD
        COMMAND cppcheck --enable=all ${CMAKE_CURRENT_SOURCE_DIR}
)
add_custom_command(TARGET Catch2Test POST_BUILD
        COMMAND bash ${CMAKE_SOURCE_DIR}/run_coverage.sh catch2
        WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
)