package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func TestAccEC2EBSEncryptionByDefault_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ebs_encryption_by_default.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckEncryptionByDefaultDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEBSEncryptionByDefaultConfig_basic(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEBSEncryptionByDefault(ctx, resourceName, false),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccEBSEncryptionByDefaultConfig_basic(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEBSEncryptionByDefault(ctx, resourceName, true),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckEncryptionByDefaultDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

		response, err := conn.GetEbsEncryptionByDefaultWithContext(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
		if err != nil {
			return err
		}

		if aws.BoolValue(response.EbsEncryptionByDefault) != false {
			return fmt.Errorf("EBS encryption by default not disabled on resource removal")
		}

		return nil
	}
}

func testAccCheckEBSEncryptionByDefault(ctx context.Context, n string, enabled bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

		response, err := conn.GetEbsEncryptionByDefaultWithContext(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
		if err != nil {
			return err
		}

		if aws.BoolValue(response.EbsEncryptionByDefault) != enabled {
			return fmt.Errorf("EBS encryption by default is not in expected state (%t)", enabled)
		}

		return nil
	}
}

func testAccEBSEncryptionByDefaultConfig_basic(enabled bool) string {
	return fmt.Sprintf(`
resource "aws_ebs_encryption_by_default" "test" {
  enabled = %[1]t
}
`, enabled)
}
