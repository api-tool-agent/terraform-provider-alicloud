package alicloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAliCloudDTSSynchronizationJob_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_dts_synchronization_job.default"
	ra := resourceAttrInit(resourceId, AliCloudDTSSynchronizationJobMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DtsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDtsSynchronizationJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sdtssynchronizationjob%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudDTSSynchronizationJobBasicDependence0)
	synchronizationConfigValue := `[{\"module\":\"03\",\"name\":\"sink.batch.size.minimum\",\"value\":\"64\"},{\"module\":\"03\",\"name\":\"sink.task.number\",\"value\":\"4\"}]`

	expectedSynchronizationConfigValue := `[{"module":"03","name":"sink.batch.size.minimum","value":"64"},{"module":"03","name":"sink.task.number","value":"4"}]`
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"dts_instance_id":                    "${alicloud_dts_synchronization_instance.default.id}",
					"dts_job_name":                       "tf-testAccCase",
					"source_endpoint_instance_type":      "RDS",
					"source_endpoint_instance_id":        "${alicloud_db_instance.source.id}",
					"source_endpoint_engine_name":        "MySQL",
					"source_endpoint_region":             "${var.region_id}",
					"source_endpoint_database_name":      "test_database",
					"source_endpoint_user_name":          "${alicloud_rds_account.source_account.account_name}",
					"source_endpoint_password":           "${alicloud_rds_account.source_account.account_password}",
					"destination_endpoint_instance_type": "RDS",
					"destination_endpoint_instance_id":   "${alicloud_db_instance.target.id}",
					"destination_endpoint_engine_name":   "MySQL",
					"destination_endpoint_region":        "${var.region_id}",
					"destination_endpoint_database_name": "test_database",
					"destination_endpoint_user_name":     "${alicloud_rds_account.target_account.account_name}",
					"destination_endpoint_password":      "${alicloud_rds_account.target_account.account_password}",
					"db_list":                            "{\\\"test_database\\\":{\\\"name\\\":\\\"test_database\\\",\\\"all\\\":true,\\\"state\\\":\\\"normal\\\"}}",
					"structure_initialization":           "true",
					"data_initialization":                "true",
					"data_synchronization":               "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dts_job_name":                       "tf-testAccCase",
						"source_endpoint_instance_type":      "RDS",
						"source_endpoint_engine_name":        "MySQL",
						"source_endpoint_region":             os.Getenv("ALICLOUD_REGION"),
						"destination_endpoint_instance_type": "RDS",
						"destination_endpoint_engine_name":   "MySQL",
						"destination_endpoint_region":        os.Getenv("ALICLOUD_REGION"),
						"db_list":                            "{\"test_database\":{\"name\":\"test_database\",\"all\":true,\"state\":\"normal\"}}",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"dts_job_name": "tf-testAccCase1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dts_job_name": "tf-testAccCase1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"job_parameters": synchronizationConfigValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"job_parameters": expectedSynchronizationConfigValue,
					}),
				),
			},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"status": "Suspending",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"status": "Suspending",
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"status": "Synchronizing",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"status": "Synchronizing",
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"source_endpoint_password": "Lazypeople123+",
			//		"status":                   "Suspending",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"source_endpoint_password": "Lazypeople123+",
			//			"status":                   "Suspending",
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"source_endpoint_password": "${alicloud_rds_account.source_account.account_password}",
			//		"status":                   "Synchronizing",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"source_endpoint_password": CHECKSET,
			//			"status":                   "Synchronizing",
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"destination_endpoint_password": "Lazypeople123+",
			//		"status":                        "Retrying",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"destination_endpoint_password": "Lazypeople123+",
			//			"status":                        "Retrying",
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"destination_endpoint_password": "${alicloud_rds_account.target_account.account_password}",
			//		"status":                        "Synchronizing",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"destination_endpoint_password": CHECKSET,
			//			"status":                        "Synchronizing",
			//		}),
			//	),
			//},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"delay_notice", "error_phone", "delay_rule_time", "error_notice", "delay_phone", "reserve", "destination_endpoint_password", "source_endpoint_password"},
			},
		},
	})
}

