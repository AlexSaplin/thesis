FROM {{ registry_url }}/{{ base_image }}

ENV PACKAGE_NAME={{ call_module }}
ENV FUNCTION_NAME={{ call_object }}

RUN apt update && apt install -y libsm6 libxext6 libxrender-dev libgl1-mesa-glx ffmpeg

# Requirements for server
RUN pip install flask

# Ok let's set a workdir
# How we want everything to look:
# /tmp/random-folder/
# -> Dockerfile
# -> app
#   -> entrypoint.sh (with another random name)
#   -> server.py (random name)
#   -> whole files
# then call just docker build
# docker push

WORKDIR /app

COPY app /app

{% if requirements_path %}
RUN pip install -r {{ requirements_path }}
{% endif %}

RUN chmod +x {{ entrypoint_name }}

ENTRYPOINT ["bash", "{{ entrypoint_name }}"]
