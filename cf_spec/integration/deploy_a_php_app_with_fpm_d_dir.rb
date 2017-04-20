$: << 'cf_spec'
require 'cf_spec_helper'

describe 'CF PHP Buildpack' do
  let(:browser)  { Machete::Browser.new(@app) }

  before(:context) { @app_name = 'with_fpm_d'}

  context 'deploying a basic PHP app with custom conf files in fpm.d dir in app root' do
    before(:all) do
      @env_config = {env: {'COMPOSER_GITHUB_OAUTH_TOKEN' => ENV['COMPOSER_GITHUB_OAUTH_TOKEN']}}
      @app = deploy_app(@app_name, @env_config)
    end

    after(:all) do
      Machete::CF::DeleteApp.new.execute(@app)
    end

    it 'expects an app to be running' do
      expect(@app).to be_running
    end

    it 'sets custom configurations' do
      browser.visit_path('/index.php')
      expect(browser).to have_body 'TEST_HOME_PATH'
      expect(browser).to have_body '/home/vcap/app/test/path'
      expect(browser).to have_body 'TEST_WEBDIR'
      expect(browser).to have_body 'htdocs'
    end
  end
end

