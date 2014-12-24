package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/goamz/ec2"
	// "github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	// "github.com/hashicorp/terraform/helper/schema"
)

func TestAccAWSNetworkAclsWithEgressAndIngressRules(t *testing.T) {
	var networkAcl ec2.NetworkAcl

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSNetworkAclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSNetworkAclEgressNIngressConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSNetworkAclExists("aws_network_acl.bar", &networkAcl),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "ingress.580214135.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "ingress.580214135.rule_no", "1"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "ingress.580214135.from_port", "80"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "ingress.580214135.to_port", "80"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "ingress.580214135.action", "allow"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "ingress.580214135.cidr_block", "10.3.10.3/18"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "egress.1730430240.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "egress.1730430240.rule_no", "2"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "egress.1730430240.from_port", "443"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "egress.1730430240.to_port", "443"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "egress.1730430240.cidr_block", "10.3.2.3/18"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.bar", "egress.1730430240.action", "allow"),
				),
			},
		},
	})
}

func TestAccAWSNetworkAclsOnlyIngressRules(t *testing.T) {
	var networkAcl ec2.NetworkAcl

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSNetworkAclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSNetworkAclIngressConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSNetworkAclExists("aws_network_acl.foos", &networkAcl),
					// testAccCheckSubnetAssociation("aws_network_acl.foos", "aws_subnet.blob"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.rule_no", "1"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.from_port", "0"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.to_port", "22"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.action", "deny"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.cidr_block", "10.2.2.3/18"),
				),
			},
		},
	})
}

func TestAccAWSNetworkAclsOnlyIngressRulesChange(t *testing.T) {
	var networkAcl ec2.NetworkAcl

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSNetworkAclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSNetworkAclIngressConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSNetworkAclExists("aws_network_acl.foos", &networkAcl),
					testIngressRuleLength(&networkAcl, 2),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.rule_no", "1"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.from_port", "0"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.to_port", "22"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.action", "deny"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.cidr_block", "10.2.2.3/18"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.2438803013.from_port", "443"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.2438803013.rule_no", "2"),
				),
			},
			resource.TestStep{
				Config: testAccAWSNetworkAclIngressConfigChange,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSNetworkAclExists("aws_network_acl.foos", &networkAcl),
					testIngressRuleLength(&networkAcl, 1),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.rule_no", "1"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.from_port", "0"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.to_port", "22"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.action", "deny"),
					resource.TestCheckResourceAttr(
						"aws_network_acl.foos", "ingress.3697634361.cidr_block", "10.2.2.3/18"),
				),
			},
		},
	})
}

func TestAccAWSNetworkAclsOnlyEgressRules(t *testing.T) {
	var networkAcl ec2.NetworkAcl

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSNetworkAclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSNetworkAclEgressConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSNetworkAclExists("aws_network_acl.bond", &networkAcl),
					testAccCheckTags(&networkAcl.Tags, "foo", "bar"),
				),
			},
		},
	})
}

func TestAccNetworkAcl_SubnetChange(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSNetworkAclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSNetworkAclSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetIsAssociatedWithAcl("aws_network_acl.bar", "aws_subnet.old"),
				),
			},
			resource.TestStep{
				Config: testAccAWSNetworkAclSubnetConfigChange,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetIsNotAssociatedWithAcl("aws_network_acl.bar", "aws_subnet.old"),
					testAccCheckSubnetIsAssociatedWithAcl("aws_network_acl.bar", "aws_subnet.new"),
				),
			},
		},
	})

}

func testAccCheckAWSNetworkAclDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).ec2conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_network" {
			continue
		}

		// Retrieve the network acl
		resp, err := conn.NetworkAcls([]string{rs.Primary.ID}, ec2.NewFilter())
		if err == nil {
			if len(resp.NetworkAcls) > 0 && resp.NetworkAcls[0].NetworkAclId == rs.Primary.ID {
				return fmt.Errorf("Network Acl (%s) still exists.", rs.Primary.ID)
			}

			return nil
		}

		ec2err, ok := err.(*ec2.Error)
		if !ok {
			return err
		}
		// Confirm error code is what we want
		if ec2err.Code != "InvalidNetworkAclID.NotFound" {
			return err
		}
	}

	return nil
}

func testAccCheckAWSNetworkAclExists(n string, networkAcl *ec2.NetworkAcl) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Security Group is set")
		}
		conn := testAccProvider.Meta().(*AWSClient).ec2conn

		resp, err := conn.NetworkAcls([]string{rs.Primary.ID}, nil)
		if err != nil {
			return err
		}

		if len(resp.NetworkAcls) > 0 && resp.NetworkAcls[0].NetworkAclId == rs.Primary.ID {
			*networkAcl = resp.NetworkAcls[0]
			return nil
		}

		return fmt.Errorf("Network Acls not found")
	}
}

func testIngressRuleLength(networkAcl *ec2.NetworkAcl, length int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var ingressEntries []ec2.NetworkAclEntry
		for _, e := range networkAcl.EntrySet {
			if e.Egress == false {
				ingressEntries = append(ingressEntries, e)
			}
		}
		// There is always a default rule (ALL Traffic ... DENY)
		// so we have to increase the lenght by 1
		if len(ingressEntries) != length+1 {
			return fmt.Errorf("Invalid number of ingress entries found; count = %d", len(ingressEntries))
		}
		return nil
	}
}