func TestAccAliCloudDTSSynchronizationJob_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_dts_synchronization_job.default"
	ra := resourceAttrInit(resourceId, AliCloudDTSSynchronizationJobMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DtsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDtsSynchronizationJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sdtssynchronizationjob%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudDTSSynchronizationJobBasicDependence1)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"dts_instance_id":                    "${alicloud_dts_synchronization_instance.default.id}",
					"dts_job_name":                       "tf-testAccCase",
					"source_endpoint_instance_type":      "PolarDB",
					"source_endpoint_instance_id":        "${alicloud_polardb_cluster.source.id}",
					"source_endpoint_engine_name":        "PolarDB",
					"source_endpoint_region":             "${var.region_id}",
					"source_endpoint_database_name":      "test_database",
					"source_endpoint_user_name":          "${alicloud_polardb_account.source_account.account_name}",
					"source_endpoint_password":           "${alicloud_polardb_account.source_account.account_password}",
					"destination_endpoint_instance_type": "RDS",
					"destination_endpoint_instance_id":   "${alicloud_db_instance.target.id}",
					"destination_endpoint_engine_name":   "MySQL",
					"destination_endpoint_region":        "${var.region_id}",
					"destination_endpoint_database_name": "test_database",
					"destination_endpoint_user_name":     "${alicloud_rds_account.target_account.account_name}",
					"destination_endpoint_password":      "${alicloud_rds_account.target_account.account_password}",
					"db_list":                            "{\\\"tfaccountpri_0\\\":{\\\"name\\\":\\\"tfaccountpri_0\\\",\\\"all\\\":true,\\\"state\\\":\\\"normal\\\"}}",
					"structure_initialization":           "true",
					"data_initialization":                "true",
					"data_synchronization":               "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dts_job_name":                       "tf-testAccCase",
						"source_endpoint_instance_type":      "PolarDB",
						"source_endpoint_engine_name":        "PolarDB",
						"source_endpoint_region":             os.Getenv("ALICLOUD_REGION"),
						"destination_endpoint_instance_type": "RDS",
						"destination_endpoint_engine_name":   "MySQL",
						"destination_endpoint_region":        os.Getenv("ALICLOUD_REGION"),
						"db_list":                            "{\"tfaccountpri_0\":{\"name\":\"tfaccountpri_0\",\"all\":true,\"state\":\"normal\"}}",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"delay_notice", "error_phone", "delay_rule_time", "error_notice", "delay_phone", "reserve", "destination_endpoint_password", "source_endpoint_password"},
			},
		},
	})
}

