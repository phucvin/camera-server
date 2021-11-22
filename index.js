const gcp = require("@pulumi/gcp");

// Create a network
const network = new gcp.compute.Network("test-network");
const firewall = new gcp.compute.Firewall("test-firewall", {
    network: network.id,
    allows: [{
        protocol: "tcp",
        ports: [ "22", "80" ],
    }],
});

const startupScript = `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &`;

// Create a Virtual Machine Instance
const computeInstance = new gcp.compute.Instance("test-instance", {
    machineType: "f1-micro",
    zone: "us-west1-b",
    bootDisk: { initializeParams: { image: "debian-cloud/debian-9" } },
    networkInterfaces: [{
        network: network.id,
        // accessConfigus must includ a single empty config to request an ephemeral IP
        accessConfigs: [{}],
    }],
    metadataStartupScript: startupScript,
});

// Export the name and IP address of the Instance
exports.instanceName = computeInstance.name;
exports.instanceIP = computeInstance.networkInterfaces.apply(ni => ni[0].accessConfigs[0].natIp);