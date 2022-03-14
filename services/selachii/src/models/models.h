#pragma once

#include <memory>
#include <vector>
#include <mutex>
#include <numeric>
#include <iostream>
#include <iterator>

#ifdef __APPLE__
#include <onnxruntime/core/session/onnxruntime_cxx_api.h>
#else
#include "onnxruntime_cxx_api.h"
#include "onnxruntime_c_api.h"
#include "cpu_provider_factory.h"
#include "cuda_provider_factory.h"
#endif

#include "tensors.h"

namespace Models {
    enum DeviceType : uint8_t {
        CPU,
        GPU
    };



    struct ModelMeta {
        const std::string ModelID;
        const std::vector<std::vector<int64_t>> InputShape;
        const std::vector<std::vector<int64_t>> OutputShape;
        const ValueType Type;
        const std::string Path;

        ModelMeta(const ModelMeta&);
        ModelMeta(ModelMeta&&);
        ModelMeta(std::string, std::string, std::vector<std::vector<int64_t>>, std::vector<std::vector<int64_t>>,
                ValueType);
    };
    
    class Model {
    public:
        Model() = delete;
        Model(Model&&);
        Model(std::unique_ptr<ModelMeta>, std::vector<uint8_t>&, DeviceType, uint8_t);

        std::vector<std::unique_ptr<Tensor>>
        Run(std::vector<std::unique_ptr<Tensor>> input);
        std::shared_ptr<ModelMeta>  GetMeta() const;
    
    private:
        template<typename T>
        std::vector<std::unique_ptr<Tensor>>
        run(std::vector<std::unique_ptr<Tensor>> input);
        void setupRuntimeEnv();

        std::shared_ptr<ModelMeta> meta;
        std::unique_ptr<Ort::Session> session;
	Ort::SessionOptions options;

        inline static std::unique_ptr<Ort::Env> env;
        inline static std::mutex envMutex;
    };
};
