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
from compile_helpers import write_env_file

def preprocess_commands(ctx):
    return ((
        '$HOME/.bp/bin/rewrite',
        '"$HOME/httpd/conf"'),)


def service_commands(ctx):
    return {
        'httpd': (
            '$HOME/httpd/bin/apachectl',
            '-f "$HOME/httpd/conf/httpd.conf"',
            '-k start',
            '-DFOREGROUND')
    }


def service_environment(ctx):
    return {
        'HTTPD_SERVER_ADMIN': ctx['ADMIN_EMAIL']
    }


def compile(install):
    ctx = install.builder._ctx
    if os.getenv('PHP_FPM_LISTEN'):
        ctx['PHP_FPM_LISTEN'] = os.getenv('PHP_FPM_LISTEN')
        return 0

    print 'Installing Nginx'
    ctx['PHP_FPM_LISTEN'] = '{TMPDIR}/php-fpm.socket'

    write_env_file(ctx, 'PHP_FPM_LISTEN', '%s/php-fpm.socket' % ctx['TMPDIR'])

    print 'Installing HTTPD'
    print 'HTTPD %s' % (ctx['HTTPD_VERSION'])

    ctx['PHP_FPM_LISTEN'] = '127.0.0.1:9000'
    (install
        .package('HTTPD')
        .config()
            .from_application('.bp-config/httpd')  # noqa
            .or_from_build_pack('defaults/config/httpd')
            .to('httpd/conf')
            .rewrite()
            .done())
    return 0