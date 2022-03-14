FROM nvidia/cuda:10.1-cudnn7-devel


# Install dependencies
WORKDIR /code

# deps
RUN apt-get update &&\
    apt-get install --no-install-recommends -y sudo git bash gcc g++ cmake make \
    libssl-dev libboost-all-dev libcpprest-dev && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENV PATH /usr/local/nvidia/bin:/usr/local/cuda/bin:/code/cmake-3.14.3-Linux-x86_64/bin:/opt/miniconda/bin:${PATH}
ENV LD_LIBRARY_PATH /opt/miniconda/lib:$LD_LIBRARY_PATH

# ONNX
ARG ONNXRUNTIME_REPO=https://github.com/Microsoft/onnxruntime
ARG ONNXRUNTIME_BRANCH=master

RUN git clone --single-branch --branch ${ONNXRUNTIME_BRANCH} --recursive ${ONNXRUNTIME_REPO} onnxruntime &&\
    /bin/sh onnxruntime/dockerfiles/scripts/install_common_deps.sh &&\
    cd onnxruntime &&\
    /bin/sh ./build.sh --cuda_home /usr/local/cuda --cudnn_home /usr/lib/x86_64-linux-gnu/ \
        --use_cuda --config Release --update --build --build_shared_lib --parallel \
        --cmake_extra_defines ONNXRUNTIME_VERSION=$(cat ./VERSION_NUMBER)

# gRPC
RUN git clone https://github.com/grpc/grpc &&\
    cd grpc && \
    git checkout tags/v1.29.1 && \
    git submodule init && \
    git submodule update --recursive && \
    cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF DCMAKE_INSTALL_PREFIX=/code/grpc . && make install

# spdlog
RUN git clone https://github.com/gabime/spdlog.git && cd spdlog && mkdir build && cd build && \
    cmake .. -DSPDLOG_INSTALL=ON && make -j && make install

# cxxopts
RUN git clone https://github.com/jarro2783/cxxopts && cd cxxopts && \
    cmake . && make && make install
