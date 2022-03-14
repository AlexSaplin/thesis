#include "serialization.h"

namespace Service {

    std::vector<std::unique_ptr<Models::Tensor>> parseRunRequest(const selachii::RunRequest *req) {
        auto size = req->tensor().size();
        std::vector<std::unique_ptr<Models::Tensor>> result(size);
        for (uint32_t i = 0; i < size; ++i) {
            result[i] = std::move(parseTensorProto(req->tensor(i)));
        }
        return std::move(result);
    }

    void makeRunResponse(selachii::RunResponse *resp, std::vector<std::unique_ptr<Models::Tensor>> result) {
        for (auto &&resultTensor : result) {
            auto tensor = resp->add_tensor();
            makeTensorProto(tensor, std::move(resultTensor));
        }
    }

    std::unique_ptr<Models::Tensor> parseTensorProto(const selachii::Tensor &tensor) {
        std::vector<int64_t> shape(tensor.shape().value().begin(), tensor.shape().value().end());
        return std::move(std::make_unique<Models::Tensor>(
                parseValueType(tensor.type()),
                shape,
                const_cast<char *>(tensor.data().data()),
                tensor.data().size()));
    }

    void makeTensorProto(selachii::Tensor *result, std::unique_ptr<Models::Tensor> tensor) {
        result->set_type(serializeValueType(tensor->Type));
        result->set_data(std::string(tensor->Ptr(), tensor->Size()));
        auto shape = tensor->Shape();
        auto shape_pb = new selachii::Shape;
        std::for_each(shape.begin(), shape.end(), [shape_pb](int64_t next) { shape_pb->add_value(next); });
        result->set_allocated_shape(shape_pb);
    }

    std::unique_ptr<Models::ModelMeta> parseModelMetaProto(const selachii::ModelMeta &meta) {
        std::vector<std::vector<int64_t>> input(meta.inputshape().size());
        for (uint32_t i = 0; i < meta.inputshape().size(); ++i) {
            input[i] = std::vector<int64_t>(meta.inputshape(i).value().begin(), meta.inputshape(i).value().end());
        }

        std::vector<std::vector<int64_t>> output(meta.outputshape().size());
        for (int i = 0; i < meta.outputshape().size(); ++i) {
            output[i] = std::vector<int64_t>(meta.outputshape(i).value().begin(), meta.outputshape(i).value().end());
        }

        return std::move(std::make_unique<Models::ModelMeta>(meta.id(), meta.path(), input, output,
                                                             parseValueType(meta.type())));
    }

    Models::ValueType parseValueType(selachii::ValueType t) {
        switch (t) {
            case selachii::ValueType::UINT8:
                return Models::ValueTypeUInt8;
            case selachii::ValueType::INT8:
                return Models::ValueTypeInt8;
            case selachii::ValueType::FLOAT16:
                return Models::ValueTypeFloat16;
            case selachii::ValueType::UINT16:
                return Models::ValueTypeUInt16;
            case selachii::ValueType::INT16:
                return Models::ValueTypeInt16;
            case selachii::ValueType::FLOAT32:
                return Models::ValueTypeFloat32;
            case selachii::ValueType::UINT32:
                return Models::ValueTypeUInt32;
            case selachii::ValueType::INT32:
                return Models::ValueTypeInt32;
            case selachii::ValueType::FLOAT64:
                return Models::ValueTypeFloat64;
            case selachii::ValueType::UINT64:
                return Models::ValueTypeUint64;
            case selachii::ValueType::INT64:
                return Models::ValueTypeInt64;
            case selachii::ValueType::COMPLEX64:
                return Models::ValueTypeComplex64;
            case selachii::ValueType::COMPLEX128:
                return Models::ValueTypeComplex128;
            default:
                return Models::ValueTypeUnknown;
        };
    }

    selachii::ValueType serializeValueType(Models::ValueType t) {
        switch (t) {
            case Models::ValueTypeUInt8:
                return selachii::ValueType::UINT8;
            case Models::ValueTypeInt8:
                return selachii::ValueType::INT8;
            case Models::ValueTypeFloat16:
                return selachii::ValueType::FLOAT16;
            case Models::ValueTypeUInt16:
                return selachii::ValueType::UINT16;
            case Models::ValueTypeInt16:
                return selachii::ValueType::INT16;
            case Models::ValueTypeFloat32:
                return selachii::ValueType::FLOAT32;
            case Models::ValueTypeUInt32:
                return selachii::ValueType::UINT32;
            case Models::ValueTypeInt32:
                return selachii::ValueType::INT32;
            case Models::ValueTypeFloat64:
                return selachii::ValueType::FLOAT64;
            case Models::ValueTypeUint64:
                return selachii::ValueType::UINT64;
            case Models::ValueTypeInt64:
                return selachii::ValueType::INT64;
            case Models::ValueTypeComplex64:
                return selachii::ValueType::COMPLEX64;
            case Models::ValueTypeComplex128:
                return selachii::ValueType::COMPLEX128;
            default:
                return selachii::ValueType::UNKNOWN;
        };
    }
}
