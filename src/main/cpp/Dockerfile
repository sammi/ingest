FROM grpc/cxx:1.10

ENV CMAKE_OS Linux
ENV CMAKE_ARCHITECTURE x86_64
ENV CMAKE_RELEASE_VERSION 3.17.3
ENV CMAKE_INSTALL_PREFIX /usr/local

RUN apt-get update && apt install -y wget

RUN wget -q -O cmake-linux.sh \
    https://github.com/Kitware/CMake/releases/download/v${CMAKE_RELEASE_VERSION}/cmake-${CMAKE_RELEASE_VERSION}-${CMAKE_OS}-${CMAKE_ARCHITECTURE}.sh \
    && sh cmake-linux.sh -- --skip-license --prefix=${CMAKE_INSTALL_PREFIX} \
    && rm cmake-linux.sh

RUN mkdir -p /server

ADD . /server/

WORKDIR /server

RUN mkdir build \
	&& cd build \
	&& cmake .. \
	&& make \
	&& cp ingest /usr/local/bin \
	&& rm -rf /server/build

ENV GRPC_SERVER_PORT 50051

EXPOSE ${GRPC_SERVER_PORT}

ENTRYPOINT ["sh", "-c", "ingest ${GRPC_SERVER_PORT}"]