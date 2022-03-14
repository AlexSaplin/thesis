#pragma once

#include <memory>
#include <string>
#include <iostream>
#include <fstream>

#include <cpprest/http_client.h>
#include <cpprest/filestream.h>
#include <spdlog/spdlog.h>

#include "models/models.h"

namespace Loader {
    class ModelLoader {
    public:
        virtual std::unique_ptr<Models::Model> LoadModel(std::unique_ptr<Models::ModelMeta>) = 0;
    };

    class FileModelLoader : public ModelLoader {
    public:
        std::unique_ptr<Models::Model> LoadModel(std::unique_ptr<Models::ModelMeta> model) override;
    };


    class HTTPModelLoader : public ModelLoader {
    public:
        HTTPModelLoader(const std::string base_url, Models::DeviceType device, uint8_t cuda_device_id, std::shared_ptr<spdlog::logger> logger_);
        std::unique_ptr<Models::Model> LoadModel(std::unique_ptr<Models::ModelMeta> model) override;
    private:
        std::string baseUrl;
        Models::DeviceType device;
        uint8_t cudaDeviceID;
        std::shared_ptr<spdlog::logger> logger;
    };
}
