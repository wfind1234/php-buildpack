package integration_test

import (
	"os"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CF PHP Buildpack", func() {
	var app *cutlass.App
	AfterEach(func() { app = DestroyApp(app) })

	var manifest struct {
		Dependencies []struct{
			Name string `json:"name"`
			Version string `json:"version"`
			Modules []string `json:"modules"`
		} `json:"dependencies"`
	}
	Before(func() {
		Expect((&libbuildpack.YAML{}).Load(filepath.Join(bpDir, "manifest.yml"), &manifest)).To(Succeed())
	})

	// php_buildpack_manifest_file = File.join(File.expand_path(File.dirname(__FILE__)), '..', '..', 'manifest.yml')
	// manifest = YAML.load(File.read(php_buildpack_manifest_file))

	// php_modules = {}
	// php_modules['PHP 5'] = manifest['dependencies'].select { |dependency| dependency['name'] == "php" && dependency['version'].to_s[0] == '5' }.first
	// php_modules['PHP 7.0'] = manifest['dependencies'].select { |dependency| dependency['name'] == "php" && dependency['version'].to_s[0,3] == '7.0' }.first
	// php_modules['PHP 7.1'] = manifest['dependencies'].select { |dependency| dependency['name'] == "php" && dependency['version'].to_s[0,3] == '7.1' }.first

	// def setup_app(app_name)
	//     env_config = {env:  {'COMPOSER_GITHUB_OAUTH_TOKEN' => ENV['COMPOSER_GITHUB_OAUTH_TOKEN']}}
	//     deploy_app(app_name, env_config)
	// end

	//   shared_examples_for 'it loads all the modules' do |php_version|
	func ItLoadsAllTheModules(phpVersion string) {
		By("logs each module on the info page", func() {
			Expect(app.Stdout.String()).To(ContainSubstring("PHP " + phpVersion))
			body, err := app.GetBody("/")
			Expect(err).ToNot(HaveOccurred())

			var modules []string
			for _, d := range manifest.Dependencies {
				if d.Name == "php" && strings.HasPrefix(d.Version, phpVersion) {
					modules = d.Modules
					break
				}
			}

			//       php_modules[php_version]['modules'].each do |module_name|
			//         if module_name == 'ioncube'
			//           expect(browser).to have_body /ionCube&nbsp;PHP&nbsp;Loader&nbsp;\(enabled\)/
			//         else
			//           expect(browser).to have_body /module_(Zend[+ ])?#{module_name}/i
			//         end
			//       end
			//     end
		})
	}

	FContext("extensions are specified in .bp-config", func() {
		Context("deploying a basic PHP5 app that loads all prepackaged extensions", func() {
			BeforeEach(func() {
				app = cutlass.New(filepath.Join(bpDir, "cf_spec", "fixtures", "php_5_all_modules"))
				app.SetEnv("COMPOSER_GITHUB_OAUTH_TOKEN", os.Getenv("COMPOSER_GITHUB_OAUTH_TOKEN"))
			})

			It("warns about deprecated PHP_EXTENSIONS", func() {
				PushAppAndConfirm(app)
				Expect(app.Stdout.String()).To(ContainSubstring("Warning: PHP_EXTENSIONS in options.json is deprecated."))

				ItLoadsAllTheModules("5")
			})
		})

		// context 'deploying a basic PHP7.0 app that loads all prepackaged extensions' do
		//   before(:all) do
		//     @app = setup_app('php_7_all_modules')
		//   end

		//   after(:all) do
		//     Machete::CF::DeleteApp.new.execute(@app)
		//   end

		//   it_behaves_like 'it loads all the modules', 'PHP 7.0'

		//   it 'warns about deprecated PHP_EXTENSIONS' do
		//     expect(@app).to be_running
		//     expect(@app).to have_logged 'Warning: PHP_EXTENSIONS in options.json is deprecated.'
		//   end
		// end

		// context 'deploying a basic PHP7.1 app that loads all prepackaged extensions' do
		//   before(:all) do
		//     @app = setup_app('php_71_all_modules')
		//   end

		//   after(:all) do
		//     Machete::CF::DeleteApp.new.execute(@app)
		//   end

		//   it_behaves_like 'it loads all the modules', 'PHP 7.1'

		//   it 'warns about deprecated PHP_EXTENSIONS' do
		//     expect(@app).to be_running
		//     expect(@app).to have_logged 'Warning: PHP_EXTENSIONS in options.json is deprecated.'
		//   end
		// end
	})

	// context 'extensions are specified in composer.json' do
	//   context 'deploying a basic PHP5 app that loads all prepackaged extensions' do
	//     before(:all) do
	//       @app = setup_app('php_5_all_modules_composer')
	//     end

	//     after(:all) do
	//       Machete::CF::DeleteApp.new.execute(@app)
	//     end

	//     it_behaves_like 'it loads all the modules', 'PHP 5'

	//     it 'does not warn about deprecated PHP_EXTENSIONS' do
	//       expect(@app).to be_running
	//       expect(@app).to_not have_logged 'Warning: PHP_EXTENSIONS in options.json is deprecated.'
	//     end
	//   end

	//   context 'deploying a basic PHP7.0 app that loads all prepackaged extensions' do
	//     before(:all) do
	//       @app = setup_app('php_7_all_modules_composer')
	//     end

	//     after(:all) do
	//       Machete::CF::DeleteApp.new.execute(@app)
	//     end

	//     it_behaves_like 'it loads all the modules', 'PHP 7.0'

	//     it 'does not warn about deprecated PHP_EXTENSIONS' do
	//       expect(@app).to be_running
	//       expect(@app).to_not have_logged 'Warning: PHP_EXTENSIONS in options.json is deprecated.'
	//     end
	//   end

	//   context 'deploying a basic PHP7.1 app that loads all prepackaged extensions' do
	//     before(:all) do
	//       @app = setup_app('php_71_all_modules_composer')
	//     end

	//     after(:all) do
	//       Machete::CF::DeleteApp.new.execute(@app)
	//     end

	//     it_behaves_like 'it loads all the modules', 'PHP 7.1'

	//     it 'does not warn about deprecated PHP_EXTENSIONS' do
	//       expect(@app).to be_running
	//       expect(@app).to_not have_logged 'Warning: PHP_EXTENSIONS in options.json is deprecated.'
	//     end
	//   end
	// end
})