func testAccCheckSubnetIsAssociatedWithAcl(acl string, sub string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		networkAcl := s.RootModule().Resources[acl]
		subnet := s.RootModule().Resources[sub]

		conn := testAccProvider.Meta().(*AWSClient).ec2conn
		filter := ec2.NewFilter()
		filter.Add("association.subnet-id", subnet.Primary.ID)
		resp, err := conn.NetworkAcls([]string{networkAcl.Primary.ID}, filter)

		if err != nil {
			return err
		}
		if len(resp.NetworkAcls) > 0 {
			return nil
		}

		r, _ := conn.NetworkAcls([]string{}, ec2.NewFilter())
		fmt.Printf("\n\nall acls\n %#v\n\n", r.NetworkAcls)
		conn.NetworkAcls([]string{}, filter)

		return fmt.Errorf("Network Acl %s is not associated with subnet %s", acl, sub)
	}
}

func testAccCheckSubnetIsNotAssociatedWithAcl(acl string, subnet string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		networkAcl := s.RootModule().Resources[acl]
		subnet := s.RootModule().Resources[subnet]

		conn := testAccProvider.Meta().(*AWSClient).ec2conn
		filter := ec2.NewFilter()
		filter.Add("association.subnet-id", subnet.Primary.ID)
		resp, err := conn.NetworkAcls([]string{networkAcl.Primary.ID}, filter)

		if err != nil {
			return err
		}
		if len(resp.NetworkAcls) > 0 {
			return fmt.Errorf("Network Acl %s is still associated with subnet %s", acl, subnet)
		}
		return nil
	}
}

const testAccAWSNetworkAclIngressConfig = `
resource "aws_vpc" "foo" {
	cidr_block = "10.1.0.0/16"
}
resource "aws_subnet" "blob" {
	cidr_block = "10.1.1.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_network_acl" "foos" {
	vpc_id = "${aws_vpc.foo.id}"
	ingress = {
		protocol = "tcp"
		rule_no = 1
		action = "deny"
		cidr_block =  "10.2.2.3/18"
		from_port = 0
		to_port = 22
	}
	ingress = {
		protocol = "tcp"
		rule_no = 2
		action = "deny"
		cidr_block =  "10.2.2.3/18"
		from_port = 443
		to_port = 443
	}
	subnet_id = "${aws_subnet.blob.id}"
}
`
const testAccAWSNetworkAclIngressConfigChange = `
resource "aws_vpc" "foo" {
	cidr_block = "10.1.0.0/16"
}
resource "aws_subnet" "blob" {
	cidr_block = "10.1.1.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_network_acl" "foos" {
	vpc_id = "${aws_vpc.foo.id}"
	ingress = {
		protocol = "tcp"
		rule_no = 1
		action = "deny"
		cidr_block =  "10.2.2.3/18"
		from_port = 0
		to_port = 22
	}
	subnet_id = "${aws_subnet.blob.id}"
}
`

const testAccAWSNetworkAclEgressConfig = `
resource "aws_vpc" "foo" {
	cidr_block = "10.2.0.0/16"
}
resource "aws_subnet" "blob" {
	cidr_block = "10.2.0.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_network_acl" "bond" {
	vpc_id = "${aws_vpc.foo.id}"
	egress = {
		protocol = "tcp"
		rule_no = 2
		action = "allow"
		cidr_block =  "10.2.2.3/18"
		from_port = 443
		to_port = 443
	}

	egress = {
		protocol = "tcp"
		rule_no = 1
		action = "allow"
		cidr_block =  "10.2.10.3/18"
		from_port = 80
		to_port = 80
	}

	egress = {
		protocol = "tcp"
		rule_no = 3
		action = "allow"
		cidr_block =  "10.2.10.3/18"
		from_port = 22
		to_port = 22
	}

	tags {
		foo = "bar"
	}
}
`

const testAccAWSNetworkAclEgressNIngressConfig = `
resource "aws_vpc" "foo" {
	cidr_block = "10.3.0.0/16"
}
resource "aws_subnet" "blob" {
	cidr_block = "10.3.0.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_network_acl" "bar" {
	vpc_id = "${aws_vpc.foo.id}"
	egress = {
		protocol = "tcp"
		rule_no = 2
		action = "allow"
		cidr_block =  "10.3.2.3/18"
		from_port = 443
		to_port = 443
	}

	ingress = {
		protocol = "tcp"
		rule_no = 1
		action = "allow"
		cidr_block =  "10.3.10.3/18"
		from_port = 80
		to_port = 80
	}
}
`
const testAccAWSNetworkAclSubnetConfig = `
resource "aws_vpc" "foo" {
	cidr_block = "10.1.0.0/16"
}
resource "aws_subnet" "old" {
	cidr_block = "10.1.111.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_subnet" "new" {
	cidr_block = "10.1.1.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_network_acl" "roll" {
	vpc_id = "${aws_vpc.foo.id}"
	subnet_id = "${aws_subnet.new.id}"
}
resource "aws_network_acl" "bar" {
	vpc_id = "${aws_vpc.foo.id}"
	subnet_id = "${aws_subnet.old.id}"
}
`

const testAccAWSNetworkAclSubnetConfigChange = `
resource "aws_vpc" "foo" {
	cidr_block = "10.1.0.0/16"
}
resource "aws_subnet" "old" {
	cidr_block = "10.1.111.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_subnet" "new" {
	cidr_block = "10.1.1.0/24"
	vpc_id = "${aws_vpc.foo.id}"
	map_public_ip_on_launch = true
}
resource "aws_network_acl" "bar" {
	vpc_id = "${aws_vpc.foo.id}"
	subnet_id = "${aws_subnet.new.id}"
}
`
