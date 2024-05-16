import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";
import * as tls from "@pulumi/tls";
import * as random from "@pulumi/random";
import * as http from "http";

function fetchMyPublicIpOutput(): Promise<string> {
  return pulumi.output(
    new Promise((resolve, reject) => {
      http
        .get("http://api.ipify.org", (resp) => {
          let data = "";
          resp.on("data", (chunk) => (data += chunk));
          resp.on("end", () => resolve(data));
        })
        .on("error", (err) => {
          reject("Error: " + err.message);
        });
    })
  );
}

const config = new pulumi.Config();

const amiId = aws.ec2.getAmiOutput({
  owners: ["679593333241"], // AWS Marketplace
  mostRecent: true,
  filters: [
    {
      name: "name",
      values: ["Cisco Identity Services Engine (ISE) v3.3-*"],
    },
  ],
}).id;

const whitelistIpAddress =
  config.get("whitelistIpSecurityGroup") || fetchMyPublicIpOutput();

const vpc = new awsx.ec2.Vpc("ise-vpc", {
  numberOfAvailabilityZones: 1,
  natGateways: {
    strategy: "None",
  },
  subnetStrategy: "Auto",
  subnetSpecs: [{ type: "Public", name: "ise-provider-test-public-subnet" }],
  tags: {
    Name: "Cisco-ISE-Provider-VPC",
    Owner: "providers-team",
  },
});

const vpcSecurityGroup = new aws.ec2.SecurityGroup("ise-sg", {
  vpcId: vpc.vpc.id,
  description: "Allow inbound traffic on port 3389 from my IP",
  egress: [
    {
      protocol: "-1",
      fromPort: 0,
      toPort: 0,
      cidrBlocks: ["0.0.0.0/0"],
    },
  ],
  ingress: [
    {
      protocol: "tcp",
      fromPort: 443,
      toPort: 443,
      cidrBlocks: [pulumi.interpolate`${whitelistIpAddress}/32`],
    },
    {
      protocol: "tcp",
      fromPort: 22,
      toPort: 22,
      cidrBlocks: [pulumi.interpolate`${whitelistIpAddress}/32`],
    },
  ],
  tags: {
    Name: "CiscoISE-SG-Ingress-SSH-HTTPS-Egress-All",
    Owner: "providers-team",
  },
});

const sshKey = new tls.PrivateKey("ise-ssh-key", {
  algorithm: "RSA",
  rsaBits: 4096,
});

const ec2KeyPair = new aws.ec2.KeyPair("ise-keypair", {
  publicKey: sshKey.publicKeyOpenssh,
  tags: {
    Name: "cisco-ise-provider-keypair",
    Owner: "providers-team",
  },
});

const randomPassword = new random.RandomPassword("ise-password", {
  length: 12,
  special: true,
});

const elasticIp = new aws.ec2.Eip("ise-provider-eip", {
  domain: "vpc",
  tags: {
    Name: "cisco-ise-provider-eip",
    Owner: "providers-team",
  },
});

export const iseAdminUsername = "iseadmin";
const instance = new aws.ec2.Instance("ise-provider-instance-test-instance", {
  instanceType: "t3.xlarge", // Evaluation license size; see https://www.cisco.com/c/en/us/td/docs/security/ise/3-1/install_guide/b_ise_InstallationGuide31/m_ISEaaS.html
  vpcSecurityGroupIds: [vpcSecurityGroup.id],
  privateDnsNameOptions: {
    enableResourceNameDnsARecord: true,
    hostnameType: "resource-name",
  },
  subnetId: vpc.publicSubnetIds[0],
  ami: amiId, //Cisco ISE 3.3
  keyName: ec2KeyPair.keyName,
  ebsBlockDevices: [
    {
      deviceName: "/dev/sda1",
      volumeSize: 300,
      volumeType: "gp2",
    },
  ],
  userData: pulumi.interpolate`
    hostname=isepatch33
    username=${iseAdminUsername}
    password=${randomPassword.result}
    dnsdomain=cisco.com
    ersapi=yes
    openapi=yes
    pxGrid=no
    timezone=Etc/UTC
    pxgrid_cloud=no
    ntpserver=time.nist.gov
    primarynameserver=8.8.8.8
  `,
  tags: {
    Name: "cisco-ise-provider-test-instance",
    ProviderUrl: "https://github.com/pulumi/pulumi-ise",
    Owner: "providers-team",
  },
});

export const publicIpAddress = new aws.ec2.EipAssociation(
  "ise-eip-association",
  {
    instanceId: instance.id,
    allocationId: elasticIp.id,
  }
).publicIp;

export const iseAdminPassword = randomPassword.result;