func TestAccAliCloudDTSSynchronizationJob_basic2(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_dts_synchronization_job.default"
	ra := resourceAttrInit(resourceId, AliCloudDTSSynchronizationJobMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DtsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDtsSynchronizationJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sdtssynchronizationjob%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudDTSSynchronizationJobBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"dts_instance_id":                    "${alicloud_dts_synchronization_instance.default.id}",
					"dts_job_name":                       "tf-testAccCase",
					"source_endpoint_instance_type":      "RDS",
					"source_endpoint_instance_id":        "${alicloud_db_instance.source.id}",
					"source_endpoint_engine_name":        "MySQL",
					"source_endpoint_region":             "${var.region_id}",
					"source_endpoint_database_name":      "test_database",
					"source_endpoint_user_name":          "${alicloud_rds_account.source_account.account_name}",
					"source_endpoint_password":           "${alicloud_rds_account.source_account.account_password}",
					"destination_endpoint_instance_type": "RDS",
					"destination_endpoint_instance_id":   "${alicloud_db_instance.target.id}",
					"destination_endpoint_engine_name":   "MySQL",
					"destination_endpoint_region":        "${var.region_id}",
					"destination_endpoint_database_name": "test_database",
					"destination_endpoint_user_name":     "${alicloud_rds_account.target_account.account_name}",
					"destination_endpoint_password":      "${alicloud_rds_account.target_account.account_password}",
					"db_list":                            "{\\\"test_database\\\":{\\\"name\\\":\\\"test_database\\\",\\\"all\\\":true,\\\"state\\\":\\\"normal\\\"}}",
					"structure_initialization":           "true",
					"data_initialization":                "true",
					"data_synchronization":               "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dts_job_name":                       "tf-testAccCase",
						"source_endpoint_instance_type":      "RDS",
						"source_endpoint_engine_name":        "MySQL",
						"source_endpoint_region":             os.Getenv("ALICLOUD_REGION"),
						"destination_endpoint_instance_type": "RDS",
						"destination_endpoint_engine_name":   "MySQL",
						"destination_endpoint_region":        os.Getenv("ALICLOUD_REGION"),
						"db_list":                            "{\"test_database\":{\"name\":\"test_database\",\"all\":true,\"state\":\"normal\"}}",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_list": "{\\\"test_database\\\":{\\\"name\\\":\\\"test_database\\\",\\\"all\\\":true,\\\"state\\\":\\\"normal\\\"}}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_list": "{\"test_database\":{\"name\":\"test_database\",\"all\":true,\"state\":\"normal\"}}",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"delay_notice", "error_phone", "delay_rule_time", "error_notice", "delay_phone", "reserve", "destination_endpoint_password", "source_endpoint_password"},
			},
		},
	})
}

var AliCloudDTSSynchronizationJobMap0 = map[string]string{
	"error_phone":                      NOSET,
	"error_notice":                     NOSET,
	"delay_rule_time":                  NOSET,
	"delay_phone":                      NOSET,
	"source_endpoint_engine_name":      CHECKSET,
	"reserve":                          NOSET,
	"delay_notice":                     NOSET,
	"destination_endpoint_engine_name": CHECKSET,
	"status":                           CHECKSET,
}

func TestAccAliCloudDTSSynchronizationJob_ssl(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_dts_synchronization_job.default"
	ra := resourceAttrInit(resourceId, AliCloudDTSSynchronizationJobMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DtsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDtsSynchronizationJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testaccdtssyncjobssl%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudDTSSynchronizationJobSslDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"dts_instance_id":                    "${alicloud_dts_synchronization_instance.default.id}",
					"dts_job_name":                       "tf-testAccCaseSsl",
					"source_endpoint_instance_type":      "RDS",
					"source_endpoint_instance_id":        "${alicloud_db_instance.source.id}",
					"source_endpoint_engine_name":        "MySQL",
					"source_endpoint_region":             "${data.alicloud_regions.default.regions.0.id}",
					"source_endpoint_database_name":      "test_database",
					"source_endpoint_user_name":          "${alicloud_rds_account.source_account.account_name}",
					"source_endpoint_password":           "${alicloud_rds_account.source_account.account_password}",
					"destination_endpoint_instance_type": "RDS",
					"destination_endpoint_instance_id":   "${alicloud_db_instance.target.id}",
					"destination_endpoint_engine_name":   "MySQL",
					"destination_endpoint_region":        "${data.alicloud_regions.default.regions.0.id}",
					"destination_endpoint_database_name": "test_database",
					"destination_endpoint_user_name":     "${alicloud_rds_account.target_account.account_name}",
					"destination_endpoint_password":      "${alicloud_rds_account.target_account.account_password}",
					"db_list":                            "{\\\"test_database\\\":{\\\"name\\\":\\\"test_database\\\",\\\"all\\\":true,\\\"state\\\":\\\"normal\\\"}}",
					// reserve is set alongside the SSL attributes so that the srcSSL/destSSL keys
					// have to be merged into an existing reserve rather than composed from scratch.
					// A merge that dropped targetTableMode would change how the job treats existing
					// destination tables, and one that produced invalid JSON would fail to configure.
					"reserve":                  "{\\\"targetTableMode\\\":\\\"2\\\"}",
					"source_endpoint_ssl":      "1",
					"destination_endpoint_ssl": "1",
					"structure_initialization": "true",
					"data_initialization":      "true",
					"data_synchronization":     "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"source_endpoint_ssl":      "1",
						"destination_endpoint_ssl": "1",
						// The merge happens on the request only; state must still hold the
						// reserve string exactly as it was written, with no SSL keys folded in.
						"reserve": "{\"targetTableMode\":\"2\"}",
					}),
				),
			},
			// Turn SSL off on both endpoints in place. The fields are not ForceNew, so this must
			// go through ModifyDtsJob with ModifyTypeEnum=UPDATE_RESERVED rather than recreating
			// the job.
			{
				Config: testAccConfig(map[string]interface{}{
					"source_endpoint_ssl":      "0",
					"destination_endpoint_ssl": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"source_endpoint_ssl":      "0",
						"destination_endpoint_ssl": "0",
					}),
				),
			},
			// Turn SSL back on, covering the enable-by-update direction as well as disable.
			{
				Config: testAccConfig(map[string]interface{}{
					"source_endpoint_ssl":      "1",
					"destination_endpoint_ssl": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"source_endpoint_ssl":      "1",
						"destination_endpoint_ssl": "1",
					}),
				),
			},
			// Change only one endpoint, so the Reserved payload carries a single changed key.
			{
				Config: testAccConfig(map[string]interface{}{
					"source_endpoint_ssl": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"source_endpoint_ssl":      "0",
						"destination_endpoint_ssl": "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"delay_notice", "error_phone", "delay_rule_time", "error_notice", "delay_phone", "reserve", "destination_endpoint_password", "source_endpoint_password"},
			},
		},
	})
}

func AliCloudDTSSynchronizationJobSslDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_regions" "default" {
  		current = true
	}

	data "alicloud_db_zones" "default" {
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		instance_charge_type     = "PostPaid"
  		category                 = "HighAvailability"
  		db_instance_storage_type = "cloud_essd"
	}

	data "alicloud_db_instance_classes" "default" {
  		zone_id                  = data.alicloud_db_zones.default.zones.0.id
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		category                 = "HighAvailability"
  		db_instance_storage_type = "cloud_essd"
  		instance_charge_type     = "PostPaid"
	}

	resource "alicloud_vpc" "default" {
  		vpc_name   = var.name
  		cidr_block = "172.16.0.0/16"
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = "172.16.0.0/24"
  		zone_id      = data.alicloud_db_zones.default.zones.0.id
  		vswitch_name = var.name
	}

	## RDS MySQL Source, with SSL opened
	resource "alicloud_db_instance" "source" {
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		instance_type            = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  		instance_storage         = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
  		db_instance_storage_type = "cloud_essd"
  		vswitch_id               = alicloud_vswitch.default.id
  		instance_name            = "rds-mysql-source-ssl"
  		ssl_action               = "Open"
	}

	resource "alicloud_db_database" "source_db" {
  		instance_id = alicloud_db_instance.source.id
  		name        = "test_database"
	}

	resource "alicloud_rds_account" "source_account" {
  		db_instance_id   = alicloud_db_instance.source.id
  		account_name     = "test_mysql"
  		account_password = "N1cetest"
	}

	resource "alicloud_db_account_privilege" "source_privilege" {
  		instance_id  = alicloud_db_instance.source.id
  		account_name = alicloud_rds_account.source_account.name
  		privilege    = "ReadWrite"
  		db_names     = alicloud_db_database.source_db.*.name
	}

	## RDS MySQL Target, with SSL opened
	resource "alicloud_db_instance" "target" {
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		instance_type            = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  		instance_storage         = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
  		db_instance_storage_type = "cloud_essd"
  		vswitch_id               = alicloud_vswitch.default.id
  		instance_name            = "rds-mysql-target-ssl"
  		ssl_action               = "Open"
	}

	resource "alicloud_rds_account" "target_account" {
  		db_instance_id   = alicloud_db_instance.target.id
  		account_name     = "test_mysql"
  		account_password = "N1cetest"
	}

	## DTS Data Synchronization
	resource "alicloud_dts_synchronization_instance" "default" {
  		payment_type                     = "PayAsYouGo"
  		source_endpoint_engine_name      = "MySQL"
  		source_endpoint_region           = data.alicloud_regions.default.regions.0.id
  		destination_endpoint_engine_name = "MySQL"
  		destination_endpoint_region      = data.alicloud_regions.default.regions.0.id
  		instance_class                   = "4xlarge"
  		sync_architecture                = "oneway"
	}
