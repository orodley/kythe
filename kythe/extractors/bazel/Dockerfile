# Copyright 2019 The Kythe Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM gcr.io/cloud-builders/bazel

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
      # Install C++ compilers for cc_* rule support.
      gcc clang && \
    apt-get clean

# Extract the Kythe release archive to /kythe
COPY kythe/release/kythe-v*.tar.gz /tmp/
RUN tar xzf /tmp/kythe-v*.tar.gz && \
    mv kythe-v*/ /kythe && \
    rm /tmp/kythe-v*.tar.gz

# Tools and configuration
ADD kythe/extractors/bazel/extract.sh /kythe/
ADD kythe/release/base/fix_permissions.sh /kythe/

# Bazel repository setup
ADD kythe/extractors/bazel/extractors.bazelrc /kythe/bazelrc
RUN cat /kythe/bazelrc > ~/.bazelrc

ENTRYPOINT ["/kythe/extract.sh"]
