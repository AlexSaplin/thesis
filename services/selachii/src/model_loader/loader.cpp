#include "loader.h"

namespace Loader {
    std::unique_ptr<Models::Model> FileModelLoader::LoadModel(std::unique_ptr<Models::ModelMeta> meta) {
        throw std::runtime_error("file model loader disabled");
        // std::ifstream freader = std::ifstream(meta->Path, std::ios::out | std::ios::binary);
        // std::string body(std::istreambuf_iterator<char>(freader), {});
        // return std::move(std::make_unique<Models::Model>(std::move(meta), std::make_pair<char*, size_t>(body.data(), body.size()), Models::CPU, 0));
    }

    HTTPModelLoader::HTTPModelLoader(
            const std::string base_url,
            Models::DeviceType device_,
            uint8_t cuda_device_id,
            std::shared_ptr<spdlog::logger> logger_
    ) : baseUrl(base_url), device(device_), cudaDeviceID(cuda_device_id), logger(logger_) {}

    std::unique_ptr<Models::Model> HTTPModelLoader::LoadModel(std::unique_ptr<Models::ModelMeta> meta) {
        auto model_id = meta->ModelID;

        auto t_start = std::chrono::high_resolution_clock::now();

        web::http::client::http_client client(this->baseUrl);
        web::uri_builder builder(meta->Path);

        logger->debug("loading model\tmodelID={}\turl={}", model_id, builder.to_string());

        std::vector<uint8_t> data_vector;
        client.request(web::http::methods::GET, builder.to_string())
        .then([&data_vector](web::http::http_response response) {
            if (response.status_code() != web::http::status_codes::OK) {
                throw std::runtime_error(response.reason_phrase());
            }
            // avoiding extra copy
            data_vector = std::move(response.extract_vector().get());
        }).wait();

        auto t_request = std::chrono::high_resolution_clock::now();
        double elapsed_time_ms = std::chrono::duration<double, std::milli>(t_request-t_start).count();
        logger->info("finished downloading model\tmodelID={}\tlatency_ms={}", model_id, elapsed_time_ms);

        auto model = std::make_unique<Models::Model>(std::move(meta), data_vector, this->device, this->cudaDeviceID);

        auto t_create = std::chrono::high_resolution_clock::now();
        elapsed_time_ms = std::chrono::duration<double, std::milli>(t_create - t_request).count();
        logger->info("finished initializing model\tmodelID={}\tlatency_ms={}", model_id, elapsed_time_ms);

        return std::move(model);
    }
};
