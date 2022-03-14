#pragma once

#include <memory>
#include <string>
#include <unordered_map>

#include "models/models.h"
#include "models/tensors.h"
#include "model_loader/loader.h"
#include "util/uuid.h"

namespace ModelStore {
    class ModelStore {
    public:
        virtual std::string PreloadModel(std::unique_ptr<Models::ModelMeta>) = 0;
        virtual bool UnloadModel(std::string) = 0;
        virtual std::shared_ptr<Models::Model> GetModel(std::string) = 0;
    };

    class ModelStoreImpl : public ModelStore {
    public:
        ModelStoreImpl(std::shared_ptr<Loader::ModelLoader>);

        std::string PreloadModel(std::unique_ptr<Models::ModelMeta>) override;
        bool UnloadModel(std::string) override;
        std::shared_ptr<Models::Model> GetModel(std::string) override;
    private:
        std::unordered_map<std::string, std::shared_ptr<Models::Model>> store;
        std::shared_ptr<Loader::ModelLoader> loader;
    };
};