FROM golang

RUN apt-get update -qq \
    && apt-get install -yq cmake \
                           fceux \
                           gcc \
                           libboost-all-dev \
                           libjpeg-dev \
                           libjpeg62-turbo-dev \
                           libsdl2-dev \
                           make \
                           unzip \
                           wget \
                           zlib1g-dev

RUN mkdir src