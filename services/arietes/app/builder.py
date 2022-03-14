import stat
import subprocess
from io import BytesIO
import os
from shutil import copy2
import tempfile
from zipfile import ZipFile

import grpc
import requests
import yaml
from jinja2 import Template

from clients.ibis.ibis_pb2 import GetFunctionRequest, UpdateFunctionParam, FunctionState, UpdateFunctionRequest
from clients.ibis.ibis_pb2_grpc import IbisStub
from config import REGISTRY_URL, IBIS_TARGET, S3_TARGET, AVAILABLE_IMAGES
from errors import ConfigValidationError, ImageError
from utils.helpers import generate_random_str


class FunctionBuilder:

    def __init__(self, logger):
        grpc_ibis_channel = grpc.insecure_channel(IBIS_TARGET)
        self.ibis_stub = IbisStub(grpc_ibis_channel)
        self.logger = logger

    def _fetch_and_validate_yaml(self, path):
        self.logger.info('Fetching and validating yaml...')
        with open(f'{path}/deepmux.yaml') as fetch_yaml:
            params_dict = yaml.load(fetch_yaml.read(), Loader=yaml.FullLoader)
            image_name = params_dict.get('env')
            if image_name is None:
                raise Exception('There is no environment image')
            if image_name not in AVAILABLE_IMAGES:
                raise Exception('There is no such image')
            # Python 3 support
            # TODO: Move this logic to class
            python_dict = params_dict.get('python')
            if python_dict is None:
                raise Exception('DeepMux currently supports only Python')
            requirements_path = python_dict.get("requirements")
            if 'call' not in python_dict:
                raise Exception('There is no function to call. You can provide it using call argument.')
            try:
                call_module, call_object = python_dict.get('call').split(':')
            except Exception as e:
                raise Exception('Bad format of call object')
            return image_name, requirements_path, call_module, call_object

    def _make_files(self, path, image_name, requirements_path, call_module, call_object):
        # Copying server file
        server_name = f'server-{generate_random_str()}.py'
        copy2('./utils/python_image/server.py', f'{path}/app/{server_name}')

        # Copying entrypoint.sh
        entrypoint_name = f'entrypoint-{generate_random_str()}.sh'
        with open('./utils/python_image/entrypoint.sh', 'r') as entrypoint_template:
            template = Template(entrypoint_template.read())
            entrypoint_source = template.render(
                server_name=server_name,
            )

        entrypoint_path = f'{path}/app/{entrypoint_name}'

        with open(entrypoint_path, 'a+') as new_entrypoint:
            new_entrypoint.write(entrypoint_source)

        # chmod +x entrypoint.sh
        st = os.stat(entrypoint_path)
        os.chmod(entrypoint_path, st.st_mode | stat.S_IEXEC)

        # Constructing Dockerfile
        with open('./utils/python_image/Dockerfile', 'r') as df:
            template = Template(df.read())
            dockerfile_source = template.render(
                base_image=image_name,
                dir_path=path,
                requirements_path=requirements_path,
                entrypoint_name=entrypoint_name,
                call_module=call_module,
                call_object=call_object,
                registry_url=REGISTRY_URL,
            )

        with open(f'{path}/Dockerfile', 'a+') as new_df:
            new_df.write(dockerfile_source)

        # Create __init__.py if not exists
        open(f'{path}/app/__init__.py', 'a+').close()

    def _build_container(self, zip_url):
        with tempfile.TemporaryDirectory() as tmpdirname:
            # Download & extract zip
            self.logger.info(f'Starting extracting zip')
            resp = requests.get(f'{S3_TARGET}/{zip_url}')
            zip_file_object = BytesIO(resp.content)
            zip_file = ZipFile(zip_file_object)
            app_dir_path = f'{tmpdirname}/app'
            os.mkdir(app_dir_path)
            zip_file.extractall(app_dir_path)

            # Load and validate yaml
            try:
                image_name, requirements_path, call_module, call_object = self._fetch_and_validate_yaml(app_dir_path)
            except Exception as e:
                raise ConfigValidationError(f'Bad config file. Error: {e}')

            self.logger.info(f'Starting extracting files')
            # Make files (entrypoint.sh, server, Dockerfile & verify __init__.py)
            self._make_files(tmpdirname, image_name, requirements_path, call_module, call_object)

            # Build & push
            self.logger.info('Starting building')
            generated_image_name = f'{REGISTRY_URL}/{generate_random_str()}'

            x = subprocess.run(['ls', '-a', tmpdirname], capture_output=True)
            self.logger.info(x.stdout.decode())

            status = subprocess.Popen(['docker', 'build', '-t', generated_image_name, tmpdirname]).wait()

            if status != 0:
                raise ImageError('Failed to build image')

            status = subprocess.Popen(['docker', 'push', generated_image_name]).wait()

            if status != 0:
                raise ImageError('Failed to push image')
        self.logger.debug(f'ending building')
        return generated_image_name

    def process_request(self, function_id):
        self.logger.info(f'Fetching function {function_id}')
        function_data = self.ibis_stub.GetFunction(GetFunctionRequest(ID=function_id)).Function
        self.logger.info(f'Updating function {function_id}')
        self.ibis_stub.UpdateFunction(UpdateFunctionRequest(
            ID=function_id,
            Params=[
                UpdateFunctionParam(
                    State=FunctionState.STATE_PROCESSING
                )
            ]
        ))
        try:
            image_name = self._build_container(function_data.CodePath)
            self.ibis_stub.UpdateFunction(UpdateFunctionRequest(
                ID=function_id,
                Params=[
                    UpdateFunctionParam(
                        State=FunctionState.STATE_READY
                    ),
                    UpdateFunctionParam(
                        ImageURL=image_name
                    )
                ]
            ))
        except Exception as e:
            self.ibis_stub.UpdateFunction(UpdateFunctionRequest(
                ID=function_id,
                Params=[
                    UpdateFunctionParam(
                        State=FunctionState.STATE_INVALID,
                    ),
                    UpdateFunctionParam(
                        ErrStr=repr(e),
                    )
                ]
            ))
