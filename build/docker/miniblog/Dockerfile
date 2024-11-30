# Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/qylearn/miniblog.


FROM BASE_IMAGE
LABEL maintainer="<qinyun_77@163.com>"

WORKDIR /opt/miniblog

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/miniblog

COPY miniblog /opt/miniblog/bin/

ENTRYPOINT ["/opt/miniblog/bin/miniblog"]
CMD ["-c", "/etc/miniblog/miniblog.yaml"]
