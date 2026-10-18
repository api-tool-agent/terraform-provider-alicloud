package paistudio_test

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/acctest"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/service/paistudio"

	tfacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// The managed alicloud_paistudio_algorithm is created and destroyed entirely
// through Terraform, so the test needs no out-of-band fixture; the standalone
// client exists only so CheckDestroy can prove the Algorithm is really gone
// after teardown. A plain `go test` (no TF_ACC) skips before any client is
// built or cloud call is made.
func TestAccAliCloudPaiStudioAlgorithm(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set; skipping PaiStudio acceptance test")
	}
	acctest.PreCheck(t)

	// PaiStudio (POP product "PaiStudio", 2022-01-12) is served from the unified
	// PAI endpoint pai.<region>.aliyuncs.com; it is not registered with the
	// location service and paistudio.<region>.aliyuncs.com does not resolve.
	// (client).loadEndpoint reads ALIBABA_CLOUD_ENDPOINT_<PRODUCT> first, so set
	// it here (before the first ROA call) for both this standalone client and
	// the muxed provider's client.
	region := os.Getenv("ALICLOUD_REGION")
	if region == "" {
		region = os.Getenv("ALIBABA_CLOUD_REGION")
	}
	if region != "" {
		os.Setenv("ALIBABA_CLOUD_ENDPOINT_PAISTUDIO", fmt.Sprintf("pai.%s.aliyuncs.com", region))
	}
	client, err := newTestPaiStudioClient()
	if err != nil {
		t.Fatalf("building PaiStudio client: %s", err)
	}

	rand := tfacctest.RandIntRange(1000, 9999)
	algorithmName := fmt.Sprintf("tf-testAccPaistudioAlgorithm-%d", rand)

	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		// CheckDestroy proves the managed alicloud_paistudio_algorithm itself is
		// gone after step teardown, not merely that the test returned. A
		// transient 5xx on the read-back is NOT treated as destroyed.
		CheckDestroy: testAccCheckPaiStudioAlgorithmDestroy(client),
		Steps: []resource.TestStep{
			{
				Config: testAccPaiStudioAlgorithmConfig(algorithmName, "tf-testAccPaistudioAlgorithm-display", "created by terraform acc"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_paistudio_algorithm.default", "algorithm_name", algorithmName),
					resource.TestCheckResourceAttr("alicloud_paistudio_algorithm.default", "display_name", "tf-testAccPaistudioAlgorithm-display"),
					resource.TestCheckResourceAttr("alicloud_paistudio_algorithm.default", "algorithm_description", "created by terraform acc"),
					resource.TestCheckResourceAttrSet("alicloud_paistudio_algorithm.default", "id"),
					resource.TestCheckResourceAttrSet("alicloud_paistudio_algorithm.default", "workspace_id"),
					resource.TestCheckResourceAttrSet("alicloud_paistudio_algorithm.default", "gmt_create_time"),
				),
			},
			{
				// UpdateAlgorithm accepts only the description and display name.
				Config: testAccPaiStudioAlgorithmConfig(algorithmName, "tf-testAccPaistudioAlgorithm-display-2", "updated by terraform acc"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_paistudio_algorithm.default", "algorithm_name", algorithmName),
					resource.TestCheckResourceAttr("alicloud_paistudio_algorithm.default", "display_name", "tf-testAccPaistudioAlgorithm-display-2"),
					resource.TestCheckResourceAttr("alicloud_paistudio_algorithm.default", "algorithm_description", "updated by terraform acc"),
				),
			},
			{
				// Every attribute is API-backed and refreshed on read, so the
				// imported state must match without ignores.
				Config:            testAccPaiStudioAlgorithmConfig(algorithmName, "tf-testAccPaistudioAlgorithm-display-2", "updated by terraform acc"),
				ResourceName:      "alicloud_paistudio_algorithm.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// testAccCheckPaiStudioAlgorithmDestroy verifies every
// alicloud_paistudio_algorithm in state is gone after destroy: a RoaGet that
// returns the typed NoSuchObject proves the algorithm is absent, an
// unexpected error (5xx/auth/malformed) is reported rather than mistaken for
// gone, and a 200 response means the resource leaked and fails the test.
func testAccCheckPaiStudioAlgorithmDestroy(client *connectivity.AliyunClient) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "alicloud_paistudio_algorithm" {
				continue
			}
			id := rs.Primary.ID
			_, err := client.RoaGet(paistudio.PaiStudioProduct, paistudio.PaiStudioVersion,
				fmt.Sprintf("/api/v1/algorithms/%s", id), nil, nil, nil)
			if err == nil {
				return fmt.Errorf("algorithm %s still exists after destroy", id)
			}
			if !isPaiStudioNotFound(err) {
				return fmt.Errorf("destroy check for algorithm %s: %w", id, err)
			}
		}
		return nil
	}
}

