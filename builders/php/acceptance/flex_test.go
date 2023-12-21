package acceptance_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/buildpacks/internal/acceptance"
)

const (
	flex = "google.config.flex"
)

func TestAcceptance(t *testing.T) {
	imageCtx, cleanup := acceptance.ProvisionImages(t)
	t.Cleanup(cleanup)

	testCases := []acceptance.Test{
		{
			Name:                       "modified front_controller",
			App:                        "front_controller",
			VersionInclusionConstraint: "< 8.2.0",
			MustUse:                    []string{flex, composer},
		},
		{
			Name:                       "php ini override",
			App:                        "php_ini",
			MustMatch:                  "PASS_PHP_INI",
			VersionInclusionConstraint: "< 8.2.0",
			MustUse:                    []string{flex},
		},
		{
			Name:                       "nginx server config included",
			App:                        "nginx_conf_include",
			VersionInclusionConstraint: "< 8.2.0",
			MustUse:                    []string{flex, utilsNginx},
			MustMatch:                  "app.php",
		},
	}

	for _, tc := range acceptance.FilterTests(t, imageCtx, testCases) {
		tc := tc
		if tc.Name == "" {
			tc.Name = tc.App
		}
		tc.Env = append(tc.Env, []string{"X_GOOGLE_TARGET_PLATFORM=flex", "GAE_APPLICATION_YAML_PATH=app.yaml"}...)

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			acceptance.TestApp(t, imageCtx, tc)
		})
	}
}