`, name)
}

func AliCloudDTSSynchronizationJobBasicDependence0(name string) string {
	return fmt.Sprintf(` 
	variable "name" {
  		default = "%s"
	}

	variable "region_id" {
  		default = "%s"
	}

	data "alicloud_db_zones" "default" {
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		instance_charge_type     = "PostPaid"
  		category                 = "HighAvailability"
  		db_instance_storage_type = "cloud_essd"
	}

	data "alicloud_vpcs" "default" {
  		name_regex =  "^default-NODELETING$"
	}

	data "alicloud_vswitches" "default" {
  		vpc_id  = data.alicloud_vpcs.default.ids.0
  		zone_id = data.alicloud_db_zones.default.zones.0.id
	}

	data "alicloud_db_instance_classes" "default" {
  		zone_id                  = data.alicloud_db_zones.default.zones.0.id
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		category                 = "HighAvailability"
  		db_instance_storage_type = "cloud_essd"
  		instance_charge_type     = "PostPaid"
	}

	## RDS MySQL Source
	resource "alicloud_db_instance" "source" {
  		engine           = "MySQL"
  		engine_version   = "8.0"
  		instance_type    = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  		instance_storage = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
  		vswitch_id       = data.alicloud_vswitches.default.ids.0
  		instance_name    = "rds-mysql-source"
	}

	resource "alicloud_db_database" "source_db" {
  		instance_id = alicloud_db_instance.source.id
  		name        = "test_database"
	}

	resource "alicloud_rds_account" "source_account" {
  		db_instance_id   = alicloud_db_instance.source.id
  		account_name     = "test_mysql"
  		account_password = "N1cetest"
	}

	resource "alicloud_db_account_privilege" "source_privilege" {
  		instance_id  = alicloud_db_instance.source.id
  		account_name = alicloud_rds_account.source_account.name
  		privilege    = "ReadWrite"
  		db_names     = alicloud_db_database.source_db.*.name
	}

	## RDS MySQL Target
	resource "alicloud_db_instance" "target" {
  		engine           = "MySQL"
  		engine_version   = "8.0"
  		instance_type    = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  		instance_storage = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
  		vswitch_id       = data.alicloud_vswitches.default.ids.0
  		instance_name    = "rds-mysql-target"
	}

	resource "alicloud_rds_account" "target_account" {
  		db_instance_id   = alicloud_db_instance.target.id
  		account_name     = "test_mysql"
  		account_password = "N1cetest"
	}

	## DTS Data Synchronization
	resource "alicloud_dts_synchronization_instance" "default" {
  		payment_type                     = "PayAsYouGo"
  		source_endpoint_engine_name      = "MySQL"
  		source_endpoint_region           = var.region_id
  		destination_endpoint_engine_name = "MySQL"
  		destination_endpoint_region      = var.region_id
  		instance_class                   = "4xlarge"
  		sync_architecture                = "oneway"
	}
