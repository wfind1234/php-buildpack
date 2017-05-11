# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import os
#
# def _should_compile(ctx):
#     if os.getenv('PHP_FPM_LISTEN'):
#         ctx['PHP_FPM_LISTEN'] = os.getenv('PHP_FPM_LISTEN')
#         return False
#     return


def preprocess_commands(ctx):
    return ((
        '$HOME/.bp/bin/rewrite',
        '"$HOME/nginx/conf"'),)


def service_commands(ctx):
    return {
        'nginx': (
            '$HOME/nginx/sbin/nginx',
            '-c "$HOME/nginx/conf/nginx.conf"')
    }


def service_environment(ctx):
    return {}


def compile(install):
    if os.getenv('PHP_FPM_LISTEN'):
        install.builder._ctx['PHP_FPM_LISTEN'] = os.getenv('PHP_FPM_LISTEN')
        return 0

    print 'Installing Nginx'
    install.builder._ctx['PHP_FPM_LISTEN'] = '{TMPDIR}/php-fpm.socket'

    ctx = install.builder._ctx
    env_dir = '%s/%s/env' % (ctx['DEPS_DIR'], ctx['DEPS_IDX'])
    if not os.path.exists(env_dir):
        os.makedirs(env_dir)

    target = open('%s/PHP_FPM_LISTEN' % env_dir, 'w')
    target.write("%s/php-fpm.socket" % ctx['TMPDIR'])
    target.close()

    (install
        .package('NGINX')
        .config()
            .from_application('.bp-config/nginx')  # noqa
            .or_from_build_pack('defaults/config/nginx')
            .to('nginx/conf')
            .rewrite()
            .done())
    return 0
