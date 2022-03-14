import importlib
import logging
import os
import time
import traceback
from threading import Thread

from flask import Flask, request, cli


cli.show_server_banner = lambda *_: None

module_name = os.getenv('PACKAGE_NAME')
function_to_call = os.getenv('FUNCTION_NAME')

module = None
fail = None


def import_module():
    global module, fail
    try:
        module = importlib.import_module(module_name)
    except Exception as e:
        current_logger.error(f'Loading module failed. Error: {repr(e)}')
        fail = traceback.format_exc()


app = Flask(__name__)
flask_log = logging.getLogger('werkzeug')
flask_log.setLevel(logging.CRITICAL)
flask_log.disabled = True
app.logger.disabled = True


current_logger = logging.getLogger('DeepMux-Runner')


@app.route('/run', methods=['POST'])
def run():
    if module is None:
        if fail is not None:
            return fail, 500
        return "Module is not loaded", 500
    try:
        result = getattr(module, function_to_call)(request.get_data())
        return result, 200
    except Exception as e:
        current_logger.error(f'Running failed. Error: {repr(e)}')
        return traceback.format_exc(), 400


@app.route('/load')
def load():
    while module is None and fail is None:
        time.sleep(0.01)

    if fail is not None:
        return fail, 400
    elif module is not None:
        return "", 200
    else:
        return "Module is not loaded", 500


if __name__ == '__main__':
    import_thread = Thread(target=import_module)
    import_thread.start()

    app.run('0.0.0.0', 80)