`, name, defaultRegionToTest)
}

func AliCloudDTSSynchronizationJobBasicDependence1(name string) string {
	return fmt.Sprintf(` 
	variable "name" {
  		default = "%s"
	}

	variable "region_id" {
  		default = "cn-hangzhou"
	}

	data "alicloud_db_zones" "default" {
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		instance_charge_type     = "PostPaid"
  		category                 = "HighAvailability"
  		db_instance_storage_type = "cloud_essd"
	}

	data "alicloud_vpcs" "default" {
  		name_regex =  "^default-NODELETING$"
	}

	data "alicloud_vswitches" "default" {
  		vpc_id  = data.alicloud_vpcs.default.ids.0
  		zone_id = data.alicloud_db_zones.default.zones.4.id
	}

	data "alicloud_db_instance_classes" "default" {
  		zone_id                  = data.alicloud_db_zones.default.zones.4.id
  		engine                   = "MySQL"
  		engine_version           = "8.0"
  		category                 = "HighAvailability"
  		db_instance_storage_type = "cloud_essd"
  		instance_charge_type     = "PostPaid"
	}

	data "alicloud_polardb_node_classes" "default" {
  		db_type    = "MySQL"
  		db_version = "8.0"
  		pay_type   = "PostPaid"
  		zone_id    = data.alicloud_db_zones.default.zones.4.id
	}

	## PolarDB PolarDB Source
	resource "alicloud_polardb_cluster" "source" {
  		db_type       = "MySQL"
  		db_version    = "8.0"
  		pay_type      = "PostPaid"
  		//db_node_class = data.alicloud_polardb_node_classes.default.classes.0.supported_engines.0.available_resources.0.db_node_class
  		db_node_class = "polar.mysql.x4.medium.c"
  		vswitch_id    = data.alicloud_vswitches.default.ids.0
  		description   = "polardb_cluster_description"
		storage_space = 20
		storage_type  = "ESSDPL0" 
	}

	resource "alicloud_polardb_database" "source_db" {
  		db_cluster_id = alicloud_polardb_cluster.source.id
  		db_name       = "test_database"
  		account_name  = "test_polardb"
	}

	resource "alicloud_polardb_account" "source_account" {
  		db_cluster_id    = alicloud_polardb_cluster.source.id
  		account_name     = "test_polardb"
  		account_password = "N1cetest"
	}

	resource "alicloud_polardb_account_privilege" "source_privilege" {
  		db_cluster_id     = alicloud_polardb_cluster.source.id
  		account_name      = alicloud_polardb_account.source_account.account_name
  		account_privilege = "ReadWrite"
  		db_names          = alicloud_polardb_database.source_db.*.db_name
	}

	## RDS MySQL Target
		resource "alicloud_db_instance" "target" {
  		engine           = "MySQL"
  		engine_version   = "8.0"
  		instance_type    = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  		instance_storage = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
  		vswitch_id       = data.alicloud_vswitches.default.ids.0
  		instance_name    = "rds-mysql-target"
	}

	resource "alicloud_rds_account" "target_account" {
  		db_instance_id   = alicloud_db_instance.target.id
  		account_name     = "test_mysql"
  		account_password = "N1cetest"
	}

	## DTS Data Synchronization
	resource "alicloud_dts_synchronization_instance" "default" {
  		payment_type                     = "PayAsYouGo"
  		source_endpoint_engine_name      = "PolarDB"
  		source_endpoint_region           = var.region_id
  		destination_endpoint_engine_name = "MySQL"
  		destination_endpoint_region      = var.region_id
  		instance_class                   = "4xlarge"
  		sync_architecture                = "oneway"
	}
