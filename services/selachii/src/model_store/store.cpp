#include "store.h"

namespace ModelStore {
    ModelStoreImpl::ModelStoreImpl(std::shared_ptr<Loader::ModelLoader> loader_) : loader(loader_) {};

    std::string ModelStoreImpl::PreloadModel(std::unique_ptr<Models::ModelMeta> meta) {
        auto load_id = Util::GenerateUUIDV4();

        std::shared_ptr<Models::Model> model = std::move(this->loader->LoadModel(std::move(meta)));
        
        this->store[load_id] = model;
        return load_id;
    }

    bool ModelStoreImpl::UnloadModel(std::string load_id) {
        if (this->store.count(load_id) == 0) {
            return false;
        }
        this->store.erase(load_id);
        return true;
    }

    std::shared_ptr<Models::Model> ModelStoreImpl::GetModel(std::string load_id) {
        return this->store[load_id];
    }
};