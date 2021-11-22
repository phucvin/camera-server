const gcp = require("@pulumi/gcp");

const zone = "us-west1-b";
const containerImage = "us-west1-docker.pkg.dev/phucvin/my-container-repo/camera-server";

const network = new gcp.compute.Network("camera-server-network");
const firewall = new gcp.compute.Firewall("camera-server-firewall", {
    network: network.id,
    allows: [{
        protocol: "tcp",
        ports: [ "22", "21", "21000-21010", "8000" ],
    }],
});

const bootDisk = new gcp.compute.Disk("camera-server", {
    zone: zone,
    image: "projects/cos-cloud/global/images/family/cos-stable",
});

const dataDisk = new gcp.compute.Disk("camera-server-data", {
    zone: zone,
    size: 10,
});

const startupScript = `#!/bin/bash
sudo mount /dev/sdb /mnt/disks/camera-server-data/
docker run --rm -d -p 21:21 -p 21000-21010:21000-21010 -v /mnt/disks/camera-server-data:/ftp/tom/upload -e USERS="tom|12345678" delfer/alpine-ftp-server
`;

const vm = new gcp.compute.Instance("camera-server", {
    allowStoppingForUpdate: true,
    machineType: "e2-micro",
    zone: zone,
    bootDisk: { source: bootDisk.id },
    attachedDisks: [
        { source: dataDisk.id },
    ],
    networkInterfaces: [{
        network: network.id,
        accessConfigs: [{}],
    }],
    metadataStartupScript: startupScript,
    metadata: {
        "gce-container-declaration":
`spec:
  containers:
  - image: ${containerImage}
    name: camera-web-server
    securityContext:
      privileged: false
    stdin: false
    tty: false
    volumeMounts:
    - mountPath: /home/tom/ftp/upload
      name: host-path-0
      readOnly: true
  restartPolicy: Always
  volumes:
  - hostPath:
      path: /mnt/disks/camera-server-data/
    name: host-path-0
`,
    },
    serviceAccount: {
        scopes: [
            "cloud-platform",
        ],
    },
});

exports.vmName = vm.name;
exports.vmIP = vm.networkInterfaces.apply(ni => ni[0].accessConfigs[0].natIp);