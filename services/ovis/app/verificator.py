import logging
import traceback
from typing import List

import numpy
import requests
import onnxruntime

from models import Message
from ardea_pb2_grpc import ArdeaStub
from ardea_pb2 import GetModelRequest, UpdateModelStateRequest, Model, Shape, STATE_READY, STATE_INVALID

logger = logging.getLogger(__name__)


class ModelValidationError(Exception):
    ...


class ModelVerificator:

    def __init__(self, ardea: ArdeaStub, model_storage_url: str):
        self.ardea = ardea
        self.model_storage_url = model_storage_url

    def verify(self, msg: Message):
        try:
            logger.info(f"verifying model with id: {msg.model_id}")

            model = self.ardea.GetModel(GetModelRequest(ID=str(msg.model_id))).Model

            self._verify_model(model)

            self.ardea.UpdateModelState(UpdateModelStateRequest(ID=model.ID, State=STATE_READY))
            logger.info("model verification successful")
        except ModelValidationError as e:
            self.ardea.UpdateModelState(
                UpdateModelStateRequest(ID=model.ID, State=STATE_INVALID, ErrStrSet=True, ErrStr=str(e)))
            logger.info(f"model verification failed: {e}")
        except:
            logger.error(f"unhandled error verifying model: {traceback.format_exc()}")

    def _verify_model(self, model: Model):
        # get model data
        url = f'{self.model_storage_url}/{model.Path}'
        response = requests.get(url, timeout=60)
        if response.status_code != 200:
            raise Exception(f"couldn't fetch model from {url}, status code {response.status_code}")
        model_data = response.content
        logger.info("loaded model data")
        # setup session
        try:
            session = onnxruntime.InferenceSession(model_data, None)
        except Exception as e:
            raise ModelValidationError(f"failed to load model for validation: {e}")
        logger.info("session initialized")
        # check inputs
        inputs: List[onnxruntime.NodeArg] = session.get_inputs()
        transformed_inputs = [[0 if isinstance(x, str) else x for x in input.shape] for input in inputs]
        if len(inputs) != len(model.InputShape):
            raise ModelValidationError(
                f"model is expected to have {len(model.InputShape)} inputs, but has {len(inputs)}: {[x.name for x in inputs]}")

        for i in range(len(inputs)):
            meta_shape = self._shape_to_list(model.InputShape[i])

            if not self._compare_shapes(transformed_inputs[i], meta_shape):
                raise ModelValidationError(
                    f"model input shape at idx {i} is expected to be {meta_shape}, but is {inputs[i].shape}")

        logger.info("inputs validated")

        # check outputs
        outputs: List[onnxruntime.NodeArg] = session.get_outputs()
        transformed_outputs = [[0 if isinstance(x, str) else x for x in output.shape] for output in outputs]
        if len(outputs) != len(model.OutputShape):
            logger.warning(f"model can have only 1 output, but has {len(outputs)}: {[x.name for x in outputs]}")

        for i in range(len(outputs)):
            meta_shape = self._shape_to_list(model.OutputShape[i])

            if not self._compare_shapes(transformed_outputs[i], meta_shape):
                raise ModelValidationError(
                    f"model output shape at idx {i} is expected to be {meta_shape}, but is {outputs[i].shape}")

        logger.info("outputs validated")

        # run model
        try:
            sess_inputs = {
                cur_input.name: numpy.zeros(shape=[1 if isinstance(x, str) else x for x in cur_input.shape],
                dtype=numpy.float32)
                for cur_input in inputs
            }

            session.run([], sess_inputs)[0]
        except Exception as e:
            raise ModelValidationError(f"failed to test run model: {e}")
        logger.info("test run successful")

    @staticmethod
    def _shape_to_list(shape: Shape):
        return [item.Value if item.IsValid else 0 for item in shape.Value]

    @staticmethod
    def _compare_shapes(shape: list, meta_shape: list) -> bool:
        if len(shape) != len(meta_shape):
            return False

        for i in range(len(shape)):
            if shape[i] != meta_shape[i] and shape[i] is not None and meta_shape[i] is not None:
                return False

        return True
