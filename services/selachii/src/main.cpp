
#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>
#include <spdlog/spdlog.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <cxxopts.hpp>

#include "service/service.h"
#include "controller/controller.h"
#include "model_store/store.h"
#include "model_loader/loader.h"

void Run(std::string model_storage_url, bool use_gpu, uint8_t gpu_device_id) {
    spdlog::set_level(spdlog::level::debug);

    auto device = use_gpu ? Models::GPU : Models::CPU;

    auto loader = std::make_shared<Loader::HTTPModelLoader>(model_storage_url, device, gpu_device_id, spdlog::stdout_color_mt("loader"));
    auto store = std::make_shared<ModelStore::ModelStoreImpl>(loader);
    auto controller = std::make_shared<Controller::ControllerImpl>(store);

    Service::WorkerServiceImpl service(controller, spdlog::stdout_color_mt("service"));

    std::string server_address("0.0.0.0:8083");
    grpc::ServerBuilder builder;
    builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
    builder.RegisterService(&service);

    const int max_message_size = 1024 * 1024 * 64;

    builder.SetMaxReceiveMessageSize(max_message_size);
    builder.SetMaxSendMessageSize(max_message_size);

    std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
    spdlog::info("Server with listening on {}", server_address);
    server->Wait();
}

int main(int argc, char* argv[]) {
    cxxopts::Options options("selachii", "model executor");
    options.add_options()
        ("s,storage", "Model storage URL", cxxopts::value<std::string>())
        ("g,gpu", "GPU device ID to use", cxxopts::value<int>()->default_value("-1"));
    auto parsed = options.parse(argc, argv);
    std::string models_storage_url = parsed["storage"].as<std::string>();
    int gpu = parsed["gpu"].as<int>();
    Run(models_storage_url, gpu >= 0, gpu);
    return 0;
}