func testAccPaiStudioAlgorithmConfig(algorithmName, displayName, description string) string {
	return fmt.Sprintf(`
resource "alicloud_paistudio_algorithm" "default" {
  algorithm_name        = %[1]q
  display_name          = %[2]q
  algorithm_description = %[3]q
}
`, algorithmName, displayName, description)
}

// newTestPaiStudioClient builds a standalone AliyunClient from ALICLOUD_*
// env vars, mirroring alicloud_sweeper_test.go's sharedClientForRegion. The
// muxed provider meta is nil before the first step, so acctest.Provider.Meta()
// cannot be used before the first step completes.
func newTestPaiStudioClient() (*connectivity.AliyunClient, error) {
	protocol := os.Getenv("ALICLOUD_PROTOCOL")
	if protocol == "" {
		protocol = "HTTPS"
	}
	conf := connectivity.Config{
		Region:        connectivity.Region(os.Getenv("ALICLOUD_REGION")),
		RegionId:      os.Getenv("ALICLOUD_REGION"),
		AccessKey:     os.Getenv("ALICLOUD_ACCESS_KEY"),
		SecretKey:     os.Getenv("ALICLOUD_SECRET_KEY"),
		SecurityToken: os.Getenv("ALICLOUD_SECURITY_TOKEN"),
		Protocol:      protocol,
	}
	if conf.AccessKey == "" {
		conf.AccessKey = os.Getenv("ALIBABA_CLOUD_ACCESS_KEY")
	}
	if conf.SecretKey == "" {
		conf.SecretKey = os.Getenv("ALIBABA_CLOUD_SECRET_KEY")
	}
	if conf.RegionId == "" {
		conf.RegionId = os.Getenv("ALIBABA_CLOUD_REGION")
		conf.Region = connectivity.Region(conf.RegionId)
	}
	if conf.SecurityToken == "" {
		conf.SecurityToken = os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN")
	}

	accountId := os.Getenv("ALICLOUD_ACCOUNT_ID")
	if accountId == "" {
		accountId = os.Getenv("ALIBABA_CLOUD_ACCOUNT_ID")
	}
	if accountId != "" {
		conf.AccountId = accountId
	}
	accountType := os.Getenv("ALICLOUD_ACCOUNT_TYPE")
	if accountType == "" {
		accountType = os.Getenv("ALIBABA_CLOUD_ACCOUNT_TYPE")
	}
	if accountType != "" {
		conf.AccountType = accountType
	}

	// Config.Credential is dereferenced by getAuthCredential during the first
	// ROA call (describeEndpointForService). conf.Client() does not populate
	// it, so mirror alicloud_sweeper_test.go's sharedClientForRegion and set
	// it explicitly from the same ALICLOUD_* env to avoid a nil-pointer panic.
	credCfg := new(credentials.Config).SetType("access_key").
		SetAccessKeyId(conf.AccessKey).SetAccessKeySecret(conf.SecretKey)
	if strings.TrimSpace(conf.SecurityToken) != "" {
		credCfg.SetType("sts").SetSecurityToken(conf.SecurityToken)
	}
	cred, err := credentials.NewCredential(credCfg)
	if err != nil {
		return nil, fmt.Errorf("building PaiStudio credential: %w", err)
	}
	conf.Credential = cred

	conf.Endpoints = &sync.Map{}
	conf.SignVersion = &sync.Map{}
	return conf.Client()
}

// isPaiStudioNotFound reports whether err is the PaiStudio not-found error,
// using the product error contract (typed *tea.SDKError Code == NoSuchObject)
// only. A loose substring match is intentionally avoided so a non-NotFound
// message that merely contains the word "NoSuchObject" cannot be misread as
// gone and mask a leaked resource during destroy checks.
func isPaiStudioNotFound(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*tea.SDKError); ok && e.Code != nil && *e.Code == "NoSuchObject" {
		return true
	}
	return false
}
