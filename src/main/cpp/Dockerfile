FROM gcc:9.3.0

ENV CMAKE_RELEASE_VERSION 3.17.3
ENV CMAKE_OS Linux
ENV CMAKE_ARCHITECTURE x86_64

RUN wget -q -O cmake-linux.sh \
	https://github.com/Kitware/CMake/releases/download/v${CMAKE_RELEASE_VERSION}/cmake-${CMAKE_RELEASE_VERSION}-${CMAKE_OS}-${CMAKE_ARCHITECTURE}.sh \
	&& sh cmake-linux.sh -- --skip-license --prefix=/usr/local \
	&& rm cmake-linux.sh