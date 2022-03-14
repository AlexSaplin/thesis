#pragma once

#include <memory>

#include <grpcpp/grpcpp.h>
#include <spdlog/spdlog.h>

#include "controller/controller.h"
#include "models/models.h"
#include "proto/selachii.grpc.pb.h"

namespace Service{

    std::vector<std::unique_ptr<Models::Tensor>> parseRunRequest(const selachii::RunRequest*);
    void makeRunResponse(selachii::RunResponse*, std::vector<std::unique_ptr<Models::Tensor>>);

    std::unique_ptr<Models::Tensor> parseTensorProto(const selachii::Tensor&);
    void makeTensorProto(selachii::Tensor*, std::unique_ptr<Models::Tensor>);

    std::unique_ptr<Models::ModelMeta> parseModelMetaProto(const selachii::ModelMeta&);

    selachii::ValueType serializeValueType(Models::ValueType);
    Models::ValueType parseValueType(selachii::ValueType);
}