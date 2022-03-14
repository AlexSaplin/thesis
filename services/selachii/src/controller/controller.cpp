#include "controller.h"


namespace Controller {
    ControllerImpl::ControllerImpl(std::shared_ptr<ModelStore::ModelStore> store_) : store(store_) {};

    std::string ControllerImpl::LoadModel(std::unique_ptr<Models::ModelMeta> meta) {
        return this->store->PreloadModel(std::move(meta));
    }
    
    bool ControllerImpl::UnloadModel(std::string load_id) {
        return this->store->UnloadModel(load_id);
    }

    std::vector<std::unique_ptr<Models::Tensor>>
    ControllerImpl::Run(const std::string& load_id, std::vector<std::unique_ptr<Models::Tensor>> input) {
        auto model = this->store->GetModel(load_id);
        auto output = ControllerImpl::buildOutputTensors(model->GetMeta());
        return std::move(model->Run(std::move(input)));
    }


    std::vector<std::unique_ptr<Models::Tensor>>
    ControllerImpl::buildOutputTensors(std::shared_ptr<Models::ModelMeta> meta) {
        auto size = meta->OutputShape.size();
        std::vector<std::unique_ptr<Models::Tensor>> result(size);
        for (uint32_t i = 0; i < size; ++i) {
            result[i] = std::move(std::make_unique<Models::Tensor>(meta->Type, meta->OutputShape[i]));
        }
        return std::move(result);
    }
};