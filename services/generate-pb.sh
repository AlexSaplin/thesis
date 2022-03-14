#!/bin/bash
set -xueo pipefail

################
### services ###
################

# ardea
mkdir -p ardea/pkg/grpc/pb/
protoc -I proto/ proto/ardea.proto --go_out=plugins=grpc:ardea/pkg/grpc/pb

# hippo
mkdir -p hippo/pkg/grpc/pb/
protoc -I proto/ proto/hippo.proto --go_out=plugins=grpc:hippo/pkg/grpc/pb

# selachii
protoc -I proto/ --cpp_out=selachii/src/proto --grpc_out=selachii/src/proto \
 --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` proto/selachii.proto

# gorilla
protoc -I proto/ proto/gorilla.proto --go_out=plugins=grpc:gorilla/pkg/grpc/pb

# slav
mkdir -p slav/app/grpc_source
protoc -I proto/ proto/slav.proto --python_out=slav/app/grpc_source --grpc_python_out=slav/app/grpc_source \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

#picus
mkdir -p slav/app/grpc_source
protoc -I proto/ proto/picus.proto --python_out=picus/app/grpc_source --grpc_python_out=picus/app/grpc_source \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

# rhino
mkdir -p rhino/pkg/grpc/pb/
protoc -I proto/ proto/rhino.proto --go_out=plugins=grpc:rhino/pkg/grpc/pb

# ibis
mkdir -p ibis/pkg/grpc/pb/
protoc -I proto/ proto/ibis.proto --go_out=plugins=grpc:ibis/pkg/grpc/pb


# tesseract
protoc -I proto/ proto/tesseract.proto --go_out=plugins=grpc:tesseract/pkg/service/pb

###############
### clients ###
###############

# ovis -> ardea
mkdir -p ovis/app/clients/ardea
protoc -I proto/ proto/ardea.proto --python_out=ovis/app --grpc_python_out=ovis/app \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

# hippo -> ardea
mkdir -p hippo/pkg/clients/ardea/pb/
protoc -I proto/ proto/ardea.proto --go_out=plugins=grpc:hippo/pkg/clients/ardea/pb

# hippo -> selachii
mkdir -p hippo/pkg/clients/selachii/pb/
protoc -I proto/ --go_out=plugins=grpc:hippo/pkg/clients/selachii/pb proto/selachii.proto

# rhino -> ibis
mkdir -p rhino/pkg/clients/ibis/pb/
protoc -I proto/ proto/ibis.proto --go_out=plugins=grpc:rhino/pkg/clients/ibis/pb


# slav -> tesseract
mkdir -p slav/app/grpc_source
protoc -I proto/ proto/tesseract.proto --python_out=slav/app/grpc_source --grpc_python_out=slav/app/grpc_source \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

# slav -> gorilla
protoc -I proto/ proto/gorilla.proto --python_out=slav/app/grpc_source --grpc_python_out=slav/app/grpc_source \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

# lynx -> slav
mkdir -p lynx/pkg/clients/slav/pb/
protoc -I proto/ proto/slav.proto --go_out=plugins=grpc:lynx/pkg/clients/slav/pb

# lynx -> ardea
mkdir -p lynx/pkg/clients/ardea/pb/
protoc -I proto/ proto/ardea.proto --go_out=plugins=grpc:lynx/pkg/clients/ardea/pb

# lynx -> hippo
mkdir -p lynx/pkg/clients/hippo/pb/
protoc -I proto/ proto/hippo.proto --go_out=plugins=grpc:lynx/pkg/clients/hippo/pb

# lynx -> ibis
mkdir -p lynx/pkg/clients/ibis/pb/
protoc -I proto/ proto/ibis.proto --go_out=plugins=grpc:lynx/pkg/clients/ibis/pb

# lynx -> rhino
mkdir -p lynx/pkg/clients/rhino/pb/
protoc -I proto/ proto/rhino.proto --go_out=plugins=grpc:lynx/pkg/clients/rhino/pb

# lynx -> gorilla
mkdir -p lynx/pkg/clients/gorilla/pb/
protoc -I proto/ proto/gorilla.proto --go_out=plugins=grpc:lynx/pkg/clients/gorilla/pb

# lync -> picus
mkdir -p lynx/pkg/clients/picus/pb/
protoc -I proto/ proto/picus.proto --go_out=plugins=grpc:lynx/pkg/clients/picus/pb

# nalogi -> gorilla
protoc -I proto/ proto/gorilla.proto --python_out=nalogi/app --grpc_python_out=nalogi/app \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

# django -> gorilla
mkdir -p django/gorilla_pb
touch django/gorilla_pb/__init__.py 
protoc -I proto/ proto/gorilla.proto --python_out=django/gorilla_pb --grpc_python_out=django/gorilla_pb \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`

# django -> ardea
mkdir -p django/ardea_pb
touch django/ardea_pb/__init__.py
protoc -I proto/ proto/ardea.proto --python_out=django/ardea_pb --grpc_python_out=django/ardea_pb \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`


# arietes -> ibis
mkdir -p arietes/app/clients/ibis
protoc -I proto/ proto/ibis.proto --python_out=arietes/app/clients/ibis --grpc_python_out=arietes/app/clients/ibis \
 --plugin=protoc-gen-grpc_python=`which grpc_python_plugin`
