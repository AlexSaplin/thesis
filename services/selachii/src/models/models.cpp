#include "models.h"

namespace Models {

    ModelMeta::ModelMeta(const ModelMeta &other)
            : ModelID(other.ModelID), InputShape(other.InputShape), OutputShape(other.OutputShape), Type(other.Type) {};

    ModelMeta::ModelMeta(ModelMeta &&other)
            : ModelID(other.ModelID), InputShape(other.InputShape), OutputShape(other.OutputShape), Type(other.Type) {};


    ModelMeta::ModelMeta(
            const std::string model_id, const std::string path, const std::vector<std::vector<int64_t>> input_shape,
            const std::vector<std::vector<int64_t>> output_shape, const ValueType type)
            : ModelID(model_id), Path(path), InputShape(input_shape), OutputShape(output_shape), Type(type) {};

    Model::Model(Model &&other) : meta(other.meta), session(std::move(other.session)) {
        other.meta = nullptr;
    };

    Model::Model(std::unique_ptr<ModelMeta> meta_, std::vector<uint8_t> &data, DeviceType device,
                 uint8_t cuda_device_id)
            : meta(std::move(meta_)) {
        this->setupRuntimeEnv();

        if (device == CPU) {
            Ort::ThrowOnError(OrtSessionOptionsAppendExecutionProvider_CPU(this->options, 0));
        } else if (device == GPU) {
            Ort::ThrowOnError(OrtSessionOptionsAppendExecutionProvider_CUDA(this->options, cuda_device_id));
        }
        this->options.SetGraphOptimizationLevel(GraphOptimizationLevel::ORT_ENABLE_EXTENDED);

        this->session = std::make_unique<Ort::Session>(
                *(env.get()), static_cast<void *>(data.data()), data.size(), this->options);

    }

    std::vector<std::unique_ptr<Tensor>>
    Model::Run(std::vector<std::unique_ptr<Tensor>> input) {
        switch (this->meta->Type) {
            case ValueTypeFloat32:
                return std::move(this->run<float>(std::move(input)));
            case ValueTypeFloat64:
                return std::move(this->run<double>(std::move(input)));
            case ValueTypeUInt8:
                return std::move(this->run<uint8_t>(std::move(input)));
            case ValueTypeUInt16:
                return std::move(this->run<uint16_t>(std::move(input)));
            case ValueTypeUInt32:
                return std::move(this->run<uint32_t>(std::move(input)));
            case ValueTypeUint64:
                return std::move(this->run<uint64_t>(std::move(input)));
            case ValueTypeInt8:
                return std::move(this->run<int8_t>(std::move(input)));
            case ValueTypeInt16:
                return std::move(this->run<int16_t>(std::move(input)));
            case ValueTypeInt32:
                return std::move(this->run<int32_t>(std::move(input)));
            case ValueTypeInt64:
                return std::move(this->run<int64_t>(std::move(input)));
            default:
                throw std::runtime_error("unknown model value type");
        }
    }

    template<typename T>
    std::vector<std::unique_ptr<Tensor>>
    Model::run(std::vector<std::unique_ptr<Tensor>> input) {
        auto memory_info = Ort::MemoryInfo::CreateCpu(OrtArenaAllocator, OrtMemTypeDefault);

        // setup inputs
        if (this->session->GetInputCount() != input.size()) {
            throw std::runtime_error("given inputs size does not match model inputs size");
        }

        std::vector<const char *> input_names(input.size());
        // const char* input_names[input.size()];

        std::vector<Ort::Value> input_tensors;
        input_tensors.reserve(input.size());

        for (uint32_t i = 0; i < input.size(); ++i) {
            auto input_shape = input[i]->Shape();
            for (uint32_t j = 0; j < input_shape.size(); ++j) {
                if (input_shape[j] != this->meta->InputShape[i][j] && this->meta->InputShape[i][j] != 0 &&
                    input_shape[j] > 0) {
                    throw std::runtime_error("given input tensor shape does not match model input tensor shape");
                }
            }

            input_names[i] = this->session->GetInputName(i, Ort::AllocatorWithDefaultOptions());
            input_tensors.push_back(
                    Ort::Value::CreateTensor<T>(
                            memory_info,
                            reinterpret_cast<T *>(const_cast<char *>(input[i]->Ptr())),
                            input[i]->Size(),
                            input[i]->Shape().data(),
                            this->meta->InputShape[i].size())
            );
        }

        // setup outputs
        if (this->session->GetOutputCount() != this->meta->OutputShape.size()) {
            throw std::runtime_error("given outputs size does not match model outpus size");
        }

        std::vector<const char *> output_names(this->meta->OutputShape.size());
        // const char *output_names[output.size()];

        for (uint32_t i = 0; i < output_names.size(); ++i) {
            output_names[i] = this->session->GetOutputName(i, Ort::AllocatorWithDefaultOptions());
        }

        Ort::RunOptions run_options;
        // run_options.SetRunLogSeverityLevel(0);

        auto output_tensors = this->session->Run(run_options, input_names.data(), input_tensors.data(),
                                                 input_tensors.size(), output_names.data(), output_names.size());

        std::vector<std::unique_ptr<Tensor>> result(output_tensors.size());
        for (uint32_t i = 0; i < output_tensors.size(); ++i) {
            auto tensor_shape_info = output_tensors[i].GetTensorTypeAndShapeInfo();
            std::vector<int64_t> shape(tensor_shape_info.GetDimensionsCount());
            tensor_shape_info.GetDimensions(shape.data(), shape.size());

            T *tensor_data = output_tensors[i].GetTensorMutableData<T>();

            auto tensor = std::make_unique<Tensor>(this->meta->Type, shape);
            memcpy(tensor->Ptr(), tensor_data, tensor->Size());

            result[i] = std::move(tensor);
        }

        return std::move(result);
    }

    std::shared_ptr<ModelMeta> Model::GetMeta() const {
        return this->meta;
    }

    void Model::setupRuntimeEnv() {
        std::unique_lock<std::mutex> lock(envMutex);
        if (env == nullptr) {
            env = std::make_unique<Ort::Env>(ORT_LOGGING_LEVEL_WARNING, "ort");
        }
    }
}
