$: << 'cf_spec'
require 'cf_spec_helper'
require 'excon'

describe 'CF PHP Buildpack' do
  let(:browser)  { Machete::Browser.new(@app) }

  context 'deploying a basic PHP app using argon2i crypto hash' do
    before(:all) do
      @app_name = 'php_72_argon'
      @env_config = {env: {'COMPOSER_GITHUB_OAUTH_TOKEN' => ENV['COMPOSER_GITHUB_OAUTH_TOKEN']}}
      @app = deploy_app(@app_name, @env_config)
    end
    after(:all) do
      Machete::CF::DeleteApp.new.execute(@app)
    end

    it 'performs an argon2 hash' do
      expect(@app).to be_running
      expect(@app).to have_logged "-------> Buildpack version #{File.read(File.expand_path('../../../VERSION', __FILE__)).chomp}"

      browser.visit_path("/")
      expect(browser).to have_body '4j2ZFDn1fVS70ZExmlJ33rXOinafcBXrp6A6grHEPkI'
    end
  end
end

