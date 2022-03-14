#pragma once

#include <functional>
#include <sstream>
#include <numeric>
#include <vector>

namespace Models {
    enum ValueType : uint8_t {
        ValueTypeUnknown,
        ValueTypeFloat16,
        ValueTypeFloat32,
        ValueTypeFloat64,
        ValueTypeUInt8,
        ValueTypeUInt16,
        ValueTypeUInt32,
        ValueTypeUint64,
        ValueTypeInt8,
        ValueTypeInt16,
        ValueTypeInt32,
        ValueTypeInt64,
        ValueTypeComplex64,
        ValueTypeComplex128
    };

    const size_t ValueTypeSize(ValueType type);

    class Tensor {
    public:
        Tensor(ValueType type_, std::vector<int64_t> shape);
        Tensor(ValueType type_, std::vector<int64_t> shape, char* data, const size_t data_size);
        Tensor(Tensor && other);
        ~Tensor();

        char* Ptr() const;
        const size_t Size() const;
        const std::vector<int64_t> Shape() const;
        const size_t ItemSize() const;

        ValueType Type;
    private:
        std::vector<int64_t> shape;
        char* data;
        size_t size;
        bool owns_data;
    };
};