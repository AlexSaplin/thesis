#include "service.h"

namespace Service {
    WorkerServiceImpl::WorkerServiceImpl(std::shared_ptr<Controller::Controller> controller_, std::shared_ptr<spdlog::logger> logger_) 
    : controller(controller_), logger(logger_) {};

    grpc::Status WorkerServiceImpl::LoadModel(
        grpc::ServerContext* context, const selachii::LoadRequest* request, selachii::LoadResponse* response
    ) {
        auto model = request->model();
        auto t_start = std::chrono::high_resolution_clock::now();

        auto model_meta = parseModelMetaProto(model);
        auto model_id = model_meta->ModelID;
        logger->debug("handling LoadModel\tmodelID={}", model_id);

        try {
            auto result = this->controller->LoadModel(std::move(model_meta));
            response->set_loadid(result);

            auto t_end = std::chrono::high_resolution_clock::now();
            double elapsed_time_ms = std::chrono::duration<double, std::milli>(t_end-t_start).count();
            logger->info("LoadModel complete\tmodelID={}\tloadID={}\tlatency_ms={}", model_id, result, elapsed_time_ms);
        }
        catch(const std::exception& e) {
            logger->error("LoadModel failed\tmodelID={}\terror={}", model_id, e.what());
            return grpc::Status(grpc::StatusCode::INTERNAL, e.what());
        }

        return grpc::Status::OK;
    }


    grpc::Status WorkerServiceImpl::UnloadModel(
        grpc::ServerContext* context, const selachii::UnloadRequest* request, selachii::UnloadResponse* response
    ) {
        auto t_start = std::chrono::high_resolution_clock::now();
        auto load_id = request->loadid();

        logger->debug("handling UnloadModel\tloadID={}", load_id);

        try {
            auto result = this->controller->UnloadModel(load_id);
            response->set_didchange(result);

            auto t_end = std::chrono::high_resolution_clock::now();
            double elapsed_time_ms = std::chrono::duration<double, std::milli>(t_end-t_start).count();
            logger->info("UnloadModel complete\tloadID={}\tchanged={}\tlatency_ms={}", load_id, result, elapsed_time_ms);
        }
        catch(const std::exception& e) {
            logger->error("UnloadModel failed\tloadID={}\terror={}", load_id, e.what());
            return grpc::Status(grpc::StatusCode::INTERNAL, e.what());
        }

        return grpc::Status::OK;
    }

    grpc::Status WorkerServiceImpl::Run(
        grpc::ServerContext* context, const selachii::RunRequest* request, selachii::RunResponse* response
    ) {
        auto t_start = std::chrono::high_resolution_clock::now();
        auto load_id = request->loadid();

        logger->debug("handling Run\tloadID={}", load_id);

        try {
            auto input = parseRunRequest(request);
            auto output = this->controller->Run(load_id, std::move(input));
            makeRunResponse(response, std::move(output));

            auto t_end = std::chrono::high_resolution_clock::now();
            double elapsed_time_ms = std::chrono::duration<double, std::milli>(t_end-t_start).count();
            logger->info("Run complete\tloadID={}\tlatency_ms={}", load_id, elapsed_time_ms);
        }
        catch(const std::exception& e) {
            logger->error("Run failed\tloadID={}\terror={}", load_id, e.what());
            return grpc::Status(grpc::StatusCode::INTERNAL, e.what());
        }
        return grpc::Status::OK;
    }
};