`, name)
}

func TestUnitAlicloudDTSSynchronizationJobInstanceClassTransferTarget(t *testing.T) {
	cases := []struct {
		name         string
		configClass  interface{}
		actualClass  interface{}
		expectTarget string
	}{
		{"empty config skips", "", "small", ""},
		{"nil config skips", nil, "small", ""},
		{"equal class skips", "small", "small", ""},
		{"upgrade dispatches", "large", "small", "large"},
		{"downgrade dispatches", "small", "4xlarge", "small"},
		{"unknown actual dispatches", "large", "", "large"},
		{"nil actual dispatches", "large", nil, "large"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expectTarget, dtsSyncJobInstanceClassTransferTarget(c.configClass, c.actualClass))
		})
	}
}

func TestUnitAlicloudDTSSynchronizationJobInstanceClassValue(t *testing.T) {
	assert.Equal(t, nil, dtsSyncJobInstanceClassValue(map[string]interface{}{}))
	assert.Equal(t, nil, dtsSyncJobInstanceClassValue(map[string]interface{}{"DtsJobClass": nil}))
	assert.Equal(t, nil, dtsSyncJobInstanceClassValue(map[string]interface{}{"DtsJobClass": ""}))
	assert.Equal(t, "large", dtsSyncJobInstanceClassValue(map[string]interface{}{"DtsJobClass": "large"}))
}

func TestUnitAlicloudDTSSynchronizationJob(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	resSchema := schema.InternalMap(p["alicloud_dts_synchronization_job"].Schema)

	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
	}

	jobDetailResponse := func(dtsJobClass string) map[string]interface{} {
		return map[string]interface{}{
			"Success":     true,
			"DtsJobClass": dtsJobClass,
			"Status":      "Synchronizing",
			"MigrationMode": map[string]interface{}{
				"DataInitialization":      true,
				"DataSynchronization":     true,
				"StructureInitialization": true,
			},
			"DestinationEndpoint": map[string]interface{}{},
			"SourceEndpoint":      map[string]interface{}{},
		}
	}

	newResourceData := func(oldClass, newClass string) *schema.ResourceData {
		base := map[string]string{
			"dts_instance_id": "dtsi93b3q2p1t6d1****",
			"dts_job_name":    "unit-test-job",
			"db_list":         `{"unit":{"name":"unit","all":true}}`,
			"instance_class":  oldClass,
		}
		s := &terraform.InstanceState{ID: "dtsj93b3q2p1t6d1****", Attributes: base}
		newAttrs := map[string]string{}
		for k, v := range base {
			newAttrs[k] = v
		}
		newAttrs["instance_class"] = newClass
		config := map[string]interface{}{}
		for k, v := range newAttrs {
			config[k] = v
		}
		diff, err := resSchema.Diff(s, &terraform.ResourceConfig{Config: config}, nil, nil, false)
		assert.Nil(t, err)
		d, err := resSchema.Data(s, diff)
		assert.Nil(t, err)
		return d
	}

	mockSeams := func(dtsJobClass string) {
		dtsSyncJobDescribe = func(_ *connectivity.AliyunClient, _ string) (map[string]interface{}, error) {
			return jobDetailResponse(dtsJobClass), nil
		}
		dtsSyncJobQueryChangedParameters = func(_ *connectivity.AliyunClient, _ string) (string, error) {
			return "[]", nil
		}
	}
	restoreSeams := func() {
		dtsSyncJobDescribe = func(client *connectivity.AliyunClient, id string) (map[string]interface{}, error) {
			dtsService := DtsService{client}
			return dtsService.DescribeDtsSynchronizationJob(id)
		}
		dtsSyncJobRpcPost = func(client *connectivity.AliyunClient, apiProductCode string, apiVersion string, apiName string, query map[string]interface{}, body map[string]interface{}, autoRetry bool) (map[string]interface{}, error) {
			return client.RpcPost(apiProductCode, apiVersion, apiName, query, body, autoRetry)
		}
		dtsSyncJobQueryChangedParameters = func(client *connectivity.AliyunClient, id string) (string, error) {
			dtsService := DtsService{client}
			return dtsService.QueryChangedJobParameters(id)
		}
	}

	t.Run("update dispatches transfer when class differs", func(t *testing.T) {
		transferCalled := false
		dtsSyncJobDescribe = func(_ *connectivity.AliyunClient, _ string) (map[string]interface{}, error) {
			if transferCalled {
				return jobDetailResponse("large"), nil
			}
			return jobDetailResponse("small"), nil
		}
		dtsSyncJobQueryChangedParameters = func(_ *connectivity.AliyunClient, _ string) (string, error) {
			return "[]", nil
		}
		defer restoreSeams()
		var transferRequests []map[string]interface{}
		dtsSyncJobRpcPost = func(_ *connectivity.AliyunClient, _ string, _ string, apiName string, _ map[string]interface{}, body map[string]interface{}, _ bool) (map[string]interface{}, error) {
			if apiName == "TransferInstanceClass" {
				transferRequests = append(transferRequests, body)
				transferCalled = true
			}
			return map[string]interface{}{"Success": true}, nil
		}
		d := newResourceData("small", "large")
		err := resourceAlicloudDtsSynchronizationJobUpdate(d, rawClient)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(transferRequests))
		assert.Equal(t, "large", transferRequests[0]["InstanceClass"])
		assert.Equal(t, "UPGRADE", transferRequests[0]["OrderType"])
		assert.Equal(t, "large", d.Get("instance_class"))
	})

	t.Run("update skips transfer when config equals actual", func(t *testing.T) {
		mockSeams("small")
		defer restoreSeams()
		transferCalled := false
		dtsSyncJobRpcPost = func(_ *connectivity.AliyunClient, _ string, _ string, apiName string, _ map[string]interface{}, _ map[string]interface{}, _ bool) (map[string]interface{}, error) {
			if apiName == "TransferInstanceClass" {
				transferCalled = true
			}
			return map[string]interface{}{"Success": true}, nil
		}
		d := newResourceData("4xlarge", "small")
		err := resourceAlicloudDtsSynchronizationJobUpdate(d, rawClient)
		assert.Nil(t, err)
		assert.False(t, transferCalled)
		assert.Equal(t, "small", d.Get("instance_class"))
	})

	t.Run("update skips transfer without instance_class change", func(t *testing.T) {
		mockSeams("small")
		defer restoreSeams()
		describeCalls := 0
		dtsSyncJobDescribe = func(_ *connectivity.AliyunClient, _ string) (map[string]interface{}, error) {
			describeCalls++
			return jobDetailResponse("small"), nil
		}
		transferCalled := false
		dtsSyncJobRpcPost = func(_ *connectivity.AliyunClient, _ string, _ string, apiName string, _ map[string]interface{}, _ map[string]interface{}, _ bool) (map[string]interface{}, error) {
			if apiName == "TransferInstanceClass" {
				transferCalled = true
			}
			return map[string]interface{}{"Success": true}, nil
		}
		d := newResourceData("small", "small")
		err := resourceAlicloudDtsSynchronizationJobUpdate(d, rawClient)
		assert.Nil(t, err)
		assert.False(t, transferCalled)
		// describe is only invoked by the trailing read, never by the instance_class branch
		assert.Equal(t, 1, describeCalls)
	})

	t.Run("update returns error when describe fails", func(t *testing.T) {
		restoreSeams()
		defer restoreSeams()
		dtsSyncJobDescribe = func(_ *connectivity.AliyunClient, _ string) (map[string]interface{}, error) {
			return nil, fmt.Errorf("mock describe failure")
		}
		d := newResourceData("small", "large")
		err := resourceAlicloudDtsSynchronizationJobUpdate(d, rawClient)
		assert.NotNil(t, err)
	})

	t.Run("read writes back actual instance class", func(t *testing.T) {
		mockSeams("4xlarge")
		defer restoreSeams()
		d, err := resSchema.Data(&terraform.InstanceState{ID: "dtsj93b3q2p1t6d1****", Attributes: map[string]string{}}, nil)
		assert.Nil(t, err)
		err = resourceAlicloudDtsSynchronizationJobRead(d, rawClient)
		assert.Nil(t, err)
		assert.Equal(t, "4xlarge", d.Get("instance_class"))
	})

	t.Run("read keeps state when actual class is empty", func(t *testing.T) {
		mockSeams("")
		defer restoreSeams()
		d, err := resSchema.Data(&terraform.InstanceState{ID: "dtsj93b3q2p1t6d1****", Attributes: map[string]string{"instance_class": "large"}}, nil)
		assert.Nil(t, err)
		err = resourceAlicloudDtsSynchronizationJobRead(d, rawClient)
		assert.Nil(t, err)
		assert.Equal(t, "large", d.Get("instance_class"))
	})
}
