#include "tensors.h"

namespace Models {

    const size_t ValueTypeSize(ValueType t) {
        const size_t typeSize8 = 1;
        const size_t typeSize16 = 2;
        const size_t typeSize32 = 4;
        const size_t typeSize64 = 8;
        const size_t typeSize128 = 16;
        switch (t) {
            case ValueTypeUInt8:
                return typeSize8;
            case ValueTypeInt8:
                return typeSize8;
            case ValueTypeFloat16:
                return typeSize16;
            case ValueTypeUInt16:
                return typeSize16;
            case ValueTypeInt16:
                return typeSize16;
            case ValueTypeFloat32:
                return typeSize32;
            case ValueTypeUInt32:
                return typeSize32;
            case ValueTypeInt32:
                return typeSize32;
            case ValueTypeFloat64:
                return typeSize64;
            case ValueTypeUint64:
                return typeSize64;
            case ValueTypeInt64:
                return typeSize64;
            case ValueTypeComplex64:
                return typeSize64;
            case ValueTypeComplex128:
                return typeSize128;
            default:
                throw std::runtime_error("unknown ValueType");
        };
    }

    Tensor::Tensor(Tensor &&other) : shape(other.shape), data(other.data), Type(other.Type), size(other.size) {
        other.data = nullptr;
        other.shape = {};
    }

    Tensor::Tensor(ValueType type_, std::vector<int64_t> shape) {
        this->Type = type_;
        this->shape = shape;

        auto shape_product = std::accumulate(this->shape.begin(), this->shape.end(), 1, std::multiplies<size_t>());

        this->size = ValueTypeSize(type_) * shape_product;
        this->data = new char[this->size];
        this->owns_data = true;
    }

    Tensor::Tensor(ValueType type_, std::vector<int64_t> shape, char* data, const size_t data_size) {
        this->shape = shape;
        this->size = data_size;
        
        auto product = std::accumulate(this->shape.begin(), this->shape.end(), 1, std::multiplies<size_t>());
        if (!product * ValueTypeSize(type_) == this->size) {
            std::ostringstream errfmt;
            errfmt << "failed to make tensor: invalid data size " <<  product * ValueTypeSize(type_) << " and got " << data_size;
            throw std::runtime_error(errfmt.str());
        }
        this->data = data;
        this->Type = type_;
        this->owns_data = false;
    }

    Tensor::~Tensor() {
        if (this->owns_data) {
            delete this->data;
        }
    }

    char* Tensor::Ptr() const {
        return this->data;
    }

    const size_t Tensor::Size() const {
        return this->size;
    }

    const std::vector<int64_t> Tensor::Shape() const {
        return this->shape;
    }

    const size_t Tensor::ItemSize() const {
        return ValueTypeSize(this->Type);
    }
};
