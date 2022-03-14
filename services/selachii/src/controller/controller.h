#pragma once

#include <memory>
#include <string>

#include "models/tensors.h"
#include "models/models.h"
#include "model_store/store.h"


namespace Controller {

    class Controller {
    public:
        virtual std::string LoadModel(std::unique_ptr<Models::ModelMeta>) = 0;
        virtual bool UnloadModel(std::string load_id) = 0;

        virtual std::vector<std::unique_ptr<Models::Tensor>>
        Run(const std::string& load_id, std::vector<std::unique_ptr<Models::Tensor>> input) = 0;
    };

    class ControllerImpl final : public Controller {
    public:
        ControllerImpl(std::shared_ptr<ModelStore::ModelStore>);

        std::string LoadModel(std::unique_ptr<Models::ModelMeta>) override;
        bool UnloadModel(std::string model_id) override;

        std::vector<std::unique_ptr<Models::Tensor>>
        Run(const std::string& load_id, std::vector<std::unique_ptr<Models::Tensor>> input) override;
    
    private:
        static std::vector<std::unique_ptr<Models::Tensor>> buildOutputTensors(std::shared_ptr<Models::ModelMeta> meta);
        std::shared_ptr<ModelStore::ModelStore> store;
     };
};