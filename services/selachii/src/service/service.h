#pragma once

#include <memory>
#include <chrono>

#include <grpcpp/grpcpp.h>
#include <spdlog/spdlog.h>

#include "serialization.h"
#include "controller/controller.h"
#include "proto/selachii.grpc.pb.h"

namespace Service {
    class WorkerServiceImpl final : public selachii::Selachii::Service {
    public: 
        WorkerServiceImpl(std::shared_ptr<Controller::Controller>, std::shared_ptr<spdlog::logger> logger_);

        grpc::Status LoadModel(
                grpc::ServerContext* context,
                const selachii::LoadRequest* request,
                selachii::LoadResponse* response
                ) override;

        grpc::Status UnloadModel(
                grpc::ServerContext* context,
                const selachii::UnloadRequest* request,
                selachii::UnloadResponse* response
                ) override;

        grpc::Status Run(
                grpc::ServerContext* context,
                const selachii::RunRequest* request,
                selachii::RunResponse* response
                ) override;

    private:
        std::shared_ptr<Controller::Controller> controller;
        std::shared_ptr<spdlog::logger> logger;
    };
